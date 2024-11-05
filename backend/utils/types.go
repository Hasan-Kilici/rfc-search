package utils
type RFC struct {
	Number   int      `json:"number"`
	Files    []string `json:"files"`
	Title    string   `json:"title"`
	Authors  []string `json:"authors"`
	Date     string   `json:"date"`
	MoreInfo string   `json:"more_info"`
	Status   string   `json:"status"`
}