package searchmodels

// UserSearchCriteria  represents the criteria for searching users
// @Description User model.
// @Name UserSearchCriteria
type UserSearchCriteria struct {
	ID        int64  `json:"Id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}
