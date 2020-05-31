package user

// User an user of the system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
}

// Users a slice of User
type Users []User
