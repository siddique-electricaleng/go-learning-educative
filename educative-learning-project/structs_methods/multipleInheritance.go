package main

import "fmt"

// Define Models

type Camera struct{}

type Phone struct{}

type SmartPhone struct {
	Camera
	Phone
}

// Camera method

func (c *Camera) TakePic() string {
	return "Click"
}

// Phone method

func (p *Phone) Call() string {
	return "Ring Ring"
}

func mulInheritanceMain() {
	samsung := new(SmartPhone)

	fmt.Println("Out smartphone exhibits the following features")

	fmt.Println(samsung.TakePic())

	fmt.Println(samsung.Call())
}
