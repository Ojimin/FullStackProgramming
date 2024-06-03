// go 데이터를 json 형태의 문자열 저장 및 출력
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

func JsonToGoPrint() {

	superHeroes_source := SuperHero {
		SquadName: "Super hero squad",
		HomeTown: "Metro City",
		Formed: 2016,
		SecretBase: "Super tower",
		Active: true,
		Members: []struct {
			Name string `json:"name"`
			Age int `json:"age"`
			SecretIdentity string `json:"secretIdentity"`
			Powers []string `json:"powers"`
		}{
			{
				Name : "Molecule Man",
				Age : 29,
				SecretIdentity :  "Dan Jukes",
				Powers : []string{"Radiation resistance",
					"Turning tiny", 
					"Radiation blast",
			},
		},
		{
			Name: "Madame Uppercut",
			Age: 39,
			SecretIdentity: "Jane Wilson",
			Powers: []string{
				"Million tonne punch",
				"Damage resistance",
				"Superhuman reflexes",
			},
		},
		{
			Name: "Eternal Flame",
			Age: 1000000,
			SecretIdentity: "Unknown",
			Powers: []string{
				"Immortality",
				"Heat Immunity",
				"Inferno",
				"Teleportation",
				"Interdimensional travel",
			},	
		},
	  },
	}

	data, err := json.MarshalIndent(superHeroes_source, "", "  ") //go to json
	if err!=nil {
		fmt.Println("json 변환 실패", err)
		return
	}
	
	fmt.Println(string(data))



	
}


