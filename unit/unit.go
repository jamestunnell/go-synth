package unit

import "fmt"

type Unit interface {
	Initialize(
		srate float64,
		paramBuffers map[string]*Buffer,
		inBuffers,
		outBuffers []*Buffer) error
	Configure()
	Sample()
}

func FindNamedBuffer(namedBuffers map[string]*Buffer, bufferName string) (*Buffer, error) {
	buf, found := namedBuffers[bufferName]
	if !found {
		return nil, fmt.Errorf("%s buffer not found", bufferName)
	}

	return buf, nil
}
