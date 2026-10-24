package main

// migrate to encoding/json/v2
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"unicode/utf8"
)

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

func errorHandlingDemo() {
	var newErr error = errors.New("some error")
	newErrType := reflect.TypeOf(newErr)
	fmt.Printf("new error type: %s\n", newErrType)

	_, err := os.Open("missing.file")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		fmt.Printf("failed to open file: %s\n", err.Error())

		errType := reflect.TypeOf(err)
		fmt.Printf("error type: %s\n", errType)
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
	errorHandlingDemo()
}
