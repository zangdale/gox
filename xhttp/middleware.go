package xhttp

import "net/http"

/*
	http 中间件，将 http.Handler 串起来
*/

type Middleware func(http.Handler) http.Handler

type Mws struct {
	mws []Middleware
}

func NewMiddlewares(mws ...Middleware) Mws {
	return Mws{append(([]Middleware)(nil), mws...)}
}

func (c Mws) Append(mws ...Middleware) Mws {
	return c.AddEnd(mws...)
}

func (c Mws) AddStart(mws ...Middleware) Mws {
	if len(mws) == 0 {
		return c
	}
	newCons := append(mws, c.mws...)
	return Mws{newCons}
}

func (c Mws) AddEnd(mws ...Middleware) Mws {
	if len(mws) == 0 {
		return c
	}
	newCons := append(c.mws, mws...)
	return Mws{newCons}
}

func (c Mws) Func(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := range c.mws {
		h = c.mws[len(c.mws)-1-i](h)
	}

	return h
}
