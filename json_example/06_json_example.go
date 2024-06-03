// json 문자열을 go 데이터로 변환 및 출력
package json_example

import (
	"encoding/json"
	"fmt"
)

// type SuperHero struct {
// 	SquadName string `json:"squadName"`
// 	HomeTown string `json:"homeTown"`
// 	Formed int `json:"formed"`
// 	SecretBase string `json:"secretBase"`
// 	Active bool `json:"active"`
// 	Members [] struct {
// 		Name string `json:"name"`
// 		Age int `json:"age"`
// 		SecretIdentity string `json:"secretIdentity"`
// 		Powers []string `json:"powers"`
// 	} `json:"members"`
// }

func JsonStringToGo() {
	superHeroes_source := `{
		"squadName": "Super hero squad",
		"homeTown": "Metro City",
		"formed": 2016,
		"secretBase": "Super tower",
		"active": true,
		"members": [
		  {
			"name": "Molecule Man",
			"age": 29,
			"secretIdentity": "Dan Jukes",
			"powers": [
				"Radiation resistance",
				"Turning tiny",
				"Radiation blast"
			]
		  },
		  {
			"name": "Madame Uppercut",
			"age": 39,
			"secretIdentity": "Jane Wilson",
			"powers": [
				"Million tonne punch",
				"Damage resistance",
				"Superhuman reflexes"
			]
		  },
		  {
			"name": "Eternal Flame",
			"age": 1000000,
			"secretIdentity": "Unknown",
			"powers": [
				"Immortality",
				"Heat Immunity",
				"Inferno",
				"Teleportation",
				"Interdimensional travel"
			]
		  }
		]
	}`
	var superHeroes SuperHero
	err := json.Unmarshal([]byte(superHeroes_source),&superHeroes) //json to 문자열
	if err!=nil {
		fmt.Println("JSON 디코딩 실패:", err)
		return
	}   
	fmt.Println(superHeroes.HomeTown)
}
