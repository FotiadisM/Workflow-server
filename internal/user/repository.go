package user

type Role string

const (
	Admin  Role = "admin"
	Normal Role = "normal"
)

type User struct {
	ID       string `json:"id"`
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Email    string `json:"email"`
	Company  string `json:"company"`
	Position string `json:"position"`
	Role     Role   `json:"role"`
}

type Repository interface {
}
