package nooler

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gorilla/mux"
// 	"github.com/seongminnpark/nooler-server/internal/pkg/model"
// 	"github.com/seongminnpark/nooler-server/internal/pkg/util"
// )

// func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
// 	tokenHeader := r.Header.Get("access_token")

// 	// Extract uuid from token.
// 	var token model.Token
// 	if token, tokenErr = token.Decode(tokenHeader); tokenErr != nil {
// 		respondWithError(w, http.StatusUnauthorized, "Invalid token")
// 		return
// 	}

// 	// Fetch user information.
// 	user := model.User{UUID: token.UUID, Token: tokenHeader}
// 	if err := user.GetUser(app.DB); err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			respondWithError(w, http.StatusNotFound, "User not found")
// 		default:
// 			respondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, user)
// }

// func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {
// 	count, _ := strconv.Atoi(r.FormValue("count"))
// 	start, _ := strconv.Atoi(r.FormValue("start"))

// 	if count > 10 || count < 1 {
// 		count = 10
// 	}

// 	if start < 0 {
// 		start = 0
// 	}

// 	users, err := model.GetUsers(app.DB, start, count)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, users)
// }

// func (app *App) createUser(w http.ResponseWriter, r *http.Request) {

// 	// Cast user info from request to user object.
// 	var user model.User
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&user); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	// Extract parameters.
// 	email := r.FormValue("email")
// 	password := r.FormValue("password")

// 	// Create new uuid for user.
// 	var newUUID string
// 	newUUID, uuidErr := util.CreateUUID()
// 	if uuidErr != nil {
// 		respondWithError(w, http.StatusInternalServerError, uuidErr.Error())
// 		return
// 	}
// 	user.UUID = newUUID

// 	// Generate new token.
// 	token = model.Token{
// 		UUID: user.UUID,
// 		Exp:  time.Now().Add(time.Hour * 24).Unix()}

// 	// Encode into string.
// 	tokenString, encodeErr := token.Encode("secret", claims)
// 	if encodeErr != nil {
// 		respondWithError(w, http.StatusInternalServerError, encodeErr.Error())
// 		return
// 	}
// 	user.Token = tokenString

// 	if err := user.CreateUser(app.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	responseJSON := make(map[string]string)
// 	responseJSON["token"] = user.Token
// 	respondWithJSON(w, http.StatusCreated, responseJSON)
// }

// func (app *App) updateUser(w http.ResponseWriter, r *http.Request) {

// 	// Validate token.
// 	tokenHeader := r.Header.Get("access_token")
// 	var token model.Token
// 	var tokenErr error
// 	if token, tokenErr = token.Decode(tokenHeader); tokenErr != nil {
// 		respondWithError(w, http.StatusUnauthorized, "Invalid token")
// 		return
// 	}

// 	// Extract parameters.
// 	email := r.FormValue("email")

// 	// Cast user info from request to user object.
// 	var user model.User
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&user); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	user.uuid = token.uuid

// 	if err := user.UpdateUser(app.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	responseJSON := make(map[string]string)
// 	responseJSON["token"] = token.Encode()
// 	respondWithJSON(w, http.StatusOK, responseJSON)
// }

// func (app *App) deleteUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
// 		return
// 	}
// 	user := model.User{ID: id}
// 	if err := user.DeleteUser(app.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
// }
