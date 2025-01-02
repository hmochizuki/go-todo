package domain

type User struct {
	ID    string
	Name  string
	Email string
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	GetUserById(id string) (User, error)
	CreateUser(user CreateUserRequest) error
}
