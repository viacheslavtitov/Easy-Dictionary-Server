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

// GetAllForUser godoc
// @Summary      Get all translation categories for user
// @Description  Get all translation categories for user
// @Tags         translation_category
// @Accept       json
// @Produce      json
// @Success      200  {array}  domainTranslationCategory.TranslationCategory
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/translation/category/all [get]
func (controller *TranslationCategoryController) GetAllForUser(c *gin.Context) {
	zap.S().Info("GET GetAllForUser")
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
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

// Edit godoc
// @Summary      Edit translation category for user
// @Description  Update translation category for user
// @Tags         translation_category
// @Accept       json
// @Produce      json
// @Param input body domainTranslationCategory.EditTranslationCategoryRequest true "New data for translation category"
// @Success      200  {object}  domain.SuccessResponse
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/translation/category/edit [post]
func (controller *TranslationCategoryController) Edit(c *gin.Context) {
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
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

// Create godoc
// @Summary Create translation category
// @Description Create new translation category for user
// @Tags translation_category
// @Accept  json
// @Produce  json
// @Param   input body domainTranslationCategory.TranslationCategoryRequest true "Translation category data"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/translation/category/create [post]
func (controller *TranslationCategoryController) Create(c *gin.Context) {
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
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

// Delete godoc
// @Summary Delete translation category
// @Description Delete translation category for user
// @Tags translation_category
// @Accept  json
// @Produce  json
// @Param id path int true "ID translation category"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/translation/category/:id [delete]
func (controller *TranslationCategoryController) Delete(c *gin.Context) {
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
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
