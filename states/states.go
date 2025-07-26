package states

import "fmt"

type AreaResult struct {
	AreaState AreaState
	Message   string
}
type BaseState struct {
	AvailableItems []string
	User           *User
}

func (r *BaseState) TakeItem(item string) AreaResult {
	if !r.User.backpack && item != "рюкзак" {
		return AreaResult{
			Message: "некуда класть",
		}
	}
	found := false
	for i, availableItem := range r.AvailableItems {
		if availableItem == item {
			found = true
			r.AvailableItems = append(r.AvailableItems[:i], r.AvailableItems[i+1:]...)
			break
		}
	}

	if !found {
		return AreaResult{
			Message: "нет такого",
		}
	}

	if item == "рюкзак" {
		message := r.User.SetBackpack()
		return AreaResult{
			Message: message,
		}
	}

	r.User.AddItem(item)
	return AreaResult{
		Message: r.User.AddItem(item),
	}
}

type AreaState interface {
	GoRoom() AreaResult
	GoCorridor() AreaResult
	GoKitchen() AreaResult
	GoStreet() AreaResult
	GetAvailableAreas() []string
	LookAround() AreaResult
	TakeItem(string) AreaResult
	SetUser(user *User)
	ApplyItem(string, string) AreaResult
}
type RoomState struct {
	BaseState
}

func (r *RoomState) GoRoom() AreaResult {
	return AreaResult{
		Message: "нет пути в комната",
	}
}

func (r *RoomState) SetUser(user *User) {
	r.User = user
}

func (r *RoomState) ApplyItem(item string, thing string) AreaResult {
	return AreaResult{}
}

func (r *RoomState) LookAround() AreaResult {
	availableItems := r.GetAvailableAreas()
	var message string

	if len(availableItems) == 0 {
		message = "пустая комната"
	} else if len(availableItems) == 3 {
		message = fmt.Sprintf("на столе: %s, %s, на стуле: %s",
			availableItems[0], availableItems[1], availableItems[2])
	} else if len(availableItems) == 2 {
		message = fmt.Sprintf("на столе: %s, %s",
			availableItems[0], availableItems[1])
	} else if len(availableItems) == 1 {
		message = fmt.Sprintf("на столе: %s", availableItems[0])
	}

	return AreaResult{
		Message: message + ". можно пройти - коридор",
	}
}

func (r *RoomState) GoCorridor() AreaResult {
	corridor := StatesFactory.GetAreaState("corridor")
	return AreaResult{
		Message:   "ничего интересного. можно пройти - кухня, комната, улица",
		AreaState: corridor,
	}
}
func (r *RoomState) GoKitchen() AreaResult {
	return AreaResult{
		Message: "нет пути в кухня",
	}
}
func (r *RoomState) GoStreet() AreaResult {
	return AreaResult{
		Message: "нет пути в кухня",
	}
}
func (r *RoomState) GetAvailableAreas() []string {
	return r.AvailableItems
}

type CorridorState struct {
	BaseState
	IsDoorOpen bool
}

func (r *CorridorState) SetUser(user *User) {
	r.User = user
}

func (r *CorridorState) TakeItem(item string) AreaResult {
	return AreaResult{}
}

func (r *CorridorState) GoRoom() AreaResult {
	room := StatesFactory.GetAreaState("room")
	return AreaResult{
		Message:   "ты в своей комнате. можно пройти - коридор",
		AreaState: room,
	}
}

func (r *CorridorState) LookAround() AreaResult {
	result := AreaResult{
		Message: "Вот это коридор",
	}
	return result
}

func (r *CorridorState) GoCorridor() AreaResult {
	return AreaResult{
		Message: "Ты и так в коридоре",
	}
}
func (r *CorridorState) GetAvailableAreas() []string {
	return make([]string, 0)
}
func (r *CorridorState) GoKitchen() AreaResult {
	kitchen := StatesFactory.GetAreaState("kitchen")
	return AreaResult{
		Message:   "кухня, ничего интересного. можно пройти - коридор",
		AreaState: kitchen,
	}
}
func (r *CorridorState) GoStreet() AreaResult {
	if !r.IsDoorOpen {
		return AreaResult{
			Message: "дверь закрыта",
		}
	}
	return AreaResult{
		Message: "на улице весна. можно пройти - домой",
	}
}

func (r *CorridorState) ApplyItem(item string, thing string) AreaResult {
	hasKeys := r.User.GetItem(item)
	if !hasKeys {
		return AreaResult{
			Message: "нет предмета в инвентаре - " + item,
		}
	}
	if item == "ключи" && thing == "дверь" {
		r.IsDoorOpen = true
		return AreaResult{
			Message: "дверь открыта",
		}
	}
	return AreaResult{
		Message: "не к чему применить",
	}
}

type KitchenState struct {
	BaseState
}

func (r *KitchenState) SetUser(user *User) {
	r.User = user
}

func (r *KitchenState) TakeItem(item string) AreaResult {
	return AreaResult{}
}

func (r *KitchenState) LookAround() AreaResult {
	availableItems := r.GetAvailableAreas()
	message := "ты находишься на кухне"

	if len(availableItems) == 1 {
		message += ", на столе: " + availableItems[0]
	}
	if !r.User.backpack {
		message += ", надо собрать рюкзак и идти в универ. можно пройти - коридор"
		return AreaResult{
			Message: message,
		}
	}

	return AreaResult{
		Message: message + ", надо идти в универ. можно пройти - коридор",
	}
}

func (r *KitchenState) GoCorridor() AreaResult {
	corridor := StatesFactory.GetAreaState("corridor")
	return AreaResult{
		Message:   "ничего интересного. можно пройти - кухня, комната, улица",
		AreaState: corridor,
	}
}
func (r *KitchenState) GetAvailableAreas() []string {
	return r.AvailableItems
}

func (r *KitchenState) GoKitchen() AreaResult {
	return AreaResult{
		Message: "нет пути в кухня",
	}
}
func (r *KitchenState) GoStreet() AreaResult {
	return AreaResult{
		Message: "подошел к двери",
	}
}
func (r *KitchenState) GoRoom() AreaResult {
	room := StatesFactory.GetAreaState("room")
	return AreaResult{
		Message:   "нет пути в комната",
		AreaState: room,
	}
}

func (r *KitchenState) ApplyItem(item string, thing string) AreaResult {
	return AreaResult{}
}
