package handlers

import (
	"github.com/lazazael/GoWeb5HelloServer/pkg/config"
	"github.com/lazazael/GoWeb5HelloServer/pkg/render"
	"net/http"
)

//Repository pattern allows to swap components around the application

//Repository the repository type
type Repository struct {
	App *config.AppConfig
}

//Repo the repository used by the handlers
var Repo *Repository

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//_, _= fmt.Fprintf(w,"This is the home page.")
	render.RenderTemplate(w, "home.page.tmpl")
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

/*func Divide(w http.ResponseWriter, r *http.Request) {
	v, err := divideValues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, v))
}

func divideValues(x float32, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}*/
