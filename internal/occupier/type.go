package occupier

type EntityType int

const (
	TypeGrass EntityType = iota
	TypeRock
	TypeTree
	TypeHerbivore
	TypePredator
	TypeCount_
)

var EntityTypeString = map[EntityType]string{
	TypeGrass:     "Grass",
	TypeRock:      "Rock",
	TypeTree:      "Tree",
	TypeHerbivore: "Herbivore",
	TypePredator:  "Predator",
}

func (e EntityType) String() string {
	return EntityTypeString[e]
}
