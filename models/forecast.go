package models

type Forecast interface {
	Name() string
}

func Name() string {
	return "Forecast"
}