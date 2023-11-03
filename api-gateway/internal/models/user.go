package models

type User struct {
	UserId    string `json:"user_id,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	Email     string `json:"email,omitempty"`
	UserType  string `json:"user_type,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type RegisterUser struct {
	FullName        string `json:"full_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
	UserType        string `json:"user_type,omitempty"`
}

type SignIn struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
