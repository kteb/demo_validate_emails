package models

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/markbates/pop"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection
var hmacl HMAC

func init() {
	hmacl = NewHMAC(envy.Get("HMACKEY", ""))
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}

func ByToken(token string) pop.ScopeFunc {
	return func(q *pop.Query) *pop.Query {
		token_hash := hmacl.hash(token)
		return q.Where("token_hash = ?", token_hash)
	}
}
