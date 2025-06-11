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

// GetAllForDictionary godoc
// @Summary      Get all words for dictionary
// @Description  Get all words for dictionary
// @Tags         word
// @Accept       json
// @Produce      json
// @Param   dictionaryId    query     int     true     "ID dictionary"
// @Param   lastId    query     int     true     "Last id in the previous response"
// @Param   pageSize    query     int     true     "Size of items in response"
// @Success      200  {object}  domainWord.WordsWithPaginationResponse
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/word/all [get]
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

// SearchForDictionary godoc
// @Summary      Get all words for dictionary by search query
// @Description  Get all words for dictionary by search query
// @Tags         word
// @Accept       json
// @Produce      json
// @Param   dictionaryId    query     int     true     "ID dictionary"
// @Param   lastId    query     int     true     "Last id in the previous response"
// @Param   pageSize    query     int     true     "Size of items in response"
// @Param   query    query     int     true     "Search letters"
// @Success      200  {object}  domainWord.WordsWithPaginationResponse
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/word/search [get]
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

// Edit godoc
// @Summary      Edit word in dictionary
// @Description  Update word in dictionary
// @Tags         word
// @Accept       json
// @Produce      json
// @Param input body domainWord.EditWordRequest true "New data for word"
// @Success      200  {object}  domain.SuccessResponse
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/word/edit [post]
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

// Create godoc
// @Summary Create word
// @Description Create new word for dictionary
// @Tags word
// @Accept  json
// @Produce  json
// @Param   input body domainWord.WordRequest true "Word data"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/word/create [post]
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

// Delete godoc
// @Summary Delete word
// @Description Delete word for dictionary
// @Tags word
// @Accept  json
// @Produce  json
// @Param id path int true "ID word"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/word/:id [delete]
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
