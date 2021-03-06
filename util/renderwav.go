package util

import (
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/transforms"
	"github.com/go-audio/wav"

	"github.com/jamestunnell/go-synth/node"
)

// RenderParams are used in rendering a WAV file
type RenderParams struct {
	SampleRate int
	DurSec     float64
	BitDepth   int
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
func RenderWAV(out *node.Node, wavFile *os.File, params *RenderParams) error {
	numSamples := params.NumSamples()

	buffer := &audio.FloatBuffer{
		Format: &audio.Format{NumChannels: 1, SampleRate: params.SampleRate},
		Data:   make([]float64, numSamples),
	}

	chunkSize := out.Output().Length

	for i := 0; i < numSamples; i += chunkSize {
		out.Run()

		jLim := chunkSize
		if i+chunkSize > numSamples {
			jLim = numSamples - i
		}

		for j := 0; j < jLim; j++ {
			buffer.Data[i+j] = out.Output().Values[j]
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

	transforms.PCMScale(buffer, params.BitDepth)

	intBuffer := buffer.AsIntBuffer()
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
