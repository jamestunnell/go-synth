package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/akrennmair/slice"
	"github.com/go-audio/audio"
	"github.com/go-audio/transforms"
	"github.com/go-audio/wav"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/network"
)

// RenderParams are used in rendering a WAV file
type RenderParams struct {
	SampleRate int
	DurSec     float64
	BitDepth   int
	ChunkSize  int
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
func RenderWAV(net *network.Network, wavFile *os.File, params *RenderParams) error {
	net.AddDefaultBlocks()

	if errs := net.Validate(); len(errs) > 0 {
		errStrings := slice.Map(errs, func(err error) string {
			return err.Error()
		})
		lines := strings.Join(errStrings, "\n")

		return fmt.Errorf("network is invalid:\n%s", lines)
	}

	if err := net.ApplyConnections(); err != nil {
		return fmt.Errorf("failed to apply connections: %w", err)
	}

	if err := net.InitializeBlocks(float64(params.SampleRate), params.ChunkSize); err != nil {
		return fmt.Errorf("failed to apply connections: %w", err)
	}

	f, err := net.MakeConfigureAndRunFunc()
	if err != nil {
		return fmt.Errorf("failed to make config-and-run func: %w", err)
	}

	terminal := net.TerminalBlocks()[0]
	switch t := terminal.(type) {
	case *synth.MonoTerminal:
		return RenderWAVMono(wavFile, params, t, f)
	}

	return nil
}

func RenderWAVMono(
	wavFile *os.File,
	params *RenderParams,
	mono *synth.MonoTerminal,
	configAndRun func()) error {
	numSamples := params.NumSamples()

	dstBuf := &audio.FloatBuffer{
		Format: &audio.Format{NumChannels: 1, SampleRate: params.SampleRate},
		Data:   make([]float64, numSamples),
	}

	srcBuf := mono.In.Output.Buffer

	for i := 0; i < numSamples; i += params.ChunkSize {
		configAndRun()

		jLim := params.ChunkSize
		if i+params.ChunkSize > numSamples {
			jLim = numSamples - i
		}

		for j := 0; j < jLim; j++ {
			dstBuf.Data[i+j] = srcBuf[j]
		}
	}

	clipSamples(dstBuf.Data)

	transforms.PCMScale(dstBuf, params.BitDepth)

	intBuffer := dstBuf.AsIntBuffer()
	intBuffer.SourceBitDepth = params.BitDepth

	return writeWAV(wavFile, intBuffer, params.BitDepth)
}

// clipSamples clips any samples that are not within the range (-1,1)
func clipSamples(samples []float64) {
	for i := 0; i < len(samples); i++ {
		if samples[i] >= 1.0 {
			samples[i] = 1.0 - 1e-5
		} else if samples[i] <= -1.0 {
			samples[i] = -1.0 + 1e-5
		}
	}
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
