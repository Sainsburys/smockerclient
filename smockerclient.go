package smockerclient

import "fmt"

type Instance struct {
	host string
	port int
}

func DefaultInstance() Instance {
	return Instance{
		host: "localhost",
		port: 8081,
	}
}

func (i Instance) GetURL() string {
	return fmt.Sprintf("http://%s:%d", i.host, i.port)
}
