package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
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
	zap.S().Info("GET GetAllForUser")
	if userID, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		tcategories, err := controller.TranslationCategoryUseCase.GetAllForUser(c, *userID)
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
	if userID, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		var request domainTranslationCategory.EditTranslationCategoryRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		err := controller.TranslationCategoryUseCase.Update(c, *userID, request.ID, request.DictionaryId, request.Name)
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
	if userID, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		zap.S().Infof("POST Create translation category for: %d", &userID)
		var request domainTranslationCategory.TranslationCategoryRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		err := controller.TranslationCategoryUseCase.Create(c, *userID, request.DictionaryId, request.Name)
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
	if _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	translationCategoryId := c.Param("id")
	zap.S().Infof("DELETE Delete translation category %d", translationCategoryId)
	if translationCategoryIdInt, err := strconv.Atoi(translationCategoryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid translation category ID"})
		return
	} else {
		rows, err := controller.TranslationCategoryUseCase.DeleteById(c, translationCategoryIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, translationCategoryId, "Failed to delete translation category by", rows, err) {
			zap.S().Debugf("Translation category deleted %s", translationCategoryId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Translation category deleted"})
		}
	}
}
