package main

// import "fmt"

// type (
// 	Vehicle interface {
// 		Drive()
// 		Ride()
// 		Aircondition()
// 	}

// 	Car        struct{}
// 	Motorcycle struct{}
// )

// func (c *Car) Drive() {
// 	fmt.Println("Car can drive.")
// }

// func (c *Car) Ride() {
// 	panic("Car cannot ride.")
// }

// func (c *Car) Aircondition() {
// 	fmt.Println("Car have airconditioner.")
// }

// func (m *Motorcycle) Drive() {
// 	panic("Motorcycle cannot drive.")
// }

// func (m *Motorcycle) Ride() {
// 	fmt.Println("Motorcycle can ride.")
// }

// func (m *Motorcycle) Aircondition() {
// 	panic("Motorcycle have no airconditioner.")
// }

// func main() {
// 	car := &Car{}
// 	motorcycle := &Motorcycle{}

// 	car.Drive()
// 	car.Ride()
// 	car.Aircondition()

// 	motorcycle.Drive()
// 	motorcycle.Ride()
// 	motorcycle.Aircondition()
// }