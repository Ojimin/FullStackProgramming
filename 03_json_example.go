// json 파일 읽은 후 to go data로 변환
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SuperHero struct {
	SquadName string `json:"squadName"`
	HomeTown string `json:"homeTown"`
	Formed int `json:"formed"`
	SecretBase string `json:"secretBase"`
	Active bool `json:"active"`
	Members [] struct {
		Name string `json:"name"`
		Age int `json:"age"`
		SecretIdentity string `json:"secretIdentity"`
		Powers []string `json:"powers"`
	} `json:"members"`
}

func main() {
	file, err := os.Open("03_json_example.json")
	if err != nil {
		return 
	}
	defer file.Close()

	//디코딩
	var superHeroes SuperHero
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&superHeroes) //주소값을 넣어야함

	//출력
	fmt.Println(superHeroes.SquadName)
	fmt.Println(superHeroes.Active)
	fmt.Println(superHeroes.Members[1].Powers[2])
}
