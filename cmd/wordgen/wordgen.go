package main

import (
	"fmt"
	"wordgen/internal/randomizer"
)

var words = []string{
	"dizlii",
	"axogaba",
	"uazo",
	"nao",
	"dubfali",
	"winao",
	"raehus",
	"ola",
	"kezucdeed",
	"xafab",
	"iwo",
	"lun",
}


func main(){
	// for _, word := range randomizer.FetchWords(10) {
	// 	fmt.Printf("%s\n", word)
	// }

	for _, word := range words {
		for i := 0; i < 5; i++ {
			fmt.Println(randomizer.ModifyWord(word))
		}
		println()
		println()
	}
}
