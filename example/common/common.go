package common

type Server struct{}

type ServerData struct {
	A string `json:"value"`
}

type Client struct {
	Connection *Server
}

type ClientData struct {
	A string `json:"value"`
}
