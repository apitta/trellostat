package board

import (
	"testing"
	"trellostat"
	"trellostat/config"
)

func TestBoard(t *testing.T) {
	t.Parallel()
	b := Board{
		Id:   "abc",
		Name: "abc",
	}
	if b.Id != "abc" || b.Name != "abc" {
		t.Error("error creating the board")
	}
}

func TestRetrieveTrelloBoards(t *testing.T) {
	config, errConf := config.ConfigFile("../trellostat.json")
	if errConf != nil {
		t.Fatal(errConf)
	}
	credential, errCred := trellostat.TrelloCredentials(config.Key, config.Token)
	if errCred != nil {
		t.Fatal(errCred)
	}
	t.Log(credential)
	boards, errBoard := RetrieveTrelloBoards(credential)
	if errBoard != nil {
		t.Error(errBoard)
	}
	t.Log(len(boards))
	if len(boards) != 19 {
		t.Error("number of boards different than expected")
	}
}
