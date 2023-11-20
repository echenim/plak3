package entities

type Plak3UserLogins struct {
	LoginProvider       string `json:"loginProvider"`
	ProviderKey         string `json:"providerKey"`
	ProviderDisplayName string `json:"providerDisplayName"`
	UserId              int64  `json:"user_id"`
}
