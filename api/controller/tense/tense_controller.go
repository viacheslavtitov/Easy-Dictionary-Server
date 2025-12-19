package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
	"easy-dictionary-server/domain"
	tenseDomain "easy-dictionary-server/domain/dictionary"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TenseController struct {
	TenseUseCase tenseDomain.TenseUseCase
}

// Edit godoc
// @Summary      Edit tense
// @Description  Update tense
// @Tags         tense
// @Accept       json
// @Produce      json
// @Param input body tenseDomain.EditTenseRequest true "New data for tense"
// @Success      200  {object}  domain.SuccessResponse
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/tense/edit [post]
func (tenseController *TenseController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		var request tenseDomain.EditTenseRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		err := tenseController.TenseUseCase.Update(c, request.ID, request.Name)
		if err != nil {
			zap.S().Errorf("Failed to update tense by id %d", request.ID)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Tense updated %d", request.ID)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Tense updated"})
		}
	}
}

// Delete godoc
// @Summary Delete tense
// @Description Delete tense for dictionary
// @Tags tense
// @Accept  json
// @Produce  json
// @Param id path int true "ID tense"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/tense/:id [delete]
func (tenseController *TenseController) Delete(c *gin.Context) {
	tenseId := c.Param("id")
	zap.S().Infof("DELETE Delete tense %d", tenseId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	if tenseIdIdInt, err := strconv.Atoi(tenseId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tense ID"})
		return
	} else {
		rows, err := tenseController.TenseUseCase.DeleteById(c, tenseIdIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, tenseId, "Failed to delete tense by", rows, err) {
			zap.S().Debugf("Tense deleted %s", tenseId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Tense deleted"})
		}
	}
}
