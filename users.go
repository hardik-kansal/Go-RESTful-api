package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*

handleUserRegister -> get Payload from body and unmarshall it with desired struct type
-> hashpassword -> sql query for create user -> createjwt (userId got from sql entry)
-> setcookie under header "Authorization" value "Token"

handlecreateTask -> JWT auTH -> get Payload from body and unmarshall it with desired struct type
-> sql query for create task

JWT auth
->func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

->validateJWT func

->claims := token.Claims.(jwt.MapClaims)
	userID := claims["userID"].(string)
	_, err = store.GetUserByID(userID)

*/

var errEmailRequired = errors.New("email is required")
var errFirstNameRequired = errors.New("first name is required")
var errLastNameRequired = errors.New("last name is required")
var errPasswordRequired = errors.New("password is required")
var errAddressrequired=errors.New("address is required")

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
	r.HandleFunc("/users/login", s.handleUserLogin).Methods("POST")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var payload *User
	var payload1 *UserVerify
	err = json.Unmarshal(body, &payload)
	err = json.Unmarshal(body, &payload1)


	if err != nil {
		log.Printf("Error unmarshalling JSON: %v\n", err)
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateUserPayload(payload); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}


	hashedPassword, err := HashPassword(payload.Password)

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}
	payload.Password = hashedPassword
	payload1.Password = hashedPassword

	u, err := s.store.CreateUser(payload)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}
	if payload1.ethAddress != ""{
		token, err := createAndSetEWTCookie(payload1.ethAddress,payload1.msg,payload1.sig,w)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
			return
		}
		WriteJSON(w, http.StatusCreated, token)

	}else{
	    token, err := createAndSetAuthCookie(u.ID, w)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
			return
		}
		WriteJSON(w, http.StatusCreated, token)

	}



}

func (s *UserService) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	// 1. Find user in db by email
	// 2. Compare password with hashed password
	// 3. Create JWT and set it in a cookie
	// 4. Return JWT in response
}

func validateUserPayload(user *User) (error) {
	if user.Email == "" {
		return errEmailRequired
	}

	if user.FirstName == "" {
		return errFirstNameRequired
	}

	if user.LastName == "" {
		return errLastNameRequired
	}

	if user.Password == "" {
		return errPasswordRequired
	}

	return nil
}

