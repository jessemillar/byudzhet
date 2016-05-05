package controllers

// import (
// 	// Don't forget this first import or nothing will work
// 	_ "crypto/sha512"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
//
// 	"github.com/jessemillar/byudzhet/packages/app"
// 	"github.com/labstack/echo"
// 	"golang.org/x/oauth2"
// )
//
// func CallbackHandler(c echo.Context) {
// 	domain := "jessemillar.auth0.com"
//
// 	// Instantiating the OAuth2 package to exchange the Code for a Token
// 	conf := &oauth2.Config{
// 		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
// 		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
// 		RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
// 		Scopes:       []string{"openid", "name", "email", "nickname"},
// 		Endpoint: oauth2.Endpoint{
// 			AuthURL:  "https://" + domain + "/authorize",
// 			TokenURL: "https://" + domain + "/oauth/token",
// 		},
// 	}
//
// 	// Getting the Code that we got from Auth0
// 	code := r.URL.Query().Get("code")
//
// 	// Exchanging the code for a token
// 	token, err := conf.Exchange(oauth2.NoContext, code)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	// Getting now the User information
// 	client := conf.Client(oauth2.NoContext, token)
// 	resp, err := client.Get("https://" + domain + "/userinfo")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	// Reading the body
// 	raw, err := ioutil.ReadAll(resp.Body)
// 	defer resp.Body.Close()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	// Unmarshalling the JSON of the Profile
// 	var profile map[string]interface{}
// 	if err := json.Unmarshal(raw, &profile); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	// Saving the information to the session.
// 	// We're using https://github.com/astaxie/beego/tree/master/session
// 	// The GlobalSessions variable is initialized in another file
// 	// Check https://github.com/auth0/auth0-golang/blob/master/examples/regular-web-app/app/app.go
// 	session, _ := app.GlobalSessions.SessionStart(w, r)
// 	defer session.SessionRelease(w)
//
// 	session.Set("id_token", token.Extra("id_token"))
// 	session.Set("access_token", token.AccessToken)
// 	session.Set("profile", profile)
//
// 	// Redirect to logged in page
// 	http.Redirect(w, r, "/user", http.StatusMovedPermanently)
// }
