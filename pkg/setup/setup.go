package setup

import (
	"simulation/internal/action"
	"simulation/internal/action/inits"
	"simulation/internal/action/turn"
	"simulation/internal/renderer"
	"simulation/internal/terminal"
	"simulation/internal/wayfinder"
	"simulation/internal/worldmap"
	"simulation/pkg/setting"
	"simulation/pkg/ui"
)

func UserInterface() ui.Dialog {
	return terminal.NewStdOutDialog()
}

func CallbackFunction(settings setting.Settings) ui.Callback {
	buffSize := int((settings.Width * settings.Height) * 2)
	rend := renderer.NewEmojiRenderer(buffSize)
	return terminal.NewPrinter(rend).CurrState
}

func WorldMap(settings setting.Settings) *worldmap.Map {
	return worldmap.New(settings.Height, settings.Width)
}

func InitActions(settings setting.Settings) []action.Action {
	return []action.Action{
		inits.SpawnOccupiersAction{Settings: settings},
	}
}

func TurnActions(settings setting.Settings) []action.Action {
	wayFinder := wayfinder.New(settings.Algorithm)

	return []action.Action{
		turn.RecoverCreatureAction{},
		turn.HerbivoresEatAction{},
		turn.PredatorsAttackAction{},
		turn.CreatureMoveAction{Wayfinder: wayFinder},
		turn.HungerCreatureAction{},
		turn.RemoveDeadCreaturesAction{},
		turn.RespawnOccupiersAction{Settings: settings},
	}
}
