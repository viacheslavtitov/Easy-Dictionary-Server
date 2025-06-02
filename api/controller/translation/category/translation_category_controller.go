package controller

import (
	"easy-dictionary-server/domain"
	domainTranslationCategory "easy-dictionary-server/domain/translation/category"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TranslationCategoryController struct {
	TranslationCategoryUseCase domainTranslationCategory.TranslationCategoryUseCase
}

func (controller *TranslationCategoryController) GetAllForUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Info("GET GetAllForUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	idStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}
	if userIDInt, err := strconv.Atoi(idStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	} else {
		tcategories, err := controller.TranslationCategoryUseCase.GetAllForUser(c, userIDInt)
		if err != nil {
			zap.S().Error("Failed to get translation categories")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Got translation categories %d", len(*tcategories))
			c.JSON(http.StatusOK, &tcategories)
		}
	}

}

func (controller *TranslationCategoryController) Edit(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Info("POST Edit")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request domainTranslationCategory.EditTranslationCategoryRequest
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
	if userIDInt, err := strconv.Atoi(idStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	} else {
		err := controller.TranslationCategoryUseCase.Update(c, userIDInt, request.ID, request.DictionaryId, request.Name)
		if err != nil {
			zap.S().Error("Failed to update translation category with " + request.Name)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Translation category updated %s", request.Name)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Translation category updated"})
		}
	}

}

func (controller *TranslationCategoryController) Create(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Infof("POST Create translation category for: %s", userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request domainTranslationCategory.TranslationCategoryRequest
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
	if userIDInt, err := strconv.Atoi(idStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	} else {
		err := controller.TranslationCategoryUseCase.Create(c, userIDInt, request.DictionaryId, request.Name)
		if err != nil {
			zap.S().Error("Failed to create translation category with " + request.Name)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Translation category created %s", request.Name)
			c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Translation category created"})
		}
	}
}

func (controller *TranslationCategoryController) Delete(c *gin.Context) {
	userID, exists := c.Get("userID")
	translationCategoryId := c.Param("id")
	zap.S().Infof("DELETE Delete translation category %d for user: %s", translationCategoryId, userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if translationCategoryIdInt, err := strconv.Atoi(translationCategoryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID"})
		return
	} else {
		err := controller.TranslationCategoryUseCase.DeleteById(c, translationCategoryIdInt)
		if err != nil {
			zap.S().Error("Failed to delete translation category by " + translationCategoryId)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Translation category deleted %s", translationCategoryId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Translation category deleted"})
		}
	}
}
