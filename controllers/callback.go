package controllers

import (
	// Don't forget this first import or nothing will work
	_ "crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func (cg *ControllerGroup) CallbackHandler(c echo.Context) error {
	domain := "jessemillar.auth0.com"

	// Instantiating the OAuth2 package to exchange the Code for a Token
	conf := &oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  "http://byudzhet.jessemillar.com/callback",
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

	// Getting now the User information
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

	// Saving the information to the session.
	// We're using https://github.com/astaxie/beego/tree/master/session
	// The GlobalSessions variable is initialized in another file
	// Check https://github.com/auth0/auth0-golang/blob/master/examples/regular-web-app/app/app.go
	// session, _ := app.GlobalSessions.SessionStart(c, c.Request())
	// defer session.SessionRelease(c)

	// session.Set("id_token", token.Extra("id_token"))
	// session.Set("access_token", token.AccessToken)
	// session.Set("profile", profile)

	// Redirect to logged in page
	c.Redirect(http.StatusMovedPermanently, "/health")

	return c.String(http.StatusOK, "DONE")
}
