package controller

import (
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
	userID, exists := c.Get("userID")
	zap.S().Info("GET GetAllForUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	languages, err := languageController.LanguageUseCase.GetAllForUser(c, userID.(int))
	if err != nil {
		zap.S().Error("Failed to get languages")
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		zap.S().Debugf("Got languages %d", len(*languages))
		c.JSON(http.StatusOK, &languages)
	}
}

func (languageController *LanguageController) Edit(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Info("POST Edit")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request languageDomain.EditLanguageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	err := languageController.LanguageUseCase.Update(c, userID.(int), request.ID, request.Name, request.Code)
	if err != nil {
		zap.S().Error("Failed to update language with " + request.Name)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		zap.S().Debugf("Language updated %s %s", request.Name, request.Code)
		c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Language updated"})
	}
}

func (languageController *LanguageController) Create(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Infof("POST Create language for: %s", userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request languageDomain.LanguageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	err := languageController.LanguageUseCase.Create(c, userID.(int), request.Name, request.Code)
	if err != nil {
		zap.S().Error("Failed to create language with " + request.Name)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		zap.S().Debugf("Language updated %s %s", request.Name, request.Code)
		c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Language created"})
	}
}

func (languageController *LanguageController) Delete(c *gin.Context) {
	userID, exists := c.Get("userID")
	languageId := c.Param("id")
	zap.S().Infof("DELETE Delete language %d for user: %s", languageId, userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if languageIdInt, err := strconv.Atoi(languageId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID"})
		return
	} else {
		err := languageController.LanguageUseCase.DeleteById(c, languageIdInt)
		if err != nil {
			zap.S().Error("Failed to delete language by " + languageId)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Language deleted %s", languageId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Language deleted"})
		}
	}
}
