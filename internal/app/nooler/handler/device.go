package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/seongminnpark/nooler-server/internal/pkg/model"
	"github.com/seongminnpark/nooler-server/internal/pkg/util"
)

type DeviceHandler struct {
	DB *sql.DB
}

func (handler *DeviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	tokenHeader := r.Header.Get("access_token")

	// Extract uuid from token.
	var token model.Token
	if err := token.Decode(tokenHeader); err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Fetch device information.
	device := model.Device{UUID: token.UUID}
	if err := device.GetDevice(handler.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		default:
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	responseJSON := make(map[string]string)
	responseJSON["token"] = tokenHeader
	responseJSON["uuid"] = device.UUID
	responseJSON["owner"] = device.Owner
	responseJSON["location"] = device.Location
	util.RespondWithJSON(w, http.StatusOK, responseJSON)
}

func (handler *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {

	// Cast user info from request to form object.
	var form model.AddDeviceForm
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&form); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check if new info is valid.
	if !util.ValidUUID(form.Owner) {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid uuid")
		return
	}

	// Chcek if user by uuid exists.
	var existingUser model.User
	existingUser.UUID = form.Owner
	if err := existingUser.GetUser(handler.DB); err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "Owner ID does not exist")
		return
	}

	// Create new uuid for device.
	uuid, uuidErr := util.CreateUUID()
	if uuidErr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, uuidErr.Error())
		return
	}

	// Create new user instance.
	device := model.Device{UUID: uuid, Owner: form.Owner, Location: form.Location}

	// Generate new token.
	token := model.Token{
		UUID: device.UUID,
		Exp:  time.Now().Add(time.Hour * 24).Unix()}

	// Encode into string.
	tokenString, encodeErr := token.Encode()
	if encodeErr != nil {
		util.RespondWithError(w, http.StatusInternalServerError, encodeErr.Error())
		return
	}

	// Create new device.
	if err := device.CreateDevice(handler.DB); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// // Create new uuid for action.
	// actionUUID, uuidErr := util.CreateUUID()
	// if uuidErr != nil {
	// 	util.RespondWithError(w, http.StatusInternalServerError, uuidErr.Error())
	// 	return
	// }

	// // Create new action.
	// action := model.Action{UUID: actionUUID, User: device.Owner, Device: uuid, Active: true}

	// // Store action.
	// if err := action.CreateAction(handler.DB); err != nil {
	// 	util.RespondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	responseJSON := make(map[string]string)
	responseJSON["token"] = tokenString
	util.RespondWithJSON(w, http.StatusCreated, responseJSON)
}
