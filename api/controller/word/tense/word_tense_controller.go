package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
	"easy-dictionary-server/domain"
	domainWordTense "easy-dictionary-server/domain/word/tense"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WordTenseController struct {
	WordTenseUseCase domainWordTense.WordTenseUseCase
}

// GetAllForWord godoc
// @Summary      Get all word tenses for word
// @Description  Get all word tenses for word
// @Tags         word_tense
// @Accept       json
// @Produce      json
// @Param   id    query     int     true     "Word id"
// @Success      200  {array}  domainWordTense.WordTense
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/word/tense/all [get]
func (controller *WordTenseController) GetAllForWord(c *gin.Context) {
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	wordIdInt, err := controllerCommon.ParseQueryInt(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid word Id"})
		return
	}
	zap.S().Infof("GET all word tenses for word %d", wordIdInt)
	words, err := controller.WordTenseUseCase.GetAllWordTenses(c, wordIdInt)
	if err != nil {
		zap.S().Error("Failed to get word tenses")
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		count := len(*words)
		zap.S().Debugf("Got word tenses %d", count)
		if count > 0 {
			c.JSON(http.StatusOK, &words)
		} else {
			c.JSON(http.StatusOK, []domainWordTense.WordTense{})
		}
	}
}

// Edit godoc
// @Summary      Edit word tense for word
// @Description  Update word tense for word
// @Tags         word_tense
// @Accept       json
// @Produce      json
// @Param input body domainWordTense.EditWordTenseRequest true "New data for word tense"
// @Success      200  {object}  domain.SuccessResponse
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/word/tense/edit [post]
func (controller *WordTenseController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	var request domainWordTense.EditWordTenseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	err := controller.WordTenseUseCase.Update(c, request.ID, request.WordId, request.TenseId, request.Original, request.Phonetic)
	if err != nil {
		zap.S().Error("Failed to update word tense with " + request.Original)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		zap.S().Debugf("Word tense updated %s", request.Original)
		c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Word tense updated"})
	}
}

// Create godoc
// @Summary Create word tense for word
// @Description Create new word tense for word
// @Tags word_tense
// @Accept  json
// @Produce  json
// @Param   input body domainWordomainWordTensedTag.WordTenseRequest true "Word tense data"
// @Success 201 {object} domain.SuccessIdResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/word/tag/create [post]
func (controller *WordTenseController) Create(c *gin.Context) {
	zap.S().Info("POST Create word tag")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	var request domainWordTense.WordTenseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	tenseId, err := controller.WordTenseUseCase.Create(c, request.WordId, request.TenseId, request.Original, request.Phonetic)
	if err != nil {
		zap.S().Error("Failed to create word tense with " + request.Original)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else if tenseId > 0 {
		zap.S().Debugf("Word created tense %s", request.Original)
		c.JSON(http.StatusCreated, domain.SuccessIdResponse{Id: tenseId})
	} else {
		c.JSON(http.StatusInternalServerError, "Something was wrong")
	}
}

// Delete godoc
// @Summary Delete word tense for word
// @Description Delete word tense for word
// @Tags word_tense
// @Accept  json
// @Produce  json
// @Param id path int true "ID word tense"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/word/tense/:id [delete]
func (controller *WordTenseController) Delete(c *gin.Context) {
	wordTenseId := c.Param("id")
	zap.S().Infof("DELETE Delete word tense %d", wordTenseId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	if wordTenseIdInt, err := strconv.Atoi(wordTenseId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid word ID"})
		return
	} else {
		rows, err := controller.WordTenseUseCase.DeleteById(c, wordTenseIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, wordTenseId, "Failed to delete word tense by", rows, err) {
			zap.S().Debugf("Word tense deleted %s", wordTenseId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Word tense deleted"})
		}
	}
}
