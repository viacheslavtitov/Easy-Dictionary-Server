package usecase

import (
	"context"
	"time"

	domainQuiz "easy-dictionary-server/domain/quiz"
	commonUseCase "easy-dictionary-server/usecase"
)

type quizUsecase struct {
	quizRepository domainQuiz.QuizRepository
	contextTimeout int
}

func NewQuizUsecase(quizRepository domainQuiz.QuizRepository, timeout int) domainQuiz.QuizUseCase {
	return &quizUsecase{
		quizRepository: quizRepository,
		contextTimeout: timeout,
	}
}

func (usecase *quizUsecase) CreateQuiz(c context.Context, dictionaryId int, name string, time *time.Time, worIds *[]int) (int, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(usecase.contextTimeout))
	defer cancel()
	return usecase.quizRepository.CreateQuiz(ctx, domainQuiz.Quiz{
		DictionaryId: dictionaryId,
		Name:         name,
		Time:         time,
	})
}

func (usecase *quizUsecase) AddWordToQuiz(c context.Context, quizId int, wordId int) (int, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(usecase.contextTimeout))
	defer cancel()
	return usecase.quizRepository.AddWordToQuiz(ctx, quizId, wordId)
}

func (usecase *quizUsecase) UpdateQuiz(c context.Context, id int, name string, time *time.Time) (int, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(usecase.contextTimeout))
	defer cancel()
	return usecase.quizRepository.UpdateQuiz(ctx, domainQuiz.Quiz{
		ID:   id,
		Name: name,
		Time: time,
	})
}

func (usecase *quizUsecase) DeleteQuizById(c context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(usecase.contextTimeout))
	defer cancel()
	return usecase.quizRepository.DeleteQuizById(ctx, id)
}

func (usecase *quizUsecase) DeleteWordFromQuizById(c context.Context, quizWordId int) (int64, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(usecase.contextTimeout))
	defer cancel()
	return usecase.quizRepository.DeleteWordFromQuizById(ctx, domainQuiz.QuizWord{
		ID: quizWordId,
	})
}

func (usecase *quizUsecase) GetAllResultsByAllQuiz(c context.Context, userId int) ([]*domainQuiz.QuizDetail, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(usecase.contextTimeout))
	defer cancel()
	return usecase.quizRepository.GetAllResultsByAllQuiz(ctx, userId)
}

func (usecase *quizUsecase) GetAllResultsByQuizId(c context.Context, userId int, quizId int) (*domainQuiz.QuizItemDetail, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(usecase.contextTimeout))
	defer cancel()
	return usecase.quizRepository.GetAllResultsByQuizId(ctx, userId, quizId)
}
