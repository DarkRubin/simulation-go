package creature

import (
	"simulation/internal/occupier"
	"simulation/internal/positon"
)

type Herbivore struct {
	position     positon.Position
	healthPoints int
	speed        int
}

func (h Herbivore) Type() occupier.EntityType {
	return occupier.TypeHerbivore
}

func (h Herbivore) Pos() positon.Position {
	return h.position
}

func (h Herbivore) HP() int {
	return h.healthPoints
}

func (h Herbivore) Speed() int {
	return h.speed
}

func (h Herbivore) Move() {
	//TODO implement me
	panic("implement me")
}

func (h Herbivore) Eat() {
	//TODO implement me
	panic("implement me")
}

func NewHerbivore(pos positon.Position, hp int, speed int) Creature {
	return Herbivore{
		position:     pos,
		healthPoints: hp,
		speed:        speed,
	}
}
