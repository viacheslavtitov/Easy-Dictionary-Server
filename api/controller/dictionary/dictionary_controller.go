package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
	"easy-dictionary-server/domain"
	dictionaryDomain "easy-dictionary-server/domain/dictionary"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictionaryController struct {
	DictionaryUseCase dictionaryDomain.DictionaryUseCase
}

func (dictionaryController *DictionaryController) GetAllForUser(c *gin.Context) {
	zap.S().Info("GET GetAllForUser")
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		dictionaries, err := dictionaryController.DictionaryUseCase.GetAllForUser(c, *userID)
		if err != nil {
			zap.S().Error("Failed to get languages")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Got languages %d", len(*dictionaries))
			c.JSON(http.StatusOK, &dictionaries)
		}
	}
}

func (dictionaryController *DictionaryController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		var request dictionaryDomain.EditDictionaryRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		err := dictionaryController.DictionaryUseCase.Update(c, *userID, request.ID, request.Dialect, request.LangFromId, request.LangToId)
		if err != nil {
			zap.S().Errorf("Failed to update dictionary by id %d", request.ID)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Dictionary updated %d", request.ID)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Dictionary updated"})
		}
	}
}

func (dictionaryController *DictionaryController) Create(c *gin.Context) {
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		zap.S().Infof("POST Create dictionary for: %d", &userID)
		var request dictionaryDomain.DictionaryRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		err := dictionaryController.DictionaryUseCase.Create(c, *userID, request.Dialect, request.LangFromId, request.LangToId)
		if err != nil {
			zap.S().Error("Failed to create dictionary ")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Dictionary created")
			c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Dictionary created"})
		}
	}
}

func (dictionaryController *DictionaryController) Delete(c *gin.Context) {
	dictionaryId := c.Param("id")
	zap.S().Infof("DELETE Delete dictionary %d", dictionaryId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	if dictionaryIdInt, err := strconv.Atoi(dictionaryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dictionary ID"})
		return
	} else {
		rows, err := dictionaryController.DictionaryUseCase.DeleteById(c, dictionaryIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, dictionaryId, "Failed to delete dictionary by", rows, err) {
			zap.S().Debugf("Dictionary deleted %s", dictionaryId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Dictionary deleted"})
		}
	}
}
