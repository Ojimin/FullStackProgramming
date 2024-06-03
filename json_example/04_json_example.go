// go 언어 to json 파일로 저장
package json_example

import (
	"encoding/json"
	"fmt"
	"os"
)

// type SuperHero struct {
// 	SquadName string `json:"squadName"`
// 	HomeTown string `json:"homeTown"`
// 	Formed int `json:"formed"`
// 	SecretIdentity string `json:"secretBase"`
// 	Active bool `json:"active"`
// 	Members [] struct {
// 		Name string `json:"name"`
// 		Age int `json:"age"`
// 		SecretIdentity string `json:"secretIdentity"`
// 		Powers []string `json:"powers"`
// 	} `json:"members"`
// }

func GoToJson() {
	
	superHeroes := SuperHero {
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

		fmt.Println(superHeroes.HomeTown)
		fmt.Println(superHeroes.Active)
		fmt.Println(superHeroes.Members[1].Powers[2])
		
		jsonData, err := json.MarshalIndent(superHeroes, "", "  ") //go to json 변환

		file, err := os.Create("04_json_example.json")
		if err != nil {
			panic(err)
		} 
		defer file.Close()

		_, err = file.Write(jsonData)
		if err != nil {
			return
		}
}