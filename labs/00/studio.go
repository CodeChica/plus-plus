package main

type Studio int

const (
	Pixar Studio = iota
	CastleRock
	MiramaxFilms
	RegencyEnterprises
	Disney
)

func (s Studio) String() string {
	return ""
}
