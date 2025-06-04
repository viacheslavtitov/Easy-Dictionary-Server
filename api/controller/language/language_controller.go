package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
	"easy-dictionary-server/domain"
	languageDomain "easy-dictionary-server/domain/language"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LanguageController struct {
	LanguageUseCase languageDomain.LanguageUseCase
}

func (languageController *LanguageController) GetAllForUser(c *gin.Context) {
	zap.S().Info("GET GetAllForUser")
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		languages, err := languageController.LanguageUseCase.GetAllForUser(c, *userID)
		if err != nil {
			zap.S().Error("Failed to get languages")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Got languages %d", len(*languages))
			c.JSON(http.StatusOK, &languages)
		}
	}
}

func (languageController *LanguageController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		var request languageDomain.EditLanguageRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		err := languageController.LanguageUseCase.Update(c, *userID, request.ID, request.Name, request.Code)
		if err != nil {
			zap.S().Error("Failed to update language with " + request.Name)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Language updated %s %s", request.Name, request.Code)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Language updated"})
		}
	}
}

func (languageController *LanguageController) Create(c *gin.Context) {
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		zap.S().Infof("POST Create language for: %d", &userID)
		var request languageDomain.LanguageRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		err := languageController.LanguageUseCase.Create(c, *userID, request.Name, request.Code)
		if err != nil {
			zap.S().Error("Failed to create language with " + request.Name)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Language created %s %s", request.Name, request.Code)
			c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Language created"})
		}
	}
}

func (languageController *LanguageController) Delete(c *gin.Context) {
	languageId := c.Param("id")
	zap.S().Infof("DELETE Delete language %d", languageId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	if languageIdInt, err := strconv.Atoi(languageId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID"})
		return
	} else {
		rows, err := languageController.LanguageUseCase.DeleteById(c, languageIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, languageId, "Failed to delete language by", rows, err) {
			zap.S().Debugf("Language deleted %s", languageId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Language deleted"})
		}
	}
}
