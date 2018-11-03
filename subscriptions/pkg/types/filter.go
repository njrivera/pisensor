package types

import (
	"github.com/pisensor/pkg/models"
)

const (
	MatchAll = "*"
)

type ClientFilter struct {
	Serials []string `json:"serials,omitempty"`
}

type ServerFilter struct {
	Serials  map[string]struct{}
	matchAll bool
}

func NewServerFilter(cf ClientFilter) *ServerFilter {
	sf := &ServerFilter{}

	if len(cf.Serials) == 1 && cf.Serials[0] == MatchAll {
		sf.matchAll = true
	} else {
		sf.Serials = arrayToMap(cf.Serials)
	}

	return sf
}

func (filter *ServerFilter) Check(r models.TempReading) bool {
	if filter.matchAll {
		return true
	}

	_, ok := filter.Serials[r.Serial]

	return ok
}

func arrayToMap(arr []string) map[string]struct{} {
	m := map[string]struct{}{}

	if arr != nil {
		for _, str := range arr {
			m[str] = struct{}{}
		}
	}

	return m
}
