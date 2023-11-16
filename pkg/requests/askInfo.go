package drequests

type AskInfoRequest struct {
	From   string `json:"name"`
	Search string `json:"search"`
}
