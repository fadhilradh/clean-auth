package user

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fadhilradh/simple-auth/utils"
	"github.com/golang-jwt/jwt/v4"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		Repository: repository,
		timeout:    time.Duration(5) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Role:     req.Role,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
		Role:     r.Role,
		Message:  "Registration successful",
	}

	return res, nil
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func (s *service) Login(c context.Context, req *LoginReq) (*LoginRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginRes{}, err
	}

	err = utils.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginRes{}, err
	}

	fmt.Println(u.Role)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:   strconv.Itoa(int(u.ID)),
		Role: u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
		},
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return &LoginRes{}, err
	}

	return &LoginRes{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		Email:    u.Email,
		Token:    signedToken,
		Role:     u.Role,
		Message:  "Login successful",
	}, nil
}

func (s *service) GetUserById(c context.Context, id int64) (*GetUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetUserRes{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}, err
}

func (s *service) GetUsers(c context.Context) (*GetUsersRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var users []GetUserRes

	for i := 0; i < len(u); i++ {
		ur := GetUserRes{
			ID:       strconv.Itoa(int(u[i].ID)),
			Email:    u[i].Email,
			Username: u[i].Username,
			Role:     u[i].Role,
		}
		users = append(users, ur)
	}

	return &GetUsersRes{
		Users: &users,
	}, err
}
