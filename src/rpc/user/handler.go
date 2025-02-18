package main

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	user "src/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	db *gorm.DB
}

func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	log.Printf("User: %s (%s) register with passwd: %s\n", req.Username, req.Email, req.Password)

	newUser := User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	result := s.db.WithContext(ctx).Create(&newUser)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create user: %w", result.Error)
	}

	log.Printf("User registered: %s (%s)", req.Username, req.Email)

	return &user.RegisterResp{UserId: newUser.ID}, nil

}

func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	log.Printf("User: %s login with passwd: %s\n", req.Email, req.Password)

	var queryItem User
	// check email and password
	result := s.db.WithContext(ctx).Where("email = ? AND password = ?", req.Email, req.Password).First(&queryItem)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find user: %w", result.Error)
	}

	return &user.LoginResp{UserId: queryItem.ID}, nil
}
