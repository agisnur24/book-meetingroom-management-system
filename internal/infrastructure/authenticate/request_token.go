package authenticate

type RequestGetTokenUser struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
