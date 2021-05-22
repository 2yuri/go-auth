package login

type LoginRequest struct {
	Mail string `json:"email"`
	Pass string `json:"password"`
}
