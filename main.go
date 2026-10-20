package main

import (
	"context"
	"encoding/json/v2"
	"fmt"
	"reflect"
	"unicode/utf8"
)

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

func jsonWriteDemo() {
	fmt.Println("*** json write demo ***")
	car := Car{"Mercedes-Benz", "E200d"}
	// passing pointer idiomatically
	out, err := json.Marshal(&car)
	if err != nil {
		fmt.Println("json marshalling failed")
	} else {
		fmt.Printf("car json: %s\n", string(out))
	}
}

func interfaceVarDemo() {
	fmt.Println("*** initial context demo ***")
	var ctx Ctx
	fmt.Printf("context: %v\n", ctx)
}

func backgroundContextDemo() {
	ctx := context.Background()
	fmt.Printf("background context: %v\n", ctx)
}

func typeAssertDemo(obj any) {
	// pointer type assertion should be explicit
	if car, ok := obj.(*Car); ok {
		fmt.Printf("it's a car %v\n", car)
	} else {
		fmt.Println("it's not a car")
	}
}

func reflectDemo() {
	var obj any = New("Audi", "A6")
	objType := reflect.TypeOf(obj)
	fmt.Printf("obj type: %s\n", objType)
}

func structDemo() {
	fmt.Println("*** struct demo ***")
	// var car Car
	car := Car{"BMW", "320d"}
	fmt.Printf("Car: {'%s'}\n", car.Maker)
	fmt.Printf("Car maker len: %d\n", len(car.Maker))
	if utf8.RuneCountInString(car.Maker) > 0 {
		fmt.Printf("car maker: %s\n", car.Maker)
	}
}

func main() {
	structDemo()
	jsonWriteDemo()
	interfaceVarDemo()
	backgroundContextDemo()

	car := &Car{"Mercedes-Benz", "E200d"}
	typeAssertDemo(car)

	reflectDemo()
}
