package pos

import "math/rand/v2"

type Position struct {
	x uint
	y uint
}

func (p Position) X() uint {
	return p.x
}

func (p Position) Y() uint {
	return p.y
}

func Pos(x uint, y uint) Position {
	return Position{x, y}
}

func RandomPos(xLim, yLim uint) Position {
	return Position{rand.UintN(xLim), rand.UintN(yLim)}
}
