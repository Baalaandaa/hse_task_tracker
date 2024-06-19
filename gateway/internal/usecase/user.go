package usecase

import (
	"context"
	"fmt"
	"gateway/internal/model"
	"gateway/internal/model/user"
	"gateway/internal/repository"
	"gateway/pkg/hash"
	"github.com/golang-jwt/jwt"
)

type UserUC struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewUserUC(userRepo *repository.UserRepository, jwtSecret string) *UserUC {
	return &UserUC{userRepo: userRepo, jwtSecret: jwtSecret}
}

func (u *UserUC) SignUp(ctx context.Context, in *user.SignUpInput) (*user.SignUpOutput, error) {
	usr, err := u.userRepo.GetByLogin(ctx, in.Login)
	if err != nil {
		return nil, err
	}
	if usr != nil {
		return &user.SignUpOutput{
			Response: model.Response{
				Success: false,
				Error:   "login already exists",
			},
		}, nil
	}
	_, err = u.userRepo.Create(ctx, &user.User{
		Login:    in.Login,
		Password: hash.Hash(in.Password),
	})
	if err != nil {
		return nil, err
	}
	return &user.SignUpOutput{
		Response: model.Response{
			Success: true,
		},
	}, nil
}

func (u *UserUC) SignIn(ctx context.Context, in *user.SignInInput) (*user.SignInOutput, error) {
	usr, err := u.userRepo.GetByLogin(ctx, in.Login)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return &user.SignInOutput{
			Response: model.Response{
				Success: false,
				Error:   "wrong login or password",
			},
		}, nil
	}
	if usr.Password != hash.Hash(in.Password) {
		return &user.SignInOutput{
			Response: model.Response{
				Success: false,
				Error:   "wrong login or password",
			},
		}, nil
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": usr.ID.String(),
	})
	res, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		fmt.Println(err, []byte(u.jwtSecret))
		return nil, err
	}
	return &user.SignInOutput{
		Response: model.Response{
			Success: true,
		},
		Token: res,
	}, nil
}

func (u *UserUC) UpdateInfo(ctx context.Context, in *user.UpdateInfoInput) (*user.UpdateInfoOutput, error) {
	usr, err := u.userRepo.GetById(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if in.Email != nil {
		usr.Email = in.Email
	}
	if in.Birthday != nil {
		usr.Birthday = in.Birthday
	}
	if in.Name != nil {
		usr.Name = in.Name
	}
	if in.Surname != nil {
		usr.Surname = in.Surname
	}
	if in.PhoneNumber != nil {
		usr.PhoneNumber = in.PhoneNumber
	}
	err = u.userRepo.Update(ctx, usr)
	if err != nil {
		return nil, err
	}
	return &user.UpdateInfoOutput{
		Response: model.Response{
			Success: true,
		},
		User: usr,
	}, nil
}
