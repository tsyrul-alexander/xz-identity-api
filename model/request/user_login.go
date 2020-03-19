package request

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}