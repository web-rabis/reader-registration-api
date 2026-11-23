package model

import "github.com/web-rabis/reader-client/protobuf"

type User struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	State    string `json:"state,omitempty"`
}

func (u *User) ToProto() *protobuf.User {
	if u == nil {
		return nil
	}
	return &protobuf.User{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		State:    u.State,
	}
}
func UserFromProto(u *protobuf.User) *User {
	if u == nil {
		return nil
	}
	return &User{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		State:    u.State,
	}
}
