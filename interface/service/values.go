package service

import "errors"

// defines different levels of priority for services
const (
	PRIORITY_LOW = iota
	PRIORITY_MEDIUM
	PRIORITY_HIGH
)

var (
	ERR_PERMISSION_DENIED = errors.New("permission denied")
	ERR_DEFAULT           = errors.New("default error")
)
