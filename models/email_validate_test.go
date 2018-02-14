package models_test

import (
	"github.com/kteb/demo_validate_emails/models"

	uuid "github.com/satori/go.uuid"
)

func (ms *ModelSuite) Test_EmailValidate_Create() {

	count, err := ms.DB.Count("email_validates")
	ms.NoError(err)
	ms.Equal(0, count)

	uuid, _ := uuid.NewV1()
	ev := &models.EmailValidate{
		UserID: uuid,
		Token:  "randomTokenGenerated",
	}
	ms.Zero(ev.TokenHash)

	verrs, err := ev.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(ev.TokenHash)
	ms.Zero(ev.Token)

	count, err = ms.DB.Count("email_validates")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_EmailValidate_ValidationErrors() {

	count, err := ms.DB.Count("email_validates")
	ms.NoError(err)
	ms.Equal(0, count)

	ev := &models.EmailValidate{
		Token: "randomTokenGenerated",
	}
	ms.Zero(ev.TokenHash)

	verrs, err := ev.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("email_validates")
	ms.NoError(err)
	ms.Equal(0, count)
}
