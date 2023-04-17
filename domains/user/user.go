package user

import "context"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type CreateUserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
	Role     string `json:"role"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Message  string `json:"message"`
	Token    string `json:"token"`
}

type GetUserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type GetUsersRes struct {
	Users *[]GetUserRes `json:"users"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id int64) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
}

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(ctx context.Context, req *LoginReq) (*LoginRes, error)
	GetUserById(ctx context.Context, id int64) (*GetUserRes, error)
	GetUsers(ctx context.Context) (*GetUsersRes, error)
}
