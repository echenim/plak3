package entities

type Plak3UserClaims struct {
	Id         int `json:"id"`
	UserId     int64  `json:"user_id"`
	ClaimType  string `json:"type"`
	ClaimValue string `json:"value"`
}
