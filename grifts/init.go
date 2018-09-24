package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/testapp/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
