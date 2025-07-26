package states

var DefaultRoomState = RoomState{
	BaseState: BaseState{
		AvailableItems: []string{
			"ключи",
			"конспекты",
			"рюкзак",
		},
	},
}
var DefaultCorridor = CorridorState{
	BaseState: BaseState{
		AvailableItems: []string{
			"ключи",
			"конспекты",
			"рюкзак",
		},
	},
	IsDoorOpen: false,
}

var DefaultKitchen = KitchenState{
	BaseState: BaseState{
		AvailableItems: []string{
			"чай",
		},
	},
}
