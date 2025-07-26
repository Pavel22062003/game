package states

import "fmt"

type User struct {
	backpack bool
	items    map[string]bool
}

func (user *User) AddItem(item string) string {
	if !user.backpack {
		return "некуда класть"
	}
	user.items[item] = true
	return fmt.Sprintf("предмет добавлен в инвентарь: %s", item)
}

func (user *User) SetBackpack() string {
	user.backpack = true
	return "вы надели: рюкзак"

}

func (user *User) Init() {
	user.items = make(map[string]bool)
}

func (user *User) GetItem(item string) bool {
	if !user.backpack {
		return false
	}
	return user.items[item]
}
