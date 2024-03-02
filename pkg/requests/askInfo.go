package drequests

type AskInfoRequest struct {
	Id     string `json:"id"`
	From   string `json:"from"`
	Search string `json:"search"`
}
