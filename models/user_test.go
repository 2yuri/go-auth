package models_test

import (
	"go-auth/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
	{
		user := models.NewUser("name", "mail", "password")
		require.NotNil(t, user)

		withoutName := models.NewUser("", "mail", "password")
		require.Nil(t, withoutName)

		withoutMail := models.NewUser("name", "", "password")
		require.Nil(t, withoutMail)

		withoutPass := models.NewUser("name", "mail", "")
		require.Nil(t, withoutPass)

		withoutAll := models.NewUser("", "", "")
		require.Nil(t, withoutAll)
	}

	{
		user := models.NewUserWithId(1, "yuri", "mail", "password")
		require.NotNil(t, user)

		withZeroID := models.NewUserWithId(0, "name", "mail", "")
		require.Nil(t, withZeroID)

		withoutName := models.NewUserWithId(1, "", "mail", "password")
		require.Nil(t, withoutName)

		withoutMail := models.NewUserWithId(1, "name", "", "password")
		require.Nil(t, withoutMail)

		withoutPass := models.NewUserWithId(1, "name", "mail", "")
		require.Nil(t, withoutPass)

		withoutAll := models.NewUserWithId(0, "", "", "")
		require.Nil(t, withoutAll)
	}
}
