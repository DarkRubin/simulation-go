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
	moved        bool
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

func (p *Predator) Moved() bool {
	return p.moved
}

func (p *Predator) Eat(food occupier.Occupier) bool {
	if p.moved {
		return false
	}
	p.moved = true
	if food.Type() != p.TargetType() {
		panic("invalid type of provided food")
	}
	herb := food.(*Herbivore)
	if herb.HP() > 0 {
		herb.Hurt(p.damage)
		p.healthPoints += 3
		return false
	}

	p.position = food.Pos()
	p.healthPoints += 5
	return true
}

func (p *Predator) Recover() {
	p.healthPoints -= 1
	p.moved = false
}

func NewPredator(pos pos.Position, damage uint, hp int, speed uint) Creature {
	return &Predator{
		damage:       damage,
		healthPoints: hp,
		speed:        speed,
		position:     pos,
	}
}
