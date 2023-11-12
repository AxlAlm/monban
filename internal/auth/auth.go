package auth

type AuthInput struct {
	ApiKey string `json:"api_key"`
}

type Authorizer interface {
	Authorize()
}
