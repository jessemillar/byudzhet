package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Email    string
	ClientID string
	Picture  string
	UserID   string
	Nickname string
}

func ValidateJWT(context echo.Context) (User, error) {
	cookie, err := context.Cookie("id_token")
	if err != nil {
		return User{}, errors.New("User not authorized")
	}

	req, err := http.NewRequest("POST", "https://jessemillar.auth0.com/tokeninfo?id_token="+cookie.Value(), nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return User{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return User{}, err
	}

	user := User{}
	json.Unmarshal(body, &user)

	return user, nil
}
