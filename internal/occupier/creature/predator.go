package creature

import (
	"simulation/internal/occupier"
	"simulation/internal/pos"
)

type Predator struct {
	position     pos.Position
	healthPoints int
	damage       uint
	speed        uint
	tired        bool
}

func (p *Predator) Type() occupier.EntityType {
	return occupier.TypePredator
}

func (p *Predator) Pos() pos.Position {
	return p.position
}

func (p *Predator) HP() int {
	return p.healthPoints
}

func (p *Predator) Speed() uint {
	return p.speed
}

func (p *Predator) Damage() uint {
	return p.damage
}

func (p *Predator) TargetType() occupier.EntityType {
	return occupier.TypeHerbivore
}

func (p *Predator) Move(way []pos.Position) bool {
	if p.moved {
		panic("pred is already moved")
	}
	wayLen := len(way)
	if wayLen == 0 {
		return false
	}
	if uint(wayLen) <= p.Speed() {
		targ := way[wayLen-1]
		p.position = targ
	} else {
		p.position = way[p.Speed()]
	}
	p.moved = true
	return true
}

func (p *Predator) IsTired() bool {
	return p.tired
}

func (p *Predator) Attack(herbivore *Herbivore, hp int) {
	herbivore.Hurt(p.damage)
	p.Heal(hp)
	p.tired = true
}

func (p *Predator) Hurt(damage uint) {
	p.healthPoints -= int(damage)
}

func (p *Predator) Heal(hp int) {
	p.healthPoints += hp
}

func (p *Predator) Eat(food occupier.Occupier, heal int) {
	if food.Type() != p.TargetType() {
		panic("invalid type of provided food")
	}
	p.Move(food.Pos())
	p.Heal(heal)
}

func (p *Predator) Recover() {
	p.tired = false
}

func NewPredator(pos pos.Position, damage uint, hp int, speed uint) Creature {
	return &Predator{
		damage:       damage,
		healthPoints: hp,
		speed:        speed,
		position:     pos,
	}
}
