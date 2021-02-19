package trellostat_test

import (
	"os"
	"testing"
	"trellostat"
	"trellostat/board"
	"trellostat/config"
)

func TestDefineCredential(t *testing.T) {
	t.Parallel()
	credential, error := trellostat.TrelloCredentials("a", "b")
	if error != nil {
		t.Fatal("It hasn't been possible to create the credential", error)
	}
	t.Log("Credential created successfully", credential)
}

func TestRetrieveTrelloBoards(t *testing.T) {
	t.Parallel()
	t.Log(os.Executable())
	c, err := config.ConfigFile("trellostat.json")
	if err != nil {
		t.Error(err)
	}
	credential, error := trellostat.TrelloCredentials(c.Key, c.Token)
	if error != nil {
		t.Fatal("Fatal error trying to define the credential", error)
	}

	boards, error := board.RetrieveTrelloBoards(credential)
	if error != nil {
		t.Fatal("Error retrieving trello boards", error)
	}
	t.Log("Boards retrieved successfully", boards)
}
