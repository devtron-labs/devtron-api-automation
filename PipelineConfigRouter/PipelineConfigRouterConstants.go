package PipelineConfigRouter

const (
	SaveAppCiPipelineApiUrl                   string = "/orchestrator/app/ci-pipeline"
	SaveAppCiPipelineApi                      string = "SaveAppCiPipelineApi"
	CreateAppApiUrl                           string = "/orchestrator/app"
	CreateAppApi                              string = "CreateAppApi"
	DeleteAppApi                              string = "DeleteAppApi"
	CreateAppMaterialApiUrl                   string = "/orchestrator/app/material"
	CreateAppMaterialApi                      string = "CreateAppMaterialApi"
	DeleteAppMaterialApi                      string = "DeleteAppMaterialApi"
	GetAppDetailsApiUrl                       string = "/orchestrator/app/get"
	GetCiPipelineViaIdApi                     string = "GetCiPipelineViaIdApi"
	GetAppDetailsApi                          string = "FetchAppGetApi"
	GetCiPipelineViaIdApiUrl                  string = "/orchestrator/app/ci-pipeline/"
	GetContainerRegistryApi                   string = "GetContainerRegistryApi"
	GetContainerRegistryApiUrl                string = "/orchestrator/app/"
	GetChartReferenceViaAppIdApi              string = "GetChartReferenceViaAppIdApi"
	GetChartReferenceViaAppIdApiUrl           string = "/orchestrator/chartref/autocomplete/"
	GetAppTemplateViaAppIdAndChartRefIdApi    string = "GetAppTemplateViaAppIdAndChartRefIdApi"
	GetAppTemplateViaAppIdAndChartRefIdApiUrl string = "/orchestrator/app/template/"
	GetCdPipelineStrategiesApiUrl             string = "/orchestrator/app/cd-pipeline/strategies/"
	GetCdPipelineStrategiesApi                string = "GetCdPipelineStrategiesApi"
	GetPipelineSuggestedCICDApiUrl            string = "/orchestrator/app/pipeline/suggest/"
	GetPipelineSuggestedCICDApi               string = "GetPipelineSuggestedCICDApi"
	GetEnvAutocompleteApiUrl                  string = "/orchestrator/env/autocomplete"
	GetEnvAutocompleteApi                     string = "GetEnvAutocompleteApi"
	SaveDeploymentTemplateAPiUrl              string = "/orchestrator/app/template"
	SaveDeploymentTemplateApi                 string = "SaveDeploymentTemplateApi"
	CreateWorkflowApiUrl                      string = "/orchestrator/app/ci-pipeline/patch"
	CreateWorkflowApi                         string = "CreateWorkflowApi"
	GetWorkflowDetailsApi                     string = "GetWorkflowDetailsApi"
	GetWorkflowApiUrl                         string = "/orchestrator/app/app-wf/"
	FetchAllAppWorkflowApi                    string = "FetchAllAppWorkflowApi"
	FetchSuggestedCiPipelineNameApi           string = "FetchSuggestedCiPipelineNameApi"
	SaveCdPipelineApiUrl                      string = "/orchestrator/app/cd-pipeline"
	SaveCdPipelineApi                         string = "SaveCdPipelineApi"
	Automatic                                 string = "AUTOMATIC"
	Manual                                    string = "MANUAL"
	DeleteCdPipelineApiUrl                    string = "/orchestrator/app/cd-pipeline/patch"
	ForceDeleteCdPipelineApiUrl               string = "/orchestrator/app/cd-pipeline/patch?force=true"
	DeleteCdPipelineApi                       string = "DeleteCdPipelineApi"
	GetAppCdPipelineApiUrl                    string = "/orchestrator/app/cd-pipeline/"
	GetAppCdPipelineApi                       string = "GetAppCdPipelineApi"
	GetWorkflowStatusApi                      string = "GetWorkflowStatusApi"
	GetWorkflowStatusApiUrl                   string = "/orchestrator/app/workflow/status/"
	GetCiPipelineMaterialApi                  string = "GetCiPipelineMaterialApi"
	TriggerCiPipelineApiUrl                   string = "/orchestrator/app/ci-pipeline/trigger"
	TriggerCiPipelineApi                      string = "TriggerCiPipelineApi"
	UpdateAppMaterial                         string = "UpdateAppMaterial"
	GetAppListForAutocompleteApi              string = "GetAppListForAutocompleteApi"
	GetAppListForAutocompleteApiUrl           string = "/orchestrator/app/autocomplete"
	GetAppListByTeamIdsApi                    string = "GetAppListByTeamIdsApi"
	GetAppListByTeamIdsApiUrl                 string = "/orchestrator/app/min"
	FindAppsByTeamIdApiUrl                    string = "/orchestrator/app/team/by-id/"
	FindAppsByTeamNameApiUrl                  string = "/orchestrator/app/team/by-name/"
)
