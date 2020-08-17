package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/transforms"
	"github.com/go-audio/wav"
	"github.com/gorilla/mux"
	"github.com/kr/pretty"

	"github.com/jamestunnell/go-synth/unit"
	"github.com/jamestunnell/go-synth/unit/generators"
)

type MakeGeneratorDemoRequest struct {
	SampleRate float64            `json:"srate"`
	Params     map[string]float64 `json:"params,omitempty"`
}

const (
	DemoBitDepth = 16
	FormatPCM    = 1
	ChunkSize    = 50
)

func makeGeneratorDemo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plugin := findPlugin(vars["name"], generators.Builtin)
	if plugin == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := ReadRequestData(r)
	if err != nil {
		log.Printf("failed to read request data: %v", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	var request MakeGeneratorDemoRequest

	if err = json.Unmarshal(data, &request); err != nil {
		log.Printf("failed to unmarshal request JSON object: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	log.Printf("make demo request: %# v", pretty.Formatter(request))

	if !isSampleRateValid(request.SampleRate) {
		log.Printf("sample rate %f is invalid", request.SampleRate)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if plugin.Interface.NumOutputs == 0 {
		log.Print("generator has no outputs")

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	paramBuffers := make(map[string]*unit.Buffer)

	// get the parameter values from the request or use defaults
	for paramName, param := range plugin.Interface.Parameters {
		paramBuffers[paramName] = unit.NewBuffer(1)

		val, found := request.Params[paramName]
		if found {
			paramBuffers[paramName].Values[0] = val
		} else {
			if param.Required {
				log.Printf("required generator param %s missing", paramName)

				w.WriteHeader(http.StatusBadRequest)

				return
			} else {
				paramBuffers[paramName].Values[0] = param.Default
			}
		}
	}

	gen := plugin.NewUnit()
	outputBuffer := unit.NewBuffer(ChunkSize)

	err = gen.Initialize(
		request.SampleRate,
		paramBuffers,
		[]*unit.Buffer{},
		[]*unit.Buffer{outputBuffer})
	if err != nil {
		log.Printf("failed to initialize generator: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	gen.Configure()

	numSamples := int(request.SampleRate)

	buffer := &audio.FloatBuffer{
		Format: &audio.Format{NumChannels: 1, SampleRate: int(request.SampleRate)},
		Data:   make([]float64, numSamples),
	}

	for i := 0; i < numSamples; i += ChunkSize {
		gen.Sample()
		for j := 0; j < ChunkSize; j++ {
			buffer.Data[i+j] = outputBuffer.Values[j]
		}
	}

	// Clip any samples that are not within the range (-1,1)
	for i := 0; i < numSamples; i++ {
		if buffer.Data[i] >= 1.0 {
			buffer.Data[i] = 1.0 - 1e-5
		} else if buffer.Data[i] <= -1.0 {
			buffer.Data[i] = -1.0 + 1e-5
		}
	}

	transforms.PCMScale(buffer, DemoBitDepth)

	intBuffer := buffer.AsIntBuffer()
	intBuffer.SourceBitDepth = 16

	// We can use a pattern of "pre-*.txt" to get an extension like: /tmp/pre-123456.txt
	tmpFile, err := ioutil.TempFile(os.TempDir(), "demo-*.wav")
	if err != nil {
		log.Printf("failed to create temporary file: %v", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	defer os.Remove(tmpFile.Name())

	err = writeWAV(tmpFile, intBuffer, DemoBitDepth)
	if err != nil {
		log.Printf("failed to write WAV file: %v", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	tmpFile.Close()

	wavFileName := tmpFile.Name()
	// flacFileName := strings.Replace(wavFileName, ".wav", ".flac", 1)

	// err = audio.EncodeFLAC(tmpFile.Name(), flacFileName)
	// if err != nil {
	// 	log.Printf("failed to encode FLAC file: %v", err)

	// 	w.WriteHeader(http.StatusInternalServerError)

	// 	return
	// }

	// defer os.Remove(flacFileName)

	w.Header().Set("Content-Type", "audio/wav")
	http.ServeFile(w, r, wavFileName)
}

func writeWAV(file *os.File, buf *audio.IntBuffer, bitDepth int) error {
	// setup the encoder and write all the frames
	encoder := wav.NewEncoder(file, buf.Format.SampleRate,
		bitDepth, buf.Format.NumChannels, FormatPCM)

	if err := encoder.Write(buf); err != nil {
		return err
	}

	// close the encoder to make sure the headers are properly
	// set and the data is flushed.
	if err := encoder.Close(); err != nil {
		return err
	}

	return nil
}

// func EncodeFLAC(srcFile, dstFlacFile string) error {
// 	cmd := exec.Command("flac","-o",dstFlacFile,srcFile)
// 	_, err := cmd.Output()

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func isSampleRateValid(srate float64) bool {
	switch srate {
	case 22050:
		return true
	case 44100:
		return true
	case 48000:
		return true
	case 96000:
		return true
	case 192000:
		return true
	}

	return false
}
