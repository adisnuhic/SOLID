package main

import "fmt"

/*
	Definition:
	The Liskov Substitution Principle encourages you to express the dependencies between your packages in terms of interfaces,
	not concrete types. By defining small interfaces, we can be more confident that implementations will faithfully satisfy their
	contract.
*/

/*----------------------------------------------------------------------------------------------------------

But LSP principle is safe in GO language since GO doesn't have inheritance but it has more powerful composition.
In composition it is not allowed to substitute parent struct by child struct

----------------------------------------------------------------------------------------------------------*/

// Example of impossible substitution parent struct by child struct

type Animal struct {
	Name string
}

func (a Animal) PrintName() {
	fmt.Println(a.Name)
}

type Bird struct {
	Animal
}

func ImpossibleLiskovSubstitution(a *Animal) {
	a.PrintName()
}

/*------------------------------------------------------------------------------------------*/

// In GO lang we are obliged to define behaviour in an interface

type IAnimal interface {
	PrintMyName()
}

type AnimalGood struct {
	Name string
}

func (a AnimalGood) PrintMyName() {
	fmt.Println(a.Name)
}

type BirdGood struct {
	AnimalGood
}

func PossibleLiskovSubstitution(a IAnimal) {
	a.PrintMyName()
}

func main() {
	// This will not work
	//eagle := &Bird{Animal{"Eagle"}}
	//ImpossibleLiskovSubstitution(eagle) // Error: cannot use eagle (type *Bird) as type *Animal in argument to ImpossibleLiskovSubstitution

	wolf := &AnimalGood{"Wolf"}
	PossibleLiskovSubstitution(wolf)
	fmt.Println()
	owl := &BirdGood{AnimalGood{"Owl"}}
	PossibleLiskovSubstitution(owl)

}
