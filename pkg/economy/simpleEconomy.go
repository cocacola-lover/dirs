package economy

import (
	"encoding/json"
	"os"
)

type SimpleEconomy struct {
	wants    map[string]bool
	packages map[string]string
	friends  map[string]bool
}

func (e SimpleEconomy) GetPackage(id string) (string, bool) {
	val, ok := e.packages[id]
	return val, ok
}

func (e SimpleEconomy) GetWants(id string) map[string]bool {
	return e.wants
}

func (e SimpleEconomy) Need(id string) {
	e.wants[id] = true
}

func (e SimpleEconomy) Put(id, val string) {
	e.packages[id] = val
}

func NewSimpleEconomy() SimpleEconomy {

	var friendsArr []string
	err := json.Unmarshal([]byte(os.Getenv("friends")), &friendsArr)
	if err != nil {
		return SimpleEconomy{
			wants:    map[string]bool{},
			packages: map[string]string{},
			friends:  map[string]bool{},
		}
	}

	friends := map[string]bool{}
	for _, v := range friendsArr {
		friends[v] = true
	}

	return SimpleEconomy{
		wants:    map[string]bool{},
		packages: map[string]string{},
		friends:  friends,
	}
}
