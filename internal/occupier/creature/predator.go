package creature

import (
	"simulation/internal/occupier"
	"simulation/internal/positon"
)

type Predator struct {
	position     positon.Position
	damage       int
	healthPoints int
	speed        int
}

func (p Predator) Type() occupier.EntityType {
	return occupier.TypePredator
}

func (p Predator) Pos() positon.Position {
	return p.position
}

func (p Predator) HP() int {
	return p.healthPoints
}

func (p Predator) Speed() int {
	return p.speed
}

func (p Predator) Damage() int {
	return p.damage
}

func (p Predator) Move() {
	//TODO implement me
	panic("implement me")
}

func NewPredator(pos positon.Position, damage int, hp int, speed int) Creature {
	return Predator{
		damage:       damage,
		healthPoints: hp,
		speed:        speed,
		position:     pos,
	}
}
