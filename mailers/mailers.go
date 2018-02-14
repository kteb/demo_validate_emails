package mailers

import (
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
)

var smtp mail.Sender
var r *render.Engine

func init() {

	// Pulling config from the env.
	fmt.Println("a")
	port := envy.Get("SMTP_PORT", "1025")
	host := envy.Get("SMTP_HOST", "localhost")
	user := envy.Get("SMTP_USER", "")
	password := envy.Get("SMTP_PASSWORD", "")
	fmt.Println("b")

	var err error
	smtp, err = mail.NewSMTPSender(host, port, user, password)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("c")

	r = render.New(render.Options{
		HTMLLayout:   "layout.html",
		TemplatesBox: packr.NewBox("../templates/mail"),
		Helpers: render.Helpers{
			"genURL":      genURL,
			"genURLParam": genURLParam,
		},
	})
}

func genURL(baseURL string, endURL template.HTML, help plush.HelperContext) template.HTML {
	return template.HTML(strings.Trim(baseURL, " /")) + endURL
}

func genURLParam(paramName, paramValue string, help plush.HelperContext) template.HTML {
	return template.HTML("?" + strings.Trim(paramName, " /?") + "=" + strings.Trim(paramValue, " "))
}
