package entity

type TaskStatus int

const (
	TODO TaskStatus = iota
	IN_PROGRESS
	DONE
)
