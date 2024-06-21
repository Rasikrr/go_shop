package requests

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}
