package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer 返回注册路由的negroni
func NewServer() *negroni.Negroni {
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	mx := mux.NewRouter()

	formatter := render.New(render.Options{IndentJSON: true})

	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
}
