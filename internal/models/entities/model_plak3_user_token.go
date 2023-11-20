package entities

type Plak3UserTokens struct {
	UserId        int64  `json:"user_id"`
	LoginProvider string `json:"login_provider"`
	Name          string `json:"name"`
	Value         string `json:"value"`
}
