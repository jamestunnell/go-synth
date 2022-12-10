package network

import (
	"golang.org/x/exp/slices"
)

type Connections []*Connection

func (conns Connections) Equal(other Connections) bool {
	if len(conns) != len(other) {
		return false
	}

	for _, conn := range conns {
		conn2, found := other.FindBySource(conn.Source)
		if !found {
			return false
		}

		if !conn2.Dest.Equal(conn.Dest) {
			return false
		}
	}

	return true
}

func (conns Connections) FindBySource(source *Address) (*Connection, bool) {
	for _, conn := range conns {
		if conn.Source.Equal(source) {
			return conn, true
		}
	}

	return nil, false
}

func (conns Connections) FindByDest(dest *Address) (*Connection, bool) {
	for _, conn := range conns {
		if conn.Dest.Equal(dest) {
			return conn, true
		}
	}

	return nil, false
}

func (conns Connections) OverusedSources() []*Address {
	strs := []string{}
	duplicated := []*Address{}

	for _, conn := range conns {
		str := conn.Source.String()
		if slices.Contains(strs, str) {
			duplicated = append(duplicated, conn.Source)
		} else {
			strs = append(strs, str)
		}
	}

	return duplicated
}

func (conns Connections) OverusedDests() []*Address {
	destStrings := []string{}
	duplicated := []*Address{}

	for _, conn := range conns {
		str := conn.Dest.String()
		if slices.Contains(destStrings, str) {
			duplicated = append(duplicated, conn.Dest)
		} else {
			destStrings = append(destStrings, str)
		}
	}

	return duplicated
}
