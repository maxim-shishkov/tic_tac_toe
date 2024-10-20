package api

import (
	"errors"
)

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

const (
	CodeNotFound   = "not_found"
	CodeBadRequest = "bad_request"
	CodeInternal   = "internal_error"
)

var (
	ErrNotFound   = errors.New(CodeNotFound)
	ErrBadRequest = errors.New(CodeBadRequest)
	ErrInternal   = errors.New(CodeInternal)
)
