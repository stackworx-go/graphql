// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/99designs/gqlgen/graphql"
)

type Node interface {
	IsNode()
}

type CreateTodoInput struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type CreateTodoPayload struct {
	Todo *Todo `json:"todo"`
}

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

func (Todo) IsNode() {}

type UploadFile struct {
	ID   int            `json:"id"`
	File graphql.Upload `json:"file"`
}

type UploadFileInput struct {
	ID   string          `json:"id"`
	File *graphql.Upload `json:"file"`
}

type UploadFilePayload struct {
	ID   string `json:"id"`
	File *File  `json:"file"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (User) IsNode() {}
