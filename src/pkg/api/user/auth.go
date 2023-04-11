package api_user

import (
	"log"
	"net/http"
	"strings"

	db_user "go-sql/pkg/db/user"
	"go-sql/pkg/helper"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	email, found := c.Params.Get("email")
	if !found {
		msg := "api.user.SignIn: Missing email"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, "auth/missing-email")
		return
	}
	email = strings.Replace(email, "%40", "@", 1)

	password_hash, found := c.Params.Get("password")
	if !found {
		msg := "api.user.SignIn: Missing password"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, "auth/missing-password")
		return
	}

	user, err := db_user.GetByEmail(email)

	if user.PasswordHash != password_hash {
		log.Println("api.user.SignIn: Passwords do not match")
		c.JSON(http.StatusBadRequest, "auth/password-does-not-match")
		return
	}

	if err != "" {
		c.JSON(http.StatusBadRequest, "auth/email-not-found")
		return
	}

	c.JSON(http.StatusOK, user)
}

func SignUp(c *gin.Context) {
	name := c.DefaultPostForm("name", "nil")
	password_hash := c.DefaultPostForm("password", "nil") // todo: hash the password on the client side
	email := c.DefaultPostForm("email", "nil")

	if name == "nil" {
		log.Println("api.user.SignUp: Missing name")
		c.JSON(http.StatusBadRequest, "auth/missing-username")
		return
	}
	if password_hash == "nil" {
		log.Println("api.user.SignUp: Missing password")
		c.JSON(http.StatusBadRequest, "auth/missing-password")
		return
	}
	if email == "nil" {
		log.Println("api.user.SignUp: Missing email")
		c.JSON(http.StatusBadRequest, "auth/missing-email")
		return
	}

	user := helper.User{Name: name, PasswordHash: string(password_hash), Email: email}

	msg := db_user.Create(user)

	c.JSON(0, msg)
}
