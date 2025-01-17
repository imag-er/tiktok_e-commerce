package main

import (
	"context"
	"log"
	user "github.com/imag-er/tiktok_e-commerce/src/user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	log.Printf("User: %s register with passwd: %s, confirm passwd: %s\n", req.Email, req.Password,req.ConfirmPassword)
	return &user.RegisterResp{
		UserId: 123456,
	}, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	log.Printf("User: %s login with passwd: %s\n", req.Email, req.Password)
	return &user.LoginResp{
		UserId: 123456,
	}, nil
}
