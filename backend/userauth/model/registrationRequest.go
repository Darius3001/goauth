package userauth

type registrationRequest struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
