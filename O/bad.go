package main

import "fmt"

type (
	//Player struct{}
	Bread  struct{}
)

func (b *Bread) Heal() {
	fmt.Println("Healing 50 HP.")
}

func (b *Bread) DamageBuff() {
	fmt.Println("Buffing 10% of max damage.")
}

func (p *Player) EatBread(bread *Bread) {
	bread.Heal()
	bread.DamageBuff()
}

// func main() {
// 	player := &Player{}
// 	bread := &Bread{}

// 	player.EatBread(bread)
// }