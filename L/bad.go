// package main

// import "fmt"

// type (
// 	Aircraft interface {
// 		Fly()
// 		Fire()
// 	}

// 	Boeing747 struct{}
// 	Spitfire  struct{}
// )

// func (a *Boeing747) Fly() {
// 	fmt.Println("Boeing can fly.")
// }

// func (a *Boeing747) Fire() {
// 	panic("Boeing can fire???")
// }

// func (a *Spitfire) Fly() {
// 	fmt.Println("Spitfire can fly.")
// }

// func (a *Spitfire) Fire() {
// 	fmt.Println("Spitfire can fire.")
// }

// func main() {
// 	boeing747 := &Boeing747{}
// 	spitfire := &Spitfire{}

// 	boeing747.Fire()
// 	boeing747.Fly()

// 	spitfire.Fire()
// 	spitfire.Fly()
// }