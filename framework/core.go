package framework

import "net/http"

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	// TODO
}
