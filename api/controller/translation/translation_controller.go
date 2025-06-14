package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
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

// GetAllForWord godoc
// @Summary      Get all translation for word
// @Description  Get all translation for word
// @Tags         translation
// @Accept       json
// @Produce      json
// @Param   id    path     int     true     "Word id"
// @Success      200  {array}  domainTranslation.Translation
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/translation/all/:id [get]
func (controller *TranslationController) GetAllForWord(c *gin.Context) {
	wordId := c.Param("id")
	zap.S().Infof("GET GetAllForWord %s", wordId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
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

// Edit godoc
// @Summary      Edit translation
// @Description  Update translation
// @Tags         translation
// @Accept       json
// @Produce      json
// @Param input body domainTranslation.EditTranslationRequest true "New data for translation"
// @Success      200  {object}  domain.SuccessResponse
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/translation/edit [post]
func (controller *TranslationController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	var request domainTranslation.EditTranslationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
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

// Create godoc
// @Summary Create translation
// @Description Create new translation for word
// @Tags translation
// @Accept  json
// @Produce  json
// @Param   input body domainTranslation.TranslationRequest true "Translation data"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/translation/create [post]
func (controller *TranslationController) Create(c *gin.Context) {
	zap.S().Infof("POST Create translation category")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	var request domainTranslation.TranslationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
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

// Delete godoc
// @Summary Delete translation
// @Description Delete translation for word
// @Tags translation
// @Accept  json
// @Produce  json
// @Param id path int true "ID translation"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/translation/:id [delete]
func (controller *TranslationController) Delete(c *gin.Context) {
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	translationId := c.Param("id")
	if translationIdInt, err := strconv.Atoi(translationId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid translation ID"})
		return
	} else {
		rows, err := controller.TranslationUseCase.DeleteById(c, translationIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, translationId, "Failed to delete translation by", rows, err) {
			zap.S().Debugf("Translation deleted %s", translationId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Translation deleted"})
		}
	}
}
