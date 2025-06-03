package controller

import (
	"easy-dictionary-server/domain"
	domainTranslation "easy-dictionary-server/domain/translation"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TranslationController struct {
	TranslationUseCase domainTranslation.TranslationUseCase
}

func (controller *TranslationController) GetAllForWord(c *gin.Context) {
	_, exists := c.Get("userID")
	wordId := c.Param("id")
	zap.S().Info("GET GetAllForWord")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if wordIDInt, err := strconv.Atoi(wordId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	} else {
		translations, err := controller.TranslationUseCase.GetAllForWord(c, wordIDInt)
		if err != nil {
			zap.S().Error("Failed to get translation")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Got translation %d", len(*translations))
			c.JSON(http.StatusOK, &translations)
		}
	}

}

func (controller *TranslationController) Edit(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Info("POST Edit")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request domainTranslation.EditTranslationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	idStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}
	if _, err := strconv.Atoi(idStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	} else {
		err := controller.TranslationUseCase.Update(c, request.ID, request.CategoryId, request.Translate, request.Description)
		if err != nil {
			zap.S().Error("Failed to update translation with " + request.Translate)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Translation updated %s", request.Translate)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Translation updated"})
		}
	}

}

func (controller *TranslationController) Create(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Infof("POST Create translation category for: %s", userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request domainTranslation.TranslationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	idStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}
	if _, err := strconv.Atoi(idStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	} else {
		err := controller.TranslationUseCase.Create(c, request.WordId, request.CategoryId, request.Translate, request.Description)
		if err != nil {
			zap.S().Error("Failed to create translation with " + request.Translate)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Translation created %s", request.Translate)
			c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Translation created"})
		}
	}
}

func (controller *TranslationController) Delete(c *gin.Context) {
	_, exists := c.Get("userID")
	translationId := c.Param("id")
	zap.S().Infof("DELETE Delete translation %d", translationId)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if translationIdInt, err := strconv.Atoi(translationId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid translation ID"})
		return
	} else {
		err := controller.TranslationUseCase.DeleteById(c, translationIdInt)
		if err != nil {
			zap.S().Error("Failed to delete translation by " + translationId)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Translation deleted %s", translationId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Translation deleted"})
		}
	}
}
