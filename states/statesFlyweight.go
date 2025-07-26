package states

type AreaStateFactory struct {
	areaStates map[string]AreaState
	User       *User
}

func NewAreaStateFactory() *AreaStateFactory {
	return &AreaStateFactory{
		areaStates: make(map[string]AreaState),
	}
}

func (f *AreaStateFactory) GetAreaState(key string) AreaState {
	if state, ok := f.areaStates[key]; ok {
		return state
	}
	var state AreaState
	switch key {
	case "room":
		roomState := DefaultRoomState
		roomState.AvailableItems = make([]string, len(DefaultRoomState.AvailableItems))
		copy(roomState.AvailableItems, DefaultRoomState.AvailableItems)
		roomState.User = f.User
		state = &roomState
	case "corridor":
		corridorState := DefaultCorridor
		corridorState.AvailableItems = make([]string, len(DefaultCorridor.AvailableItems))
		copy(corridorState.AvailableItems, DefaultCorridor.AvailableItems)
		corridorState.User = f.User
		state = &corridorState
	case "kitchen":
		kitchenState := DefaultKitchen
		kitchenState.AvailableItems = make([]string, len(DefaultKitchen.AvailableItems))
		copy(kitchenState.AvailableItems, DefaultKitchen.AvailableItems)
		kitchenState.User = f.User
		state = &kitchenState
	default:
		return nil
	}
	f.areaStates[key] = state
	return state
}

func (f *AreaStateFactory) ResetStates() {
	f.areaStates = make(map[string]AreaState)
}

var StatesFactory = NewAreaStateFactory()
