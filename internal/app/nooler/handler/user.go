package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/seongminnpark/nooler-server/internal/pkg/model"
	"github.com/seongminnpark/nooler-server/internal/pkg/util"
)

type UserHandler struct {
	DB *sql.DB
}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	tokenHeader := r.Header.Get("access_token")

	// Extract uuid from token.
	var token *model.Token
	if err := token.Decode(tokenHeader); err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Fetch user information.
	user := model.User{UUID: token.UUID}
	if err := user.GetUser(handler.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	util.RespondWithJSON(w, http.StatusOK, user)
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	// Cast user info from request to form object.
	var form model.SignupForm
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&form); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Chcek if user by email already exists.
	var existingUser model.User
	existingUser.Email = form.Email
	if err := existingUser.GetUserByEmail(handler.DB); err == nil {
		util.RespondWithError(w, http.StatusUnauthorized, "Email already exists")
		return
	}

	// Create new uuid for user.
	uuid, uuidErr := util.CreateUUID()
	if uuidErr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, uuidErr.Error())
		return
	}

	// Create password hash for user.
	passwordHash, err := util.SaltAndHashPassword(form.Password)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Create new user instance.
	user := model.User{UUID: uuid, Email: form.Email, PasswordHash: passwordHash}

	// Generate new token.
	token := model.Token{
		UUID: user.UUID,
		Exp:  time.Now().Add(time.Hour * 24).Unix()}

	// Encode into string.
	tokenString, encodeErr := token.Encode()
	if encodeErr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, encodeErr.Error())
		return
	}

	// Create new user.
	if err := user.CreateUser(handler.DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON := make(map[string]string)
	responseJSON["token"] = tokenString
	util.RespondWithJSON(w, http.StatusCreated, responseJSON)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	// Cast user info from request to form object.
	var form model.LoginForm
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&form); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Fetch user instance.
	var user model.User
	user.Email = form.Email
	if err := user.GetUserByEmail(handler.DB); err != nil {
		util.RespondWithError(w, http.StatusNonAuthoritativeInfo, "Wrong email")
		return
	}

	// Check if password hashes match.
	if !util.CompareHashAndPassword(user.PasswordHash, form.Password) {
		util.RespondWithError(w, http.StatusUnauthorized, "Wrong password")
		return
	}

	// Generate new token.
	token := model.Token{
		UUID: user.UUID,
		Exp:  time.Now().Add(time.Hour * 24).Unix()}

	// Encode into string.
	tokenString, err := token.Encode()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJSON := make(map[string]string)
	responseJSON["token"] = tokenString
	util.RespondWithJSON(w, http.StatusCreated, responseJSON)
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	// Validate token.
	tokenHeader := r.Header.Get("access_token")
	var token model.Token
	if err := token.Decode(tokenHeader); err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Fetch user corresponding to token.
	existingUser := model.User{UUID: token.UUID}
	if err := existingUser.GetUser(handler.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Cast user info from request to user object.
	var user model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()

	// Check if new info is valid.

	// Check if new email already exists.
	if user.Email != "" && existingUser.Email != user.Email {
		var userWithEmail model.User
		if err := userWithEmail.GetUserByEmail(handler.DB); err != nil {
			util.RespondWithError(w, http.StatusBadRequest, "Email already exists")
			return
		}
	}

	// Update user.
	if err := user.UpdateUser(handler.DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Construct and send response.
	responseJSON := make(map[string]string)
	tokenString, encodeErr := token.Encode()
	if encodeErr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, encodeErr.Error())
		return
	}
	responseJSON["token"] = tokenString
	util.RespondWithJSON(w, http.StatusOK, responseJSON)
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Validate token.
	tokenHeader := r.Header.Get("access_token")

	var token model.Token
	if err := token.Decode(tokenHeader); err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Delete
	user := model.User{UUID: token.UUID}
	if err := user.DeleteUser(handler.DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
