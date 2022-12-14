package handles

import (
	"github.com/lcxznpy/bookings/pkg/config"
	"github.com/lcxznpy/bookings/pkg/models"
	"github.com/lcxznpy/bookings/pkg/render"
	"net/http"
)

// 一个Repository 结构体给handles使用
var Repo *Repository

// 一个repository类型的结构体
type Repository struct {
	App *config.AppConfig
}

// 创建一个新的repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// 设置repository给handles
func NewHandles(r *Repository) {
	Repo = r
}

// Home render the room page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//获取访问home的电脑的ip地址,将其传进map中
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About render the room page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello,again."

	//从会话中找到key为remote_ip对应的value并显示它
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation render the room page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// One render the room page
func (m *Repository) One(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "one.page.tmpl", &models.TemplateData{})
}

// Two render the room page
func (m *Repository) Two(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "two.page.tmpl", &models.TemplateData{})
}

// Availability render the room page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact render the room page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
