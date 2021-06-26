package alfred

import (
	"encoding/json"
	"fmt"
	"log"
)

type Result struct {
	Items []interface{} `json:"items"`
}

type Item struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
}

type Error struct {
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Valid    bool   `json:"valid"`
}

func SendResult(items []Item) {
	result := Result{Items: make([]interface{}, len(items))}

	for i, item := range items {
		result.Items[i] = item
	}

	send(result)
}

func SendError(err error) {
	send(Result{
		Items: []interface{}{
			Error{
				Title:    "Something went wrong!",
				SubTitle: err.Error(),
				Valid:    false,
			},
		},
	})
}

func send(data interface{}) {
	bytes, err := json.Marshal(data)

	if err == nil {
		fmt.Println(string(bytes))

		return
	}

	log.Panicln(err)
}