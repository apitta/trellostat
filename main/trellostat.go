package main

import (
	"fmt"
	"os"
	"trellostat"
	"trellostat/board"
	"trellostat/config"
)

func main() {
	config, errConf := config.ConfigFile("trellostat.json")
	if errConf != nil {
		fmt.Println(errConf)
		os.Exit(2)
	}
	credential, err := trellostat.TrelloCredentials(config.Key, config.Token)
	if err != nil {
		fmt.Println("error initializing credential")
		os.Exit(3)
	}
	boards, err := board.RetrieveTrelloBoards(credential)
	if err != nil {
		fmt.Println("error retrieving boards")
		fmt.Println(err)
	}
	for i, b := range boards {
		fmt.Printf("%2d '%s' (%s)\n", i+1, b.Name, b.Id)
	}
}
