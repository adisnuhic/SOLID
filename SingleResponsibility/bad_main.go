// The example below represnts a way of breaking Single Responsibility Principle

package main

import "encoding/json"

import "os"

type Animal struct {
	Name string
	Age  uint16
}

func (a Animal) GetName() string {
	return a.Name
}

func (a Animal) GetAge() uint16 {
	return a.Age
}

// This is not responsibility of the Animal struct
func (a Animal) formatJSON() ([]byte, error) {
	json, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	return json, nil
}

func main() {
	tiger := &Animal{
		Name: "Tiger",
		Age:  12,
	}
	jsonTiger, errTiger := tiger.formatJSON()
	if errTiger != nil {
		panic(errTiger)
	}
	os.Stdout.Write(jsonTiger)

}
