package controllers

import "../models"

type (
	UserResource struct {
		Data models.User		`json:"data"`
	}

	LoginModel struct {
		Username    string	`json:"username"`
		Password 	string 	`json:"password"`
	}

	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}

	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
)

