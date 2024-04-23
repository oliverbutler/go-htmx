package pages

import (
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"
)
func Navbar() g.Node {
	return Nav(
		Class("flex flex-row justify-between items-center mt-2"),
		A(
			hx.Boost("true"),
			Href("/"),
			Class("text-xl  font-bold no-underline"),
			g.Text("TheStartup 2"),
		),
	)
}

