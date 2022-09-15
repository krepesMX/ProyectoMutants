package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type DnaData struct {
	Dna []string `json:"dna,omitempty"`
}

func (u *DnaData) ToJson() ([]byte, error) {
	return json.Marshal(u.Dna)
}

type Metadata interface {
}
