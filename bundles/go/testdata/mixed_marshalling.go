package testdata

import (
	"fmt"

	"encoding/json"

	"github.com/juju/errors"
)

// Pretend this is defined in the client
type ClientData struct {
	A string
}

func SendData() error {
	myServer := &common.Server{}

	data, err := MakeJSONRequest(myServer)
	if err != nil {
		return errors.Trace(err)
	}

	fmt.Printf("sending data %v", data)
	return nil
}

func MakeJSONRequest(s *common.Server) ([]byte, error) {
	cliStr := common.ClientData{
		A: "a fixed value",
	}

	data, err := json.Marshal(cliStr)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return data, nil
}
