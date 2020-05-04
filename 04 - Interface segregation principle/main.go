package main

import "fmt"

/*
	Definition:
	Clients should not be forced to depend on methods they do not use.
	In other words we shouldn’t create big struct with a lot of behaviors, we should isolate behaviors,
	it can be also called interface polution
*/

/*--------------------------------------------------------*/
/*	 BREAKING INTEFRACE SEGREGATION PRINCIPLE	          */
/*--------------------------------------------------------*/

type Vehicle struct {
	Name string
}

func (v Vehicle) Run() {
	fmt.Println("Running... " + v.Name)
}

func (v Vehicle) PlayCD() {
	fmt.Println("Playing CD... " + v.Name)
}

type Motorbike struct {
	Vehicle
}

/*--------------------------------------------------------*/
/*	 IMPLEMENTING INTEFRACE SEGREGATION PRINCIPLE        */
/*--------------------------------------------------------*/
type IPlayCD interface {
	Play()
}

type IRun interface {
	Run()
}

type VehicleCDPlay struct {
	Song string
}

func (v VehicleCDPlay) Play() {
	fmt.Println("Paying CD... " + v.Song)
}

type VehicleRun struct {
	Name string
}

func (v VehicleRun) Run() {
	fmt.Println("Vehicle running... " + v.Name)
}

// Compose Vehicle that can run and play song
type VehicleGood struct {
	VehicleCDPlay
	VehicleRun
}

// Compose Motorbike that can run
type MotorbikeGood struct {
	VehicleRun
}

func main() {
	// Running bad example
	// We really don't need to playCD in motorbike right?
	m := &Motorbike{Vehicle{"Honda"}}
	m.PlayCD()

	// Running good example
	v1 := &VehicleGood{VehicleCDPlay{"My Song"}, VehicleRun{"BMW"}}
	v1.Run()
	v1.Play()

	m2 := &MotorbikeGood{VehicleRun{"Kawsaki"}}
	m2.Run()

}
