package handlers

import (
	"net/http"
	"udemy/pkg/config"
	"udemy/pkg/models"
	"udemy/pkg/renders"
)

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

type Repository struct {
	App *config.AppConfig
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"
	renders.RenderTemplate(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIp

	renders.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
