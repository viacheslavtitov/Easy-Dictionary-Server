package controller

import (
	"easy-dictionary-server/domain"
	domainWord "easy-dictionary-server/domain/word"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WordController struct {
	WordUseCase domainWord.WordUseCase
}

func (controller *WordController) GetAllForDictionary(c *gin.Context) {
	userID, exists := c.Get("userID")
	dictionaryId := c.Param("id")
	zap.S().Info("GET GetAllForDictionary")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	_, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	if dictionaryIdInt, err := strconv.Atoi(dictionaryId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dictionary ID"})
		return
	} else {
		words, err := controller.WordUseCase.GetAllForDictionary(c, dictionaryIdInt)
		if err != nil {
			zap.S().Error("Failed to get words")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Got words %d", len(*words))
			c.JSON(http.StatusOK, &words)
		}
	}
}

func (controller *WordController) Edit(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Info("POST Edit")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request domainWord.EditWordRequest
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
		err := controller.WordUseCase.Update(c, request.ID, request.DictionaryId, request.Original, &request.Phonetic, request.Type, &request.CategoryId)
		if err != nil {
			zap.S().Error("Failed to update word with " + request.Original)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Word updated %s", request.Original)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Word updated"})
		}
	}

}

func (controller *WordController) Create(c *gin.Context) {
	userID, exists := c.Get("userID")
	zap.S().Infof("POST word for: %s", userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var request domainWord.WordRequest
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
		err := controller.WordUseCase.Create(c, request.DictionaryId, request.Original, &request.Phonetic, request.Type, &request.CategoryId)
		if err != nil {
			zap.S().Error("Failed to create word with " + request.Original)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Word created %s", request.Original)
			c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Word created"})
		}
	}
}

func (controller *WordController) Delete(c *gin.Context) {
	userID, exists := c.Get("userID")
	wordId := c.Param("id")
	zap.S().Infof("DELETE Delete word %d for user: %s", wordId, userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if wordIdInt, err := strconv.Atoi(wordId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID"})
		return
	} else {
		err := controller.WordUseCase.DeleteById(c, wordIdInt)
		if err != nil {
			zap.S().Error("Failed to delete word by " + wordId)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Word deleted %s", wordId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Word deleted"})
		}
	}
}
