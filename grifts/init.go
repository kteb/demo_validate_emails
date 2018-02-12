package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kteb/demo_validate_emails/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
