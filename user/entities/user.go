package entities

import (
	"errors"
	"regexp"
	"time"
)

type CreateUserRequest struct {
	Phone     Phone
	FirstName Name
	LastName  Name
}

type User struct {
	Id        int32
	Phone     Phone
	Avatar    string
	FirstName Name
	LastName  Name
}

func (u User) AvatarPointer() *string {
	if len(u.Avatar) == 0 {
		return nil
	}
	return &u.Avatar
}

func (u User) GetFullName() string {
	return string(u.FirstName) + " " + string(u.LastName)
}

type UserAuth struct {
	Code      OtpCode
	UserId    UserId
	CreatedAt time.Time
}

type UserId int32

type Name string

type OtpCode int32

type OtpId int32

func NewName(name string) (*Name, error) {
	// Check that the name is not empty
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	// Check that the name matches the pattern for a valid name
	match, err := regexp.MatchString(`^[A-Za-z][A-Za-z'-]*[A-Za-z]$`, name)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, errors.New("name is not valid")
	}
	res := Name(name)
	return &res, nil
}
