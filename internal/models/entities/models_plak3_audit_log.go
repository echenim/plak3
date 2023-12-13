package entities

import "time"

type Plak3AuditLog struct {
	LogID          int64     `json:"id"`
	UserID         int64     `json:"user_id,omitempty"` // Pointer to handle nullable
	ActionType     string    `json:"action_type"`
	EntityType     string    `json:"entity_type"`
	EntityID       int64     `json:"entity_id"`
	Timestamp      time.Time `json:"timestamp"`
	IPAddress      string    `json:"ip_address,omitempty"`      // Empty string for nullable
	UserAgent      string    `json:"user_agent,omitempty"`      // Empty string for nullable
	ChangeDetails  string    `json:"change_details,omitempty"`  // JSON string
	AdditionalInfo string    `json:"additional_info,omitempty"` // JSON string
}
