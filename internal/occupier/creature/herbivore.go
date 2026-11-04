package creature

import (
	"simulation/internal/occupier"
	"simulation/internal/pos"
)

type Herbivore struct {
	position     pos.Position
	healthPoints int
	speed        int
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

func (h *Herbivore) Speed() int {
	return h.speed
}

func (h *Herbivore) Move(way []pos.Position) {
	wayLen := len(way)
	if wayLen == 0 {
		return
	}
	if wayLen <= h.Speed() {
		targ := way[wayLen-1]
		h.position = targ
	} else {
		h.position = way[h.Speed()]
	}
}

func (h *Herbivore) Eat(food occupier.Occupier) {
	h.position = food.Pos()
	h.healthPoints += 5
}

func (h *Herbivore) TargetType() occupier.EntityType {
	return occupier.TypeGrass
}

func NewHerbivore(pos pos.Position, hp int, speed int) Creature {
	return &Herbivore{
		position:     pos,
		healthPoints: hp,
		speed:        speed,
	}
}
