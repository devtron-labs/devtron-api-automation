package ResponseDTOs

type ChartGroupInstallResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result struct {
	} `json:"result"`
}
