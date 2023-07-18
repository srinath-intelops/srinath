package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/srinath-intelops/srinath/test-component/pkg/rest/server/models"
	"github.com/srinath-intelops/srinath/test-component/pkg/rest/server/services"
	"net/http"
	"strconv"
)

type TestresourceController struct {
	testresourceService *services.TestresourceService
}

func NewTestresourceController() (*TestresourceController, error) {
	testresourceService, err := services.NewTestresourceService()
	if err != nil {
		return nil, err
	}
	return &TestresourceController{
		testresourceService: testresourceService,
	}, nil
}

func (testresourceController *TestresourceController) CreateTestresource(context *gin.Context) {
	// validate input
	var input models.Testresource
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger testresource creation
	if _, err := testresourceController.testresourceService.CreateTestresource(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Testresource created successfully"})
}

func (testresourceController *TestresourceController) UpdateTestresource(context *gin.Context) {
	// validate input
	var input models.Testresource
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testresource update
	if _, err := testresourceController.testresourceService.UpdateTestresource(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Testresource updated successfully"})
}

func (testresourceController *TestresourceController) FetchTestresource(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testresource fetching
	testresource, err := testresourceController.testresourceService.GetTestresource(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, testresource)
}

func (testresourceController *TestresourceController) DeleteTestresource(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger testresource deletion
	if err := testresourceController.testresourceService.DeleteTestresource(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Testresource deleted successfully",
	})
}

func (testresourceController *TestresourceController) ListTestresources(context *gin.Context) {
	// trigger all testresources fetching
	testresources, err := testresourceController.testresourceService.ListTestresources()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, testresources)
}

func (*TestresourceController) PatchTestresource(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*TestresourceController) OptionsTestresource(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*TestresourceController) HeadTestresource(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
