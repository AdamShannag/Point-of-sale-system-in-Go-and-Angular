package auth

import (
	"net/http"

	"github.com/AdamShannag/toolkit"
)

func (a *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	// Call auth-service and register the user
	a.kit.WriteJSON(w, 200, toolkit.JSONResponse{Error: false, Message: "SignUp works"})
}

func (a *Auth) Signin(w http.ResponseWriter, r *http.Request) {
	// Call auth-service and sign in the user
	a.kit.WriteJSON(w, 200, toolkit.JSONResponse{Error: false, Message: "SignIn works"})
}

func (a *Auth) Signout(w http.ResponseWriter, r *http.Request) {
	// Logout the user, and message the auth service
	a.kit.WriteJSON(w, 200, toolkit.JSONResponse{Error: false, Message: "Signout works!"})
}

func (a *Auth) Signedin(w http.ResponseWriter, r *http.Request) {
	// Call auth-service and check if user is signed in
	a.kit.WriteJSON(w, 200, toolkit.JSONResponse{Error: false, Message: "SignedIn works!"})
}

func (a *Auth) Username(w http.ResponseWriter, r *http.Request) {
	// Call auth-service and check if user is signed in
	a.kit.WriteJSON(w, 200, toolkit.JSONResponse{Error: false, Message: "Username works!"})
}
