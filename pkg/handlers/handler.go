package handlers

import (
	"net/http"

	"github.com/thiruthanikaiarasu/udemy-go/bookings/pkg/config"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/pkg/models"
	"github.com/thiruthanikaiarasu/udemy-go/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig 
}

// NewRepo creates a new repository 
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	} 
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {   

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again" 

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}


// Reservation renders the make a reservation page and display from 
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.gohtml", &models.TemplateData{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.gohtml", &models.TemplateData{})
}

// Majors renders the room page 
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.gohtml", &models.TemplateData{})
}

// Majors renders the room page 
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.gohtml", &models.TemplateData{})
}

// Majors renders the room page 
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.gohtml", &models.TemplateData{})
}