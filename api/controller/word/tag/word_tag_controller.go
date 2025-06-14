package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
	"easy-dictionary-server/domain"
	domainWordTag "easy-dictionary-server/domain/word/tag"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WordTagController struct {
	WordTagUseCase domainWordTag.WordTagUseCase
}

// GetAllForDictionary godoc
// @Summary      Get all word tags for dictionary
// @Description  Get all word tags for dictionary
// @Tags         word_tag
// @Accept       json
// @Produce      json
// @Param   id    query     int     true     "Dictionary id"
// @Success      200  {array}  domainWordTag.WordTag
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/word/tag/all [get]
func (controller *WordTagController) GetAllForDictionary(c *gin.Context) {
	dictionaryId := c.Query("dictionaryId")
	zap.S().Infof("GET all word tags for dictionary %s", dictionaryId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	if dictionaryIdInt, err := strconv.Atoi(dictionaryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dictionary ID"})
		return
	} else {
		words, err := controller.WordTagUseCase.GetAllForDictionary(c, dictionaryIdInt)
		if err != nil {
			zap.S().Error("Failed to get word tags")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Got word tags %d", len(*words))
			c.JSON(http.StatusOK, &words)
		}
	}
}

// Edit godoc
// @Summary      Edit word tag for dictionary
// @Description  Update word tag for dictionary
// @Tags         word_tag
// @Accept       json
// @Produce      json
// @Param input body domainWordTag.EditWordTagRequest true "New data for word tag"
// @Success      200  {object}  domain.SuccessResponse
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/word/tag/edit [post]
func (controller *WordTagController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	var request domainWordTag.EditWordTagRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	err := controller.WordTagUseCase.Update(c, request.ID, request.DictionaryId, request.Name)
	if err != nil {
		zap.S().Error("Failed to update word tag with " + request.Name)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		zap.S().Debugf("Word tag updated %s", request.Name)
		c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Word tag updated"})
	}
}

// Create godoc
// @Summary Create word tag for dictionary
// @Description Create new word tag for dictionary
// @Tags word_tag
// @Accept  json
// @Produce  json
// @Param   input body domainWordTag.WordTagRequest true "Word tag data"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/word/tag/create [post]
func (controller *WordTagController) Create(c *gin.Context) {
	zap.S().Info("POST Create word tag")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	var request domainWordTag.WordTagRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	err := controller.WordTagUseCase.Create(c, request.DictionaryId, request.Name)
	if err != nil {
		zap.S().Error("Failed to create word tag with " + request.Name)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		zap.S().Debugf("Word created tag %s", request.Name)
		c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Word tag created"})
	}
}

// Delete godoc
// @Summary Delete word tag for dictionary
// @Description Delete word tag for dictionary
// @Tags word_tag
// @Accept  json
// @Produce  json
// @Param id path int true "ID word tag"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/word/tag/:id [delete]
func (controller *WordTagController) Delete(c *gin.Context) {
	wordTagId := c.Param("id")
	zap.S().Infof("DELETE Delete word tag %d", wordTagId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	if wordTagIdInt, err := strconv.Atoi(wordTagId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid word ID"})
		return
	} else {
		rows, err := controller.WordTagUseCase.DeleteById(c, wordTagIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, wordTagId, "Failed to delete word tag by", rows, err) {
			zap.S().Debugf("Word tag deleted %s", wordTagId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Word tag deleted"})
		}
	}
}
