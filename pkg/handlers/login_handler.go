package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ashmilhussain/GO-RestWithUnitTest/pkg/auth"
	"github.com/ashmilhussain/GO-RestWithUnitTest/pkg/models"
	"github.com/ashmilhussain/GO-RestWithUnitTest/pkg/responses"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	expire := time.Now().Add(20 * time.Minute)
	cookie := &http.Cookie{Name: "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  expire}
	http.SetCookie(w, cookie)
	responses.JSON(w, http.StatusOK, "Login Success")
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.UserId)
}
