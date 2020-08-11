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

	"github.com/jamestunnell/go-synth/pkg/unit"
	"github.com/jamestunnell/go-synth/pkg/unit/generators"
)

type MakeGeneratorDemoRequest struct {
	Params map[string]float64 `json:"params,omitempty"`
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

	srate, ok := getSrate(vars)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := ReadRequestData(r)
	if err != nil {
		log.Printf("failed to read request data: %v", err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	var requestParams map[string]float64

	if err = json.Unmarshal(data, &requestParams); err != nil {
		log.Printf("failed to unmarshal request data: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	log.Printf("request params: %# v", pretty.Formatter(requestParams))

	ifc := plugin.GetInterface(float64(srate))

	if ifc.NumOutputs == 0 {
		log.Print("generator has no outputs")

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	paramBuffers := make(map[string]*unit.Buffer)

	// get the parameter values from the request or use defaults
	for _, param := range ifc.Parameters {
		paramBuffers[param.Name] = unit.NewBuffer(1)

		val, found := requestParams[param.Name]
		if found {
			paramBuffers[param.Name].Values[0] = val
		} else {
			if param.Required {
				log.Printf("required generator param %s missing", param.Name)

				w.WriteHeader(http.StatusBadRequest)

				return
			} else {
				paramBuffers[param.Name].Values[0] = param.Default
			}
		}
	}

	gen := plugin.NewUnit()
	outputBuffer := unit.NewBuffer(ChunkSize)

	err = gen.Initialize(
		float64(srate),
		paramBuffers,
		[]*unit.Buffer{},
		[]*unit.Buffer{outputBuffer})
	if err != nil {
		log.Printf("failed to initialize generator: %v", err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	numSamples := srate

	buffer := &audio.FloatBuffer{
		Format: &audio.Format{NumChannels: 1, SampleRate: srate},
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
