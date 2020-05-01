package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*----------------------------------------------------*/
/*		BREAKING SINGLE RESPONSIBILITY PRINCIPLE	  */
/*----------------------------------------------------*/
type AnimalBad struct {
	Name string
	Age  uint16
}

func (a AnimalBad) GetName() string {
	return a.Name
}

func (a AnimalBad) GetAge() uint16 {
	return a.Age
}

// This is not responsibility of the Animal struct
func (a AnimalBad) formatJSON() ([]byte, error) {
	json, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return json, nil
}

/*--------------------------------------------------------*/
/*		IMPLEMENTING SINGLE RESPONSIBILITY PRINCIPLE	  */
/*--------------------------------------------------------*/
type AnimalGood struct {
	Name string
	Age  uint16
}

func (a AnimalGood) GetName() string {
	return a.Name
}

func (a AnimalGood) GetAge() uint16 {
	return a.Age
}

type ToJSON struct{}

// FormatJSON is a responsibility of ToJSON struct
func (ToJSON) FormatJson(a *AnimalGood) ([]byte, error) {
	myJSON, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return myJSON, nil
}

func main() {
	// Running a bad way
	tiger := &AnimalBad{
		Name: "Tiger",
		Age:  12,
	}
	jsonTiger, errTiger := tiger.formatJSON()
	if errTiger != nil {
		panic(errTiger)
	}
	os.Stdout.Write(jsonTiger)
	fmt.Println()

	// Running a good way
	duck := &AnimalGood{
		Name: "Duck",
		Age:  7,
	}
	jsonDuck, errDuck := (&ToJSON{}).FormatJson(duck)
	if errDuck != nil {
		panic(errDuck)
	}
	os.Stdout.Write(jsonDuck)
	fmt.Println()

}
