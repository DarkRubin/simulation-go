package positon

import "math/rand/v2"

type Position struct {
	x uint
	y uint
}

func RandomPos(xLim, yLim uint) Position {
	return Position{rand.UintN(xLim), rand.UintN(yLim)}
}

func Pos(x uint, y uint) Position {
	return Position{x, y}
}
