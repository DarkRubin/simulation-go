package creature

import (
	"simulation/internal/occupier"
)

type Creature interface {
	occupier.Occupier
	HP() int
	Speed() int
	Move()
}
