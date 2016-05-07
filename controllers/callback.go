package controllers

import (
	// Don't forget this first import or nothing will work
	_ "crypto/sha512"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func (cg *ControllerGroup) CallbackHandler(c echo.Context) error {
	domain := "jessemillar.auth0.com"

	// Instantiating the OAuth2 package to exchange the Code for a Token
	conf := &oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8000/callback",
		Scopes:       []string{"openid", "name", "email", "nickname"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	// Getting the Code that we got from Auth0
	code := c.QueryParam("code")

	// Exchanging the code for a token
	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Getting the user information
	client := conf.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://" + domain + "/userinfo")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Reading the body
	raw, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Unmarshalling the JSON of the Profile
	var profile map[string]interface{}
	if err := json.Unmarshal(raw, &profile); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// session, err := helpers.Store.Get(c.Request().(*standard.Request).Request, "session-name")
	// if err != nil {
	// 	c.String(http.StatusInternalServerError, err.Error())
	// }
	//
	// session.Values["id_token"] = token.Extra("id_token")
	// session.Values["access_token"] = token.AccessToken
	// session.Values["profile"] = profile
	//
	// // Save it before we write to the response/return from the handler.
	// session.Save(c.Request().(*standard.Request).Request, c.Response().(*standard.Response).ResponseWriter)

	// stuff := session.Get("profile")
	fmt.Printf("Token: %+v\n", token)

	helpers.MakeCookie(c, "id_token", token.Extra("id_token").(string))

	// Redirect to logged in page
	c.Redirect(http.StatusMovedPermanently, "/health")

	return c.String(http.StatusOK, "Callback finished") // We'll never actually hit this...?
}
