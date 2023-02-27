package AppStoreDiscoverRouter

import (
	"automation-suite/TeamRouter"
	Base "automation-suite/testUtils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
)

func (suite *AppStoreDiscoverTestSuite) TestChartGroupInstall() {

	suite.Run("A=1=InstallingOneHelmAppWithValidPayloadWithoutChartGroup", func() {
		log.Println("=== Fetching Payload ===")
		responseFetchAllTeamApi := TeamRouter.HitFetchAllTeamApi(suite.authToken)
		log.Println("=== Here we are using airflow chart repo ===")
		queryParams := map[string]string{"appStoreName": "apache"}
		PollForGettingHelmAppData(queryParams, suite.authToken)
		discoverAppApiResponse := HitDiscoverAppApi(queryParams, suite.authToken)
		projectId := responseFetchAllTeamApi.Result[0].Id
		appName := "automation-preset-" + strings.ToLower(Base.GetRandomStringOfGivenLength(5))
		log.Printf("=== App Name :- %s ===\n", appName)
		environmentId := 1
		appstoreVersionId := discoverAppApiResponse.Result[0].AppStoreApplicationVersionId
		referenceValueId := discoverAppApiResponse.Result[0].AppStoreApplicationVersionId
		requestPayload := getPayloadforDeployBulkApps(projectId, appName, environmentId, appstoreVersionId,
			"", referenceValueId, "DEFAULT", 0,
			true, 0, 1)
		payloadByteArray, _ := json.Marshal(requestPayload)
		responseofHitApiDeployBulkApps := HitApiDeployBulkApps(string(payloadByteArray), suite.authToken)
		log.Printf("=== reponse of Deploy Bulk Apps API :- code: %d, status: %s ===\n",
			responseofHitApiDeployBulkApps.Code, responseofHitApiDeployBulkApps.Status)
		assert.Equal(suite.T(), 200, responseofHitApiDeployBulkApps.Code)
		assert.Equal(suite.T(), "OK", responseofHitApiDeployBulkApps.Status)
		assert.Empty(suite.T(), responseofHitApiDeployBulkApps.Result)
		log.Println("Deleting the Helm App created via API")
		respOfGetAllInstallAppApi := HitApiGetAllInstalledApps(suite.authToken)
		arrhelmApps := respOfGetAllInstallAppApi.Result.HelmApps
		var appId string
		for _, helmApp := range arrhelmApps {
			if helmApp.AppName == appName {
				appId = helmApp.AppId
			}
		}
		HitDeleteInstalledAppApi(appId, suite.authToken)
	})
}
