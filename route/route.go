package route

import (
	// "banana/view"
	"github.com/redaced/banana/controller"

	// "fmt"
	// "log"
	// "net/http"
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
	routes []*route
	// Options *Options
}

// type Options struct {
// 	Model        *model.Model
// 	View         view.View
// }

type route struct {
	pattern string   // url pattern e.g /home
	methods []string // http methods e.g GET, POST etc
	// ctrl    string   // the name of the controller
	// fn      string   // the name of the controller's method to be executed
}

var r = mux.NewRouter()

func init() {
	r.HandleFunc("/", controller.Index)
}

func Resource(Url string) {
	r.HandleFunc(Url, controller.Index)
}

// func Get(Url string, cont func(w http.ResponseWriter, r *http.Request)){
// 	r.HandleFunc(Url, cont).Methods("GET")
// }

func Invoker() *mux.Router {
	return r
}
