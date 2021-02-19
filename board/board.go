package board

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"trellostat"
)

// Board structure that represents name and id from a Trello Board
type Board struct {
	Name string
	Id   string
}

var client http.Client

// RetrieveTrelloBoards uses TrelloCredential to connect on API an
// retrieves Trello Boards
func RetrieveTrelloBoards(credential trellostat.TrelloCredential) ([]Board, error) {
	request, error := http.NewRequest("GET", "https://api.trello.com/1/members/me/boards", nil)
	if error != nil {
		return []Board{}, errors.New("Error creating request")
	}
	query := request.URL.Query()

	query.Add("key", credential.Key)
	query.Add("token", credential.Token)
	request.URL.RawQuery = query.Encode()

	response, error := client.Do(request)
	if error != nil {
		return []Board{}, error
	}
	content, error := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if error != nil {
		return []Board{}, error
	}
	decoder := json.NewDecoder(strings.NewReader(string(content)))
	_, err := decoder.Token()
	if err != nil {
		fmt.Println(err)
	}
	var boards []Board
	for decoder.More() {
		var b Board
		err := decoder.Decode(&b)
		if err != nil {
			fmt.Println(err)
		}
		boards = append(boards, b)
	}
	return boards, nil
}
