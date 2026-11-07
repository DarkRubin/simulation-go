package creature

import (
	"simulation/internal/occupier"
	"simulation/internal/pos"
)

type Herbivore struct {
	position     pos.Position
	healthPoints int
	speed        uint
	tired        bool
}

func (h *Herbivore) Type() occupier.EntityType {
	return occupier.TypeHerbivore
}

func (h *Herbivore) Pos() pos.Position {
	return h.position
}

func (h *Herbivore) HP() int {
	return h.healthPoints
}

func (h *Herbivore) Speed() uint {
	return h.speed
}

func (h *Herbivore) IsTired() bool {
	return h.tired
}

func (h *Herbivore) Move(way []pos.Position) bool {
	if h.moved {
		panic("herb is already moved")
	}
	wayLen := len(way)
	if wayLen == 0 {
		return false
	}
	if uint(wayLen) <= h.speed {
		targ := way[wayLen-1]
		h.position = targ
	} else {
		h.position = way[h.speed]
	}
	h.moved = true
	return true
}

func (h *Herbivore) Eat(food occupier.Occupier) bool {
	if food.Type() != h.TargetType() {
		panic("invalid type of provided food")
	}
	if h.Moved() {
		return false
	}
	h.position = food.Pos()
	h.moved = true
	h.healthPoints += 5
	return true
}

func (h *Herbivore) TargetType() occupier.EntityType {
	return occupier.TypeGrass
}

func (h *Herbivore) Hurt(damage uint) {
	h.healthPoints -= int(damage)
}

func (h *Herbivore) Recover() {
	h.moved = false
}

func NewHerbivore(pos pos.Position, hp int, speed uint) Creature {
	return &Herbivore{
		position:     pos,
		healthPoints: hp,
		speed:        speed,
	}
}
