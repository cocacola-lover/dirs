package drequests

type SendInfoRequest struct {
	Search string `json:"search"`
	Info   string `json:"info"`
}
