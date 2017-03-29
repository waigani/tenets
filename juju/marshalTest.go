package main

import (
	"fmt"

	"encoding/json"

	"github.com/juju/errors"
)

type server struct{}

type client struct {
	connection server
}

type serverStruct struct {
	A string `json:"value"`
}

type clientStruct struct {
	B string `json:"value"`
}

func main() {
	myServer := server{}
	myClient := client{myServer}

	response, err := myClient.Request()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("The server sent us:")
	fmt.Println(response)
}

func (c *client) Request() (*clientStruct, error) {
	jsonResponse, err := c.connection.MakeJSONRequest()
	if err != nil {
		return nil, errors.Trace(err)
	}

	var cs clientStruct
	err = json.Unmarshal(jsonResponse, &cs)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &cs, nil
}

func (s *server) MakeJSONRequest() ([]byte, error) {
	cliStr := clientStruct{
		B: "a fixed value",
	}

	data, err := json.Marshal(cliStr)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return data, nil
}
