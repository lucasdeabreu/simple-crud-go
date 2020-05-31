package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Handler handler for requests to manager User
type Handler struct {
	Service *Service
}

// GetAll users
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := h.Service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userJson, err := toJSON(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/userJson; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, userJson)
}

// Create new user
func (h *Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err = h.Service.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userJson, err := toJSON(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/userJson; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, userJson)
}

// GetByID Get an user by ID
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.Service.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userJson, err := toJSON(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/userJson; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, userJson)
}

// Update Update an user
func (h *Handler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID, err = strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err = h.Service.Update(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userJson, err := toJSON(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/userJson; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, userJson)
}

// DeleteByID Delete an user by ID
func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.Service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func toJSON(entity interface{}) (string, error) {
	userJson, err := json.MarshalIndent(&entity, "", "\t")
	if err != nil {
		return "", err
	}
	return string(userJson), nil
}
