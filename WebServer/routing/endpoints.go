package routing

import (
	"net/http"
	"text/template"
)

// HomePage contains the index.
func (route *Router) HomePage(w http.ResponseWriter, r *http.Request) {
	type MetaHome struct {
		Name    string
		Address string
		Port    int
		Page    string
	}
	templates, err := template.ParseFiles(
		"templates/sidebar.html",
		"templates/base.html",
		"templates/index.html",
		"templates/modal.html",
	)
	if err != nil {
		route.logger.Error(err.Error())
	}
	templates.ExecuteTemplate(w, "base", MetaHome{Name: route.projectTitle, Address: route.apiIP, Port: route.apiPort, Page: "home"})
	route.logger.Debug("An Home Request was made.")
}

// NotificationPage contains the template event page.
func (route *Router) NotificationPage(w http.ResponseWriter, r *http.Request) {
	type MetaEvent struct {
		Name    string
		Address string
		Port    int
		Page    string
	}
	templates, err := template.ParseFiles(
		"templates/sidebar.html",
		"templates/base.html",
		"templates/notifications.html",
		"templates/modal.html",
	)
	if err != nil {
		route.logger.Error(err.Error())
	}
	templates.ExecuteTemplate(w, "base", MetaEvent{Name: route.projectTitle, Address: route.apiIP, Port: route.apiPort, Page: "notifications"})
	route.logger.Debug("An Notification Request was made.")
}

// TablesPage contains the template event page.
func (route *Router) TablesPage(w http.ResponseWriter, r *http.Request) {
	type MetaTables struct {
		Name    string
		Address string
		Port    int
		Page    string
	}
	templates, err := template.ParseFiles(
		"templates/sidebar.html",
		"templates/base.html",
		"templates/tables.html",
		"templates/modal.html",
	)
	if err != nil {
		route.logger.Error(err.Error())
	}
	templates.ExecuteTemplate(w, "base", MetaTables{Name: route.projectTitle, Address: route.apiIP, Port: route.apiPort, Page: "tables"})
	route.logger.Debug("An Tables Request was made.")
}

// TypographyPage contains the template event page.
func (route *Router) TypographyPage(w http.ResponseWriter, r *http.Request) {
	type MetaTypo struct {
		Name    string
		Address string
		Port    int
		Page    string
	}
	templates, err := template.ParseFiles(
		"templates/sidebar.html",
		"templates/base.html",
		"templates/typography.html",
		"templates/modal.html",
	)
	if err != nil {
		route.logger.Error(err.Error())
	}
	templates.ExecuteTemplate(w, "base", MetaTypo{Name: route.projectTitle, Address: route.apiIP, Port: route.apiPort, Page: "typography"})
	route.logger.Debug("An Typography Request was made.")
}
