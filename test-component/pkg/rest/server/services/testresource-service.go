package services

import (
	"github.com/srinath-intelops/srinath/test-component/pkg/rest/server/daos"
	"github.com/srinath-intelops/srinath/test-component/pkg/rest/server/models"
)

type TestresourceService struct {
	testresourceDao *daos.TestresourceDao
}

func NewTestresourceService() (*TestresourceService, error) {
	testresourceDao, err := daos.NewTestresourceDao()
	if err != nil {
		return nil, err
	}
	return &TestresourceService{
		testresourceDao: testresourceDao,
	}, nil
}

func (testresourceService *TestresourceService) CreateTestresource(testresource *models.Testresource) (*models.Testresource, error) {
	return testresourceService.testresourceDao.CreateTestresource(testresource)
}

func (testresourceService *TestresourceService) UpdateTestresource(id int64, testresource *models.Testresource) (*models.Testresource, error) {
	return testresourceService.testresourceDao.UpdateTestresource(id, testresource)
}

func (testresourceService *TestresourceService) DeleteTestresource(id int64) error {
	return testresourceService.testresourceDao.DeleteTestresource(id)
}

func (testresourceService *TestresourceService) ListTestresources() ([]*models.Testresource, error) {
	return testresourceService.testresourceDao.ListTestresources()
}

func (testresourceService *TestresourceService) GetTestresource(id int64) (*models.Testresource, error) {
	return testresourceService.testresourceDao.GetTestresource(id)
}
