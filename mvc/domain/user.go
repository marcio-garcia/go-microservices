package domain

// User the user model
type User struct {
	ID       uint64 `json:"id"`
	FistName string `json:"fist_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}
