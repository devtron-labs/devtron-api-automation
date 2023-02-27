package RequestDTOs

type ChartGroupInstallRequestDTO struct {
	ProjectId                     int                             `json:"projectId"`
	ChartGroupInstallChartRequest []ChartGroupInstallChartRequest `json:"charts"`
	ChartGroupId                  int                             `json:"chartGroupId"` //optional
}

type ChartGroupInstallChartRequest struct {
	AppName                 string `json:"appName"`
	EnvironmentId           int    `json:"environmentId"`
	AppStoreVersion         int    `json:"appStoreVersion"`
	ValuesOverrideYaml      string `json:"valuesOverrideYaml"` //optional
	ReferenceValueId        int    `json:"referenceValueId"`
	ReferenceValueKind      string `json:"referenceValueKind"`
	ChartGroupEntryId       int    `json:"chartGroupEntryId"` //optional
	DefaultClusterComponent bool   `json:"-"`
}
