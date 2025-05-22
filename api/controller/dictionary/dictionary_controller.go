package controller

import (
	"easy-dictionary-server/domain"
	dictionaryDomain "easy-dictionary-server/domain/dictionary"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictionaryController struct {
	DictionaryUseCase dictionaryDomain.DictionaryUseCase
}

func (dictionaryController *DictionaryController) GetAllForUser(c *gin.Context) {
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
		dictionaries, err := dictionaryController.DictionaryUseCase.GetAllForUser(c, userIDInt)
		if err != nil {
			zap.S().Error("Failed to get languages")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Got languages %d", len(*dictionaries))
			c.JSON(http.StatusOK, &dictionaries)
		}
	}

}

func (dictionaryController *DictionaryController) Edit(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Info("POST Edit")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request dictionaryDomain.EditDictionaryRequest
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
		err := dictionaryController.DictionaryUseCase.Update(c, userIDInt, request.ID, request.Dialect, request.LangFromId, request.LangToId)
		if err != nil {
			zap.S().Errorf("Failed to update dictionary by id %d", request.ID)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Dictionary updated %d", request.ID)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Dictionary updated"})
		}
	}

}

func (dictionaryController *DictionaryController) Create(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Infof("POST Create dictionary for: %s", userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request dictionaryDomain.DictionaryRequest
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
		err := dictionaryController.DictionaryUseCase.Create(c, userIDInt, request.Dialect, request.LangFromId, request.LangToId)
		if err != nil {
			zap.S().Error("Failed to create dictionary ")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Dictionary created")
			c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Dictionary created"})
		}
	}
}

func (dictionaryController *DictionaryController) Delete(c *gin.Context) {
	userID, exists := c.Get("userID")
	dictionaryId := c.Param("id")
	zap.S().Infof("DELETE Delete dictionary %d for user: %s", dictionaryId, userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if dictionaryIdInt, err := strconv.Atoi(dictionaryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dictionary ID"})
		return
	} else {
		err := dictionaryController.DictionaryUseCase.DeleteById(c, dictionaryIdInt)
		if err != nil {
			zap.S().Error("Failed to delete dictionary by " + dictionaryId)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Dictionary deleted %s", dictionaryId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Dictionary deleted"})
		}
	}
}
