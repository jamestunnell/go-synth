package util

import (
	"fmt"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/transforms"
	"github.com/go-audio/wav"
	"github.com/jamestunnell/go-synth"
)

// RenderParams are used in rendering a WAV file
type RenderParams struct {
	SampleRate int
	DurSec     float64
	BitDepth   int
}

type ErrWrongOutputCount struct {
	Expected, Actual int
}

type ErrWrongOutputType struct {
	Expected, Actual string
}

func (err *ErrWrongOutputCount) Error() string {
	const strFmt = "block has wrong output count: expected %d, actual %d"

	return fmt.Sprintf(strFmt, err.Expected, err.Actual)
}

func (err *ErrWrongOutputType) Error() string {
	const strFmt = "block has wrong output type: expected %s, actual %s"

	return fmt.Sprintf(strFmt, err.Expected, err.Actual)
}

// NumSamples determines the total number of samples based on
// duration and sample rate.
func (p *RenderParams) NumSamples() int {
	return int(p.DurSec * float64(p.SampleRate))
}

// FormatPCM is the PCM format used in rendering
const FormatPCM = 1

// RenderWAV renders audio to a WAV file.
// Returns a non-nil error in case of failure.
func RenderWAV(src synth.Block, wavFile *os.File, params *RenderParams) error {
	numSamples := params.NumSamples()

	dstBuf := &audio.FloatBuffer{
		Format: &audio.Format{NumChannels: 1, SampleRate: params.SampleRate},
		Data:   make([]float64, numSamples),
	}

	ifc := synth.GetInterface(src)
	if len(ifc.Outputs) != 1 {
		return &ErrWrongOutputCount{
			Expected: 1,
			Actual:   len(ifc.Outputs),
		}
	}

	var output synth.Output

	for _, o := range ifc.Outputs {
		if o.Type() != "float64" {
			return &ErrWrongOutputType{
				Expected: "float64",
				Actual:   o.Type(),
			}
		}

		output = o
	}

	srcBuf := ConnectedBuffer()
	chunkSize := len(srcBuf)

	for i := 0; i < numSamples; i += chunkSize {
		src.Run()

		jLim := chunkSize
		if i+chunkSize > numSamples {
			jLim = numSamples - i
		}

		for j := 0; j < jLim; j++ {
			dstBuf.Data[i+j] = srcBuf[j]
		}
	}

	// Clip any samples that are not within the range (-1,1)
	for i := 0; i < numSamples; i++ {
		if dstBuf.Data[i] >= 1.0 {
			dstBuf.Data[i] = 1.0 - 1e-5
		} else if dstBuf.Data[i] <= -1.0 {
			dstBuf.Data[i] = -1.0 + 1e-5
		}
	}

	transforms.PCMScale(dstBuf, params.BitDepth)

	intBuffer := dstBuf.AsIntBuffer()
	intBuffer.SourceBitDepth = params.BitDepth

	return writeWAV(wavFile, intBuffer, params.BitDepth)
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
