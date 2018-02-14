package mailers

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/kteb/demo_validate_emails/models"
	"github.com/pkg/errors"
)

// SendValidateEmail send an email when a user suscribe to the Validate
func SendValidateEmail(u models.User, token string, c buffalo.Context) error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Validate Email"
	m.From = "validate-email@example.com"
	m.To = []string{u.Email}
	data := render.Data{
		"user":              u,
		"token":             token,
		"HOST":              c.Value("HOST"),
		"validateTokenPath": c.Data()["validateTokenPath"],
	}
	err := m.AddBody(r.HTML("validate_email.html"), data)
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
