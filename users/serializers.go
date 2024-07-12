package users

import (
	"github.com/gin-gonic/gin"

	"github.com/gothinkster/golang-gin-realworld-example-app/common"
)

type ProfileSerializer struct {
	C *gin.Context
	UserModel
}

// Declare your response schema here
type ProfileResponse struct {
	ID        uint    `json:"-"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}

// Put your response logic including wrap the userModel here.
func (ps *ProfileSerializer) Response() ProfileResponse {
	myUserModel := ps.C.MustGet("my_user_model").(UserModel)
	profile := ProfileResponse{
		ID:        ps.ID,
		Username:  ps.Username,
		Bio:       ps.Bio,
		Image:     ps.Image,
		Following: myUserModel.isFollowing(ps.UserModel),
	}
	return profile
}

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func (us *UserSerializer) Response() UserResponse {
	myUserModel := us.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}
