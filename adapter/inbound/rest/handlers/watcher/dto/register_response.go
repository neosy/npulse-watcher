package dto

type RegisterStatus string

const (
	RegisterStatusSuccess RegisterStatus = "success"
	RegisterStatusFailed  RegisterStatus = "failed"
)

type RegisterResponse struct {
	Status RegisterStatus `json:"status"`
}
