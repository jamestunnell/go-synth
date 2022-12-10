package main

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-synth"
	"github.com/jamestunnell/go-synth/network"
	"github.com/jamestunnell/go-synth/unit/gen/osc"
	"github.com/jamestunnell/go-synth/unit/proc/math"
	"github.com/jamestunnell/go-synth/util"
)

func main() {
	n := network.New()

	n.Blocks["A"] = osc.NewTriangle()
	n.Blocks["B"] = osc.NewSine()
	n.Blocks["C"] = math.NewMul()
	n.Blocks["D"] = synth.NewConst(20.0)
	n.Blocks["E"] = synth.NewConst(200.0)
	n.Blocks["F"] = synth.NewMonoTerminal()

	n.Connections = network.Connections{
		{
			Source: &network.Address{Block: "D", Port: "Out"},
			Dest:   &network.Address{Block: "A", Port: "Freq"},
		},
		{
			Source: &network.Address{Block: "E", Port: "Out"},
			Dest:   &network.Address{Block: "B", Port: "Freq"},
		},
		{
			Source: &network.Address{Block: "A", Port: "Out"},
			Dest:   &network.Address{Block: "C", Port: "In1"},
		},
		{
			Source: &network.Address{Block: "B", Port: "Out"},
			Dest:   &network.Address{Block: "C", Port: "In2"},
		},
		{
			Source: &network.Address{Block: "C", Port: "Out"},
			Dest:   &network.Address{Block: "F", Port: "In"},
		},
	}

	wavFile, err := os.Create("test.wav")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create WAV")
	}

	params := &util.RenderParams{
		SampleRate: 48000,
		DurSec:     2.0,
		BitDepth:   16,
		ChunkSize:  100,
	}

	err = util.RenderWAV(n, wavFile, params)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to render WAV")
	}
}
