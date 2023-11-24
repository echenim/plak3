package views

type Plak3UserRoles struct {
	RoleId   string `json:"role_id"`
	UserId   int64  `json:"user_id"`
	RoleName string `json:"role_name"`
}

type Plak3UserAndRoles struct {
	UserId int64    `json:"user_id"`
	Roles  []string `json:"roles"`
}
