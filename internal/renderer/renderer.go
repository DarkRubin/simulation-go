package renderer

import (
	"simulation/internal/worldmap"
)

type Renderer interface {
	Render(m worldmap.Map) string
}
