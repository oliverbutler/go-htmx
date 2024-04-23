package pages

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func IndexPage() g.Node {
	return Div(
		Navbar(),
		H2(
			g.Text("My index content"),
		),
	)
}
