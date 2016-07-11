package handlers

import (
	// Don't forget this first import or nothing will work
	_ "crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jessemillar/byudzhet/helpers"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func (handlerGroup *HandlerGroup) CallbackHandler(context echo.Context) error {
	domain := "jessemillar.auth0.com"

	// Instantiating the OAuth2 package to exchange the Code for a Token
	conf := &oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("AUTH0_CALLBACK"),
		Scopes:       []string{"openid", "name", "email", "nickname"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	// Getting the Code that we got from Auth0
	code := context.QueryParam("code")

	// Exchanging the code for a token
	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}

	// Getting the user information
	client := conf.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://" + domain + "/userinfo")
	if err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}

	// Reading the body
	raw, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}

	// Unmarshal the JSON of the Auth0 profile
	var profile map[string]interface{}
	if err := json.Unmarshal(raw, &profile); err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}

	helpers.MakeCookie(context, "id_token", token.Extra("id_token").(string))

	// Redirect to logged in page
	context.Redirect(http.StatusMovedPermanently, "/frontend")

	return context.String(http.StatusOK, "Callback finished") // We'll never actually hit this...?
}
