package action

import (
	"simulation/internal/worldmap"
)

type Action interface {
	Act(*worldmap.Map)
}
