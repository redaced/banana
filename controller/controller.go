package controller
import (
	"net/http"
	"html/template"
	// "os"
	// "strings"
)
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("view/*.gohtml"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}
func Get(w http.ResponseWriter, r *http.Request) {
	error := tpl.ExecuteTemplate(w, "welcome.gohtml", nil)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
}