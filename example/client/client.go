package client

import (
	"fmt"

	"encoding/json"

	"github.com/codelingo/tenets/example/common"
	"github.com/codelingo/tenets/example/server"
	"github.com/juju/errors"
)

func GetData() {
	myClient := &common.Client{&common.Server{}}

	response, err := Request(myClient)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("The server sent us:")
	fmt.Println(response)
}

func Request(c *common.Client) (string, error) {
	jsonResponse, err := server.MakeJSONRequest(c.Connection)
	if err != nil {
		return "", errors.Trace(err)
	}

	var cs common.ServerData
	err = json.Unmarshal(jsonResponse, &cs)
	if err != nil {
		return "", errors.Trace(err)
	}
	return cs.A, nil
}
