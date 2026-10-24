package model

type Car struct {
	Maker string `json:"maker"`
	Model string `json:"model"`
}

func New(maker string, model string) *Car {
	return &Car{maker, model}
}

type Ctx interface {
	Do()
}
