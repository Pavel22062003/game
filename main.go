package game

import (
	"fmt"
	"strings"
)

type Player struct {
	location    string
	inventory   []string
	hasBackpack bool
}

type Item struct {
	name     string
	location string
}

type Room struct {
	name        string
	description string
	items       []Item
	exits       map[string]string
	actions     map[string]func(*Player, []string) string
}

type Game struct {
	player   *Player
	rooms    map[string]*Room
	doorOpen bool
}

var game *Game

func initGame() {
	game = &Game{
		player: &Player{
			location:    "кухня",
			inventory:   []string{},
			hasBackpack: false,
		},
		rooms:    make(map[string]*Room),
		doorOpen: false,
	}

	kitchen := &Room{
		name:        "кухня",
		description: "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
		items: []Item{
			{name: "чай", location: "стол"},
		},
		exits: map[string]string{
			"коридор": "коридор",
		},
		actions: make(map[string]func(*Player, []string) string),
	}

	corridor := &Room{
		name:        "коридор",
		description: "ничего интересного. можно пройти - кухня, комната, улица",
		items:       []Item{},
		exits: map[string]string{
			"кухня":   "кухня",
			"комната": "комната",
			"улица":   "улица",
		},
		actions: make(map[string]func(*Player, []string) string),
	}

	room := &Room{
		name:        "комната",
		description: "ты в своей комнате. можно пройти - коридор",
		items: []Item{
			{name: "ключи", location: "стол"},
			{name: "конспекты", location: "стол"},
			{name: "рюкзак", location: "стул"},
		},
		exits: map[string]string{
			"коридор": "коридор",
		},
		actions: make(map[string]func(*Player, []string) string),
	}

	street := &Room{
		name:        "улица",
		description: "на улице весна. можно пройти - домой",
		items:       []Item{},
		exits: map[string]string{
			"домой": "коридор",
		},
		actions: make(map[string]func(*Player, []string) string),
	}

	game.rooms["кухня"] = kitchen
	game.rooms["коридор"] = corridor
	game.rooms["комната"] = room
	game.rooms["улица"] = street

	setupRoomActions()
}

func setupRoomActions() {
	game.rooms["кухня"].actions["осмотреться"] = func(p *Player, args []string) string {
		if p.hasBackpack {
			return "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
		}
		return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
	}

	game.rooms["комната"].actions["осмотреться"] = func(p *Player, args []string) string {
		items := game.rooms["комната"].items
		if len(items) == 0 {
			return "пустая комната. можно пройти - коридор"
		}

		tableItems := []string{}
		chairItems := []string{}

		for _, item := range items {
			if item.location == "стол" {
				tableItems = append(tableItems, item.name)
			} else if item.location == "стул" {
				chairItems = append(chairItems, item.name)
			}
		}

		description := ""
		if len(tableItems) > 0 {
			description += "на столе: " + strings.Join(tableItems, ", ")
		}
		if len(chairItems) > 0 {
			if len(tableItems) > 0 {
				description += ", "
			}
			description += "на стуле: " + strings.Join(chairItems, ", ")
		}
		description += ". можно пройти - коридор"

		return description
	}

	game.rooms["коридор"].actions["осмотреться"] = func(p *Player, args []string) string {
		return "ничего интересного. можно пройти - кухня, комната, улица"
	}

	game.rooms["улица"].actions["осмотреться"] = func(p *Player, args []string) string {
		return "на улице весна. можно пройти - домой"
	}
}

func handleCommand(command string) string {
	parts := strings.Split(command, " ")
	if len(parts) == 0 {
		return "неизвестная команда"
	}

	action := parts[0]
	args := parts[1:]

	switch action {
	case "осмотреться":
		return handleLookAround(args)
	case "идти":
		return handleGo(args)
	case "взять":
		return handleTake(args)
	case "надеть":
		return handleWear(args)
	case "применить":
		return handleApply(args)
	default:
		return "неизвестная команда"
	}
}

func handleLookAround(args []string) string {
	currentRoom := game.rooms[game.player.location]
	if action, exists := currentRoom.actions["осмотреться"]; exists {
		return action(game.player, args)
	}
	return currentRoom.description
}

func handleGo(args []string) string {
	if len(args) == 0 {
		return "куда идти?"
	}

	target := args[0]
	currentRoom := game.rooms[game.player.location]

	if exit, exists := currentRoom.exits[target]; exists {
		if target == "улица" && !game.doorOpen {
			return "дверь закрыта"
		}

		game.player.location = exit
		newRoom := game.rooms[exit]

		if exit == "кухня" {
			if game.player.hasBackpack {
				return "кухня, ничего интересного. можно пройти - коридор"
			}
		}

		return newRoom.description
	}

	return "нет пути в " + target
}

func handleTake(args []string) string {
	if len(args) == 0 {
		return "что взять?"
	}

	itemName := args[0]
	currentRoom := game.rooms[game.player.location]

	if !game.player.hasBackpack {
		return "некуда класть"
	}

	for i, item := range currentRoom.items {
		if item.name == itemName {
			currentRoom.items = append(currentRoom.items[:i], currentRoom.items[i+1:]...)
			game.player.inventory = append(game.player.inventory, itemName)
			return "предмет добавлен в инвентарь: " + itemName
		}
	}

	return "нет такого"
}

func handleWear(args []string) string {
	if len(args) == 0 {
		return "что надеть?"
	}

	itemName := args[0]
	currentRoom := game.rooms[game.player.location]

	for i, item := range currentRoom.items {
		if item.name == itemName && itemName == "рюкзак" {
			currentRoom.items = append(currentRoom.items[:i], currentRoom.items[i+1:]...)
			game.player.hasBackpack = true
			return "вы надели: " + itemName
		}
	}

	return "нет такого"
}

func handleApply(args []string) string {
	if len(args) < 2 {
		return "что к чему применить?"
	}

	itemName := args[0]
	target := args[1]

	hasItem := false
	for _, item := range game.player.inventory {
		if item == itemName {
			hasItem = true
			break
		}
	}

	if !hasItem {
		return "нет предмета в инвентаре - " + itemName
	}

	if itemName == "ключи" && target == "дверь" {
		game.doorOpen = true
		return "дверь открыта"
	}

	return "не к чему применить"
}

func main() {
	initGame()
	fmt.Println("Текстовая игра запущена!")
	fmt.Println("Доступные команды: осмотреться, идти <направление>, взять <предмет>, надеть <предмет>, применить <предмет> <цель>")
}
