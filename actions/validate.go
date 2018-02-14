package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/kteb/demo_validate_emails/models"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// ValidateToken default implementation.
func ValidateToken(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty EmailValidate
	ev := &models.EmailValidate{}

	//Getting the current user if it exist
	u := &models.User{}
	// check if the user is logged in
	if c.Value("current_user_id") != nil {
		err := tx.Find(u, c.Value("current_user_id"))
		if err != nil {
			c.Session().Set("current_user_id", nil)
			c.Flash().Add("warning", "We didn't find the session you were connected to, so please reconnect")
			return c.Redirect(302, "/signin")
		}
		//u = c.Value("current_user").(*models.User)
	}

	token := c.Param("token")

	if token == "" {
		if u.ID != uuid.Nil {
			c.Flash().Add("warning", "You're account is already validated. If you want to validate an other user, please logout first.")
			return c.Redirect(302, "/")
		}
		c.Set("email_validate", ev)
		return c.Render(200, r.HTML("validate_token/new.html"))
	}

	// To find the Token the parameter token is used.
	if err := tx.Scope(models.ByToken(token)).First(ev); err != nil {
		c.Set("email_validate", ev)
		c.Flash().Add("danger", "The token is invalid")
		return c.Render(200, r.HTML("validate_token/new.html"))
	}

	// TODO: refactor the double if user maybe u.Id insted of c.Value("current_user")
	if u.ID != uuid.Nil {
		if u.ID != ev.UserID {
			c.Flash().Add("warning", "You're trying to validate an other user, please logout first.")
			c.Logger().Warn(u.Email + " is trying to validate " + ev.UserID.String())
			return c.Redirect(302, "/")
		}
	} else {
		err := tx.Find(u, ev.UserID)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	// Validate the user
	u.Validated = true
	err := tx.Update(u)
	if err != nil {
		return errors.WithStack(err)
	}

	// Delete the validation account
	if err := tx.Destroy(ev); err != nil {
		return errors.WithStack(err)
	}

	// Finished the validation
	c.Flash().Add("success", "Your email was successfully validated!")
	if c.Value("current_user") != nil {
		return c.Redirect(302, "/")
	}
	return c.Redirect(302, "/signin")
}
