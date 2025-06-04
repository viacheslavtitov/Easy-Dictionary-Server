package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
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
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	dictionaryIdInt, err := controllerCommon.ParseQueryInt(c, "dictionaryId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dictionary Id"})
		return
	}
	lastIdInt, err := controllerCommon.ParseQueryInt(c, "lastId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid last read Id"})
		return
	}
	pageSizeInt, err := controllerCommon.ParseQueryInt(c, "pageSize")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}
	zap.S().Infof("GET all words for dictionary %d with lastId %d and pageSize %d", dictionaryIdInt, lastIdInt, pageSizeInt)
	words, err := controller.WordUseCase.GetAllForDictionary(c, dictionaryIdInt, lastIdInt, pageSizeInt)
	if err != nil {
		zap.S().Error("Failed to get words")
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		count := len(*words)
		zap.S().Debugf("Got words %d", count)
		if count > 0 {
			last := (*words)[count-1].ID
			c.JSON(http.StatusOK, domainWord.WordsWithPaginationResponse{
				Words:    *words,
				LatestId: last,
				PageSize: pageSizeInt,
			})
		} else {
			c.JSON(http.StatusOK, domainWord.WordsWithPaginationResponse{
				Words:    []domainWord.Word{},
				LatestId: 0,
				PageSize: 0,
			})
		}
	}
}

func (controller *WordController) SearchForDictionary(c *gin.Context) {
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	dictionaryIdInt, err := controllerCommon.ParseQueryInt(c, "dictionaryId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dictionary Id"})
		return
	}
	lastIdInt, err := controllerCommon.ParseQueryInt(c, "lastId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid last read Id"})
		return
	}
	pageSizeInt, err := controllerCommon.ParseQueryInt(c, "pageSize")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}
	query := c.Query("query")
	if len(query) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query is empty"})
		return
	}
	zap.S().Infof("GET search words in dictionary %d with lastId %d and pageSize %d and query %s", dictionaryIdInt, lastIdInt, pageSizeInt, query)
	words, err := controller.WordUseCase.SearchWordsForDictionary(c, query, dictionaryIdInt, lastIdInt, pageSizeInt)
	if err != nil {
		zap.S().Error("Failed to get words")
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		count := len(*words)
		zap.S().Debugf("Got words %d", count)
		if count > 0 {
			last := (*words)[count-1].ID
			c.JSON(http.StatusOK, domainWord.WordsWithPaginationResponse{
				Words:    *words,
				LatestId: last,
				PageSize: pageSizeInt,
			})
		} else {
			c.JSON(http.StatusOK, domainWord.WordsWithPaginationResponse{
				Words:    []domainWord.Word{},
				LatestId: 0,
				PageSize: 0,
			})
		}
	}
}

func (controller *WordController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	var request domainWord.EditWordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	err := controller.WordUseCase.Update(c, request.ID, request.DictionaryId, request.Original, request.Phonetic, request.Type, request.CategoryId)
	if err != nil {
		zap.S().Error("Failed to update word with " + request.Original)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		zap.S().Debugf("Word updated %s", request.Original)
		c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Word updated"})
	}
}

func (controller *WordController) Create(c *gin.Context) {
	zap.S().Info("POST Create word")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	var request domainWord.WordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	err := controller.WordUseCase.Create(c, request.DictionaryId, request.Original, request.Phonetic, request.Type, request.CategoryId)
	if err != nil {
		zap.S().Error("Failed to create word with " + request.Original)
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		zap.S().Debugf("Word created %s", request.Original)
		c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Word created"})
	}
}

func (controller *WordController) Delete(c *gin.Context) {
	wordId := c.Param("id")
	zap.S().Infof("DELETE Delete word %d", wordId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	if wordIdInt, err := strconv.Atoi(wordId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid word ID"})
		return
	} else {
		rows, err := controller.WordUseCase.DeleteById(c, wordIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, wordId, "Failed to delete word by", rows, err) {
			zap.S().Debugf("Word deleted %s", wordId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Word deleted"})
		}
	}
}
