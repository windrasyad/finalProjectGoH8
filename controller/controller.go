// Package controller loads the routes for each of the controllers.
package controller

import (
	"blueprint/blueprint/controller/article"
	"blueprint/blueprint/controller/contact"

	"blueprint/blueprint/controller/debug"

	"blueprint/blueprint/controller/about"
	"blueprint/blueprint/controller/home"
	"blueprint/blueprint/controller/login"
	"blueprint/blueprint/controller/notepad"
	"blueprint/blueprint/controller/register"
	"blueprint/blueprint/controller/static"
	"blueprint/blueprint/controller/status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	notepad.Load()
	contact.Load()
	article.Load()
}
