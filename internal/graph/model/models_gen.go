// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Admin struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	SecurityKey string `json:"securityKey"`
}

type AdminInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	SecurityKey     string `json:"securityKey"`
}

type AdminSignIn struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	SecurityKey string `json:"securityKey"`
}

type AdminSignInResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type CreateAdminResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Admin   string `json:"admin"`
}

type Mutation struct {
}

type Query struct {
}
