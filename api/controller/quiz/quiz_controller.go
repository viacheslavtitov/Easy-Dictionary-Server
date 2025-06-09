package controller

import (
	controllerCommon "easy-dictionary-server/api/controller"
	"easy-dictionary-server/domain"
	quizDomain "easy-dictionary-server/domain/quiz"
	validatorutil "easy-dictionary-server/internalenv/validator"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type QuizController struct {
	QuizUseCase quizDomain.QuizUseCase
}

func (controller *QuizController) GetAllQuizes(c *gin.Context) {
	zap.S().Info("GET GetAllQuizes")
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		quizes, err := controller.QuizUseCase.GetAllResultsByAllQuiz(c, *userID)
		if err != nil {
			zap.S().Error("Failed to get quizes")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Got quizes %d", len(quizes))
			c.JSON(http.StatusOK, &quizes)
		}
	}
}

func (controller *QuizController) GetQuizById(c *gin.Context) {
	quizId := c.Param("id")
	zap.S().Infof("GET GetQuizById %d", quizId)
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		if quizIdInt, err := strconv.Atoi(quizId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dictionary ID"})
			return
		} else {
			quizDetail, err := controller.QuizUseCase.GetAllResultsByQuizId(c, *userID, quizIdInt)
			if err != nil {
				zap.S().Error("Failed to get quiz")
				zap.S().Error(err)
				c.JSON(http.StatusInternalServerError, err.Error())
			} else {
				c.JSON(http.StatusOK, &quizDetail)
			}
		}
	}
}

func (controller *QuizController) Edit(c *gin.Context) {
	zap.S().Info("POST Edit")
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		var request quizDomain.EditQuizRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		_, err := controller.QuizUseCase.UpdateQuiz(c, request.ID, request.Name, request.Time)
		if err != nil {
			zap.S().Errorf("Failed to update quiz by id %d", request.ID)
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Quiz updated %d", request.ID)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Quiz updated"})
		}
	}
}

func (controller *QuizController) Create(c *gin.Context) {
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		zap.S().Infof("POST Create quiz for: %d", &userID)
		var request quizDomain.QuizRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		_, err := controller.QuizUseCase.CreateQuiz(c, request.DictionaryId, request.Name, request.Time, request.WordIds)
		if err != nil {
			zap.S().Error("Failed to create quiz ")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Quiz created")
			c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Quiz created"})
		}
	}
}

func (controller *QuizController) AddWord(c *gin.Context) {
	if userID, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	} else {
		zap.S().Infof("POST Add word to quiz for: %d", &userID)
		var request quizDomain.QuizWordRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			zap.S().Error(err)
			validationErrors := validatorutil.FormatValidationError(err)
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
			return
		}
		_, err := controller.QuizUseCase.AddWordToQuiz(c, request.QuizId, request.WordId)
		if err != nil {
			zap.S().Error("Failed to add word to quiz ")
			zap.S().Error(err)
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			zap.S().Debugf("Word added to quiz")
			c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Word added to quiz"})
		}
	}
}

func (controller *QuizController) Delete(c *gin.Context) {
	quizId := c.Param("id")
	zap.S().Infof("DELETE Delete quiz %d", quizId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	if quizIdInt, err := strconv.Atoi(quizId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	} else {
		rows, err := controller.QuizUseCase.DeleteQuizById(c, quizIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, quizId, "Failed to delete quiz by", rows, err) {
			zap.S().Debugf("Quiz deleted %s", quizId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Quiz deleted"})
		}
	}
}

func (controller *QuizController) DeleteWord(c *gin.Context) {
	quizWordId := c.Param("id")
	zap.S().Infof("DELETE Delete word from quiz %d", quizWordId)
	if _, _, valid := controllerCommon.ValidateUserIdInContext(c); !valid {
		return
	}
	if quizWordIdInt, err := strconv.Atoi(quizWordId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	} else {
		rows, err := controller.QuizUseCase.DeleteWordFromQuizById(c, quizWordIdInt)
		if controllerCommon.ValidateDeleteByIdResult(c, quizWordId, "Failed to delete quiz word by", rows, err) {
			zap.S().Debugf("Quiz word deleted %s", quizWordId)
			c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Quiz word deleted"})
		}
	}
}
