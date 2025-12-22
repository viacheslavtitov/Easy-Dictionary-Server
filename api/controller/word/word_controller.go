package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
	"easy-dictionary-server/domain"
	domainWord "easy-dictionary-server/domain/word"
	validatorutil "easy-dictionary-server/internalenv/validator"
	repositoryWord "easy-dictionary-server/repository/word"
	"errors"
	"time"

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
// @Param   query    query     string     false     "Filter by word"
// @Param   categoryIds    query     []int     false     "Filter by categories"
// @Param   tagIds    query     []int     false     "Filter by tags"
// @Param   wordTypes    query     []string     false     "Filter by word types"
// @Param   from    query     string     false     "Filter by time from (DateOnly)" Format(date-time)
// @Param   to    query     string     false     "Filter by time to (DateOnly)" Format(date-time)
// @Success      200  {object}  domainWord.WordsWithPaginationResponse
// @Failure      400  {object}  domain.ErrorResponse
// @Failure      404  {object}  domain.ErrorResponse
// @Failure      500  {object}  domain.ErrorResponse
// @Router       /api/word/all [get]
func (controller *WordController) GetAllForDictionary(c *gin.Context) {
	userId, _, valid := controllerCommon.ValidateUserIdInContext(c)
	if !valid {
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

	categoryIds, err := controllerCommon.GetIntSliceParam(c, "categoryIds")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if categoryIds != nil {
		zap.S().Infow("Accept filter by categoryIds", categoryIds)
	}
	tagIds, err := controllerCommon.GetIntSliceParam(c, "tagIds")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if tagIds != nil {
		zap.S().Infow("Accept filter by tagIds", tagIds)
	}
	wordTypes, _ := controllerCommon.GetStringSliceParam(c, "wordTypes")
	zap.S().Infow("Accept filter by wordTypes", wordTypes)
	createdFrom, err := controllerCommon.ParseDateTime(c, "from")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if createdFrom != nil {
		zap.S().Infof("Accept filter by createdFrom = %v", createdFrom.Format(time.DateOnly))
	}
	createdTo, err := controllerCommon.ParseDateTime(c, "to")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if createdTo != nil {
		zap.S().Infof("Accept filter by createdTo = %v", createdTo.Format(time.DateOnly))
	}
	query := c.Query("query")
	zap.S().Infof("Accept filter by query = %s", query)

	zap.S().Infof("GET all words for dictionary %d with lastId %d and pageSize %d", dictionaryIdInt, lastIdInt, pageSizeInt)
	wordsResponse, err := controller.WordUseCase.SearchWordsForDictionary(c, *userId, query, dictionaryIdInt, lastIdInt, pageSizeInt,
		createdFrom, createdTo, &wordTypes, categoryIds, tagIds)
	if err != nil {
		zap.S().Error("Failed to get words")
		zap.S().Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		count := len(wordsResponse.Words)
		zap.S().Debugf("Got words %d", count)
		c.JSON(http.StatusOK, wordsResponse)
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
	err := controller.WordUseCase.Update(c, request.ID, request.DictionaryId, request.Original, request.Phonetic, request.Type, request.TagIds)
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
	err := controller.WordUseCase.Create(c, request.DictionaryId, request.Original, request.Phonetic, request.Type)
	if err != nil {
		if errors.Is(err, repositoryWord.ErrWordAlreadyExists) {
			zap.S().Error(err)
			c.JSON(http.StatusConflict, err.Error())
		} else {
			zap.S().Error("Failed to create word with " + request.Original)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}
	} else {
		zap.S().Debugf("Word created %s", request.Original)
		c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Word created"})
	}
}

// Create godoc
// @Summary Create word with translations
// @Description Create new word with translations for dictionary
// @Tags word
// @Accept  json
// @Produce  json
// @Param   input body domainWord.WordWithTranslationRequest true "Word data"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Router /api/word/create/translations [post]
func (controller *WordController) CreateWordWithTranslations(c *gin.Context) {
	zap.S().Info("POST Create word with translations")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	var request domainWord.WordWithTranslationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		zap.S().Error(err)
		validationErrors := validatorutil.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}
	err := controller.WordUseCase.CreateWithTranslations(c, request.DictionaryId, request.Original, request.Phonetic, request.Type, request.Translations,
		request.WordTags, request.WordTenses)
	if err != nil {
		if errors.Is(err, repositoryWord.ErrWordAlreadyExists) {
			zap.S().Error(err)
			c.JSON(http.StatusConflict, err.Error())
		} else {
			zap.S().Error("Failed to create word with " + request.Original)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		}
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
