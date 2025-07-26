package game_process

import "github.com/Pavel22062003/game/states"

var currentGameProcess *GameProcessManager

var availableSteps = map[string]func(states.AreaState) states.AreaResult{
	"осмотреться": func(a states.AreaState) states.AreaResult {
		return a.LookAround()
	},
	"идти коридор": func(a states.AreaState) states.AreaResult {
		return a.GoCorridor()
	},
	"взять ключи": func(a states.AreaState) states.AreaResult {
		return a.TakeItem("ключи")
	},
	"надеть рюкзак": func(a states.AreaState) states.AreaResult {
		return a.TakeItem("рюкзак")
	},
	"идти комната": func(a states.AreaState) states.AreaResult {
		return a.GoRoom()
	},
	"взять конспекты": func(a states.AreaState) states.AreaResult {
		return a.TakeItem("конспекты")
	},
	"применить ключи дверь": func(a states.AreaState) states.AreaResult {
		return a.ApplyItem("ключи", "дверь")
	},
	"взять телефон": func(a states.AreaState) states.AreaResult {
		return a.TakeItem("телефон")
	},
	"идти кухня": func(a states.AreaState) states.AreaResult {
		return a.GoKitchen()
	},
	"идти улица": func(a states.AreaState) states.AreaResult {
		return a.GoStreet()
	},
	"применить телефон шкаф": func(a states.AreaState) states.AreaResult {
		return a.ApplyItem("телефон", "шкаф")
	},
	"применить ключи шкаф": func(a states.AreaState) states.AreaResult {
		return a.ApplyItem("ключи", "шкаф")
	},
}

type GameProcessManager struct {
	CurrentLocation states.AreaState
	User            *states.User
}

func (g *GameProcessManager) InitUser() {
	g.User = &states.User{}
	g.User.Init()
	states.StatesFactory.User = g.User
}

func (g *GameProcessManager) ManageCommand(command string) string {
	action, isOk := availableSteps[command]
	if !isOk {
		return "неизвестная команда"
	}

	result := action(g.CurrentLocation)
	if result.AreaState != nil {
		g.CurrentLocation = result.AreaState
	}
	return result.Message
}

func InitGame() {
	currentGameProcess = &GameProcessManager{}

	currentGameProcess.InitUser()

	// Сбрасываем состояния перед новой игрой
	states.StatesFactory.ResetStates()

	startArea := states.StatesFactory.GetAreaState("kitchen")
	currentGameProcess.CurrentLocation = startArea
}

func HandleCommand(command string) string {
	return currentGameProcess.ManageCommand(command)
}
