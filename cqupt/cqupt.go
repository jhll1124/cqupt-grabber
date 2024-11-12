package cqupt

import (
	"grab"
	"query"
)

type Engine struct {
	Queryer query.Queryer
	Grabber grab.Grabber
}

func New() *Engine {
	h := &Engine{
		Queryer: query.Queryer{},
		Grabber: grab.Grabber{},
	}
	return h
}
