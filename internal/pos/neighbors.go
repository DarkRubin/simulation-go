package pos

func (p Position) Neighbors() []Position {
	return []Position{
		p.Up(),
		p.Down(),
		p.Left(),
		p.Right(),
	}
}

func (p Position) Up() Position {
	return Position{p.x, p.y - 1}
}

func (p Position) Down() Position {
	return Position{p.x, p.y + 1}
}

func (p Position) Left() Position {
	return Position{p.x - 1, p.y}
}

func (p Position) Right() Position {
	return Position{p.x + 1, p.y}
}
