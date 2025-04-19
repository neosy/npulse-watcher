package dto

type RegisterRequest struct {
	IPAddress string `json:"ip_address"`
	HostName  string `json:"host_hame"`
}
