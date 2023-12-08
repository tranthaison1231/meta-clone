package models

type SocketRequest[T any] struct {
	Action  string `json:"action"`
	Payload T      `json:"payload"`
}

type SocketResponse[T any] struct {
	Action  string `json:"action"`
	Payload T      `json:"payload"`
}
