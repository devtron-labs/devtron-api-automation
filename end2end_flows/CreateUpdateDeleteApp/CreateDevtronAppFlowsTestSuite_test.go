package CreateUpdateDeleteApp

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestCreateAppFlowsSuite(t *testing.T) {
	suite.Run(t, new(CreateDevtronAppFlowsTestSuite))
}
