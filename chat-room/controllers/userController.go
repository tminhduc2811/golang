package controllers

import (
	"encoding/json"
	"net/http"
	"../common"
	"../data"
	"../models"
)

func Register(w http.ResponseWriter, r *http.Request)  {
	var dataResource UserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	repo.CreateUser(user)
	// Clear the hash password before response
	user.HashPassword = nil
	j, err := json.Marshal(UserResource{Data: *user})
	if err != nil {
		common.DisplayAppError(
				w,
				err,
				"An unexpected has occured",
				500,
			)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func Login(w http.ResponseWriter, r *http.Request)  {
	var dataResource LoginResource
	var token string
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login format",
			500,
			)
		return
	}
	loginmodel := dataResource.Data
	loginUser := models.User{
		Username:	loginmodel.Username,
		Password:	loginmodel.Password,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	user, err := repo.Login(loginUser)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			401,
			)
		return
	}
	token, err = common.GenerateJWT(user.Username, "member")
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Error while generating the access token",
			500,
			)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user.HashPassword = nil
	authUser := AuthUserModel{
		User:	user,
		Token:	token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
			)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}