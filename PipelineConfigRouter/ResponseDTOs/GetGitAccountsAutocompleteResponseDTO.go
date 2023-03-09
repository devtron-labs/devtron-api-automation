package ResponseDTOs

type GetGitAccountsAutocompleteResponseDTO struct {
	Code   int          `json:"code"`
	Status string       `json:"status"`
	Result []GitDetails `json:"result"`
}

type GitDetails struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	AuthMode  string `json:"authMode"`
	Active    bool   `json:"active"`
	GitHostId int    `json:"gitHostId"`
}
