package pos

func (p Position) Neighbors() []Position {
	return []Position{
		p.Up(),
		p.Down(),
		p.Left(),
		p.Right(),
	}
}
