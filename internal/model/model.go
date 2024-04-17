package model

import (
	"log"
	"os"
	"plugin"
)

type Model struct {
	library string
}

func NewModel(s string) *Model {
	return &Model{library: s}
}

func (m Model) Calc(s *string, x string) (float64, error) {
	plug, err := plugin.Open(m.library)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	symCalc, err := plug.Lookup("Calc")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Приведение типа к функции Calc
	calcFunc := symCalc.(func(*string, string) (float64, error))

	return calcFunc(s, x)
}
