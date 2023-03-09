package CreateUpdateDeleteApp

import (
	"automation-suite/PipelineConfigRouter"
	PipelineConfigRouterResponseDTOs "automation-suite/PipelineConfigRouter/ResponseDTOs"
	"automation-suite/TeamRouter"
	TeamRouterResponseDTOs "automation-suite/TeamRouter/ResponseDTOs"
	Base "automation-suite/testUtils"
	"github.com/stretchr/testify/suite"
)

func GetAllTeamsAutocomplete(authToken string) TeamRouterResponseDTOs.FetchAllTeamResponseDTO {
	return TeamRouter.HitFetchForAutocompleteApi(authToken)
}

func GetAllEnvsAutocomplete(queryParams map[string]string, authToken string) PipelineConfigRouter.EnvironmentDetailsResponseDTO {
	return PipelineConfigRouter.HitGetAllEnvironmentDetails(queryParams, authToken)
}

func CreatePayloadForDevtronApp(appName string, teamId int, templateId int, labels []PipelineConfigRouter.Labels) PipelineConfigRouter.CreateAppRequestDto {
	return PipelineConfigRouter.GetPayloadForCreateDevtronApp(appName, teamId, templateId, labels)
}

func CreatDevtronApp(payload []byte, appName string, teamId int, templateId int, authToken string) PipelineConfigRouter.CreateAppResponseDto {
	return PipelineConfigRouter.HitCreateDevtronAppApi(payload, authToken)
}

func GetGitAccountsAutocomplete(appId string, authToken string) PipelineConfigRouterResponseDTOs.GetGitAccountsAutocompleteResponseDTO {
	return PipelineConfigRouter.HitGetGitListAutocomplete(appId, authToken)
}

func SaveGitRepository(payload []byte, appId int, gitProviderId int, fetchSubmodules bool, authToken string) PipelineConfigRouter.CreateAppMaterialResponseDto {
	return PipelineConfigRouter.HitCreateAppMaterialApi(payload, 0, 0, true, authToken)
}

func GetContainerRegistriesAutocomplete(appId string, authToken string) PipelineConfigRouter.GetContainerRegistryResponseDTO {
	return PipelineConfigRouter.HitGetContainerRegistry(appId, authToken)
}

func SaveBuildConfigurations(payload []byte, authToken string) PipelineConfigRouter.SaveAppCiPipelineResponseDTO {
	return PipelineConfigRouter.HitSaveAppCiPipeline(payload, authToken)
}

func GetChartRefAutocompleteForBaseDeploymentTemplate(appId string, authToken string) PipelineConfigRouter.GetChartReferenceResponseDTO {
	return PipelineConfigRouter.HitGetChartReferenceViaAppId(appId, authToken)
}

func GetBaseDeploymentTemplate(appId string, chartRefId string, authToken string) PipelineConfigRouter.GetAppTemplateResponseDto {
	return PipelineConfigRouter.HitGetTemplateViaAppIdAndChartRefId(appId, chartRefId, authToken)
}

func SaveAppTemplate(payload []byte, authToken string) PipelineConfigRouter.SaveDeploymentTemplateResponseDTO {
	return PipelineConfigRouter.HitSaveDeploymentTemplateApi(payload, authToken)
}

type CreateDevtronAppFlowsTestSuite struct {
	suite.Suite
	authToken string
}

func (suite *CreateDevtronAppFlowsTestSuite) SetupSuite() {
	suite.authToken = Base.GetAuthToken()
}
