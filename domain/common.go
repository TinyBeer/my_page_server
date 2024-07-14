package domain

const (
	StatusError = "error"
	StatusOk    = "ok"

	ErrorNoToken = "no token"
)

type BaseResp struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
