package main

import "fmt"

type (
	Player struct{}
	Item   struct {
		Name      string
		Abilities []Ability
	}
	Heal       struct{}
	DamageBuff struct{}

	Ability interface {
		Execute()
	}
)

func (p *Player) Use(item *Item) {
	fmt.Println("Use: ", item.Name)
	for _, ability := range item.Abilities {
		ability.Execute()
	}
}

func (h *Heal) Execute() {
	fmt.Println("Healing 50 HP.")
}

func (d *DamageBuff) Execute() {
	fmt.Println("Buffing 10% of max damage.")
}

func main() {
	player := &Player{}
	bread := &Item{
		Name: "Bread",
		Abilities: []Ability{
			&Heal{},
			&DamageBuff{},
		},
	}
	player.Use(bread)
}