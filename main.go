package game

import "github.com/Pavel22062003/game/game_process"

func initGame() {
	game_process.InitGame()
}

func handleCommand(command string) string {
	return game_process.HandleCommand(command)
}
