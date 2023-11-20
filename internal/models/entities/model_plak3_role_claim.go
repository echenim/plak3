package entities

type Plak3RoleClaims struct {
	Id         int    `json:"id"`
	RoleId     int64  `json:"role_id"`
	ClaimType  string `json:"claim_type"`
	ClaimValue string `json:"claim_value"`
}
