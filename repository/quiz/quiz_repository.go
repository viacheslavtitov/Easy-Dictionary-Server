package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbQuiz "easy-dictionary-server/db/quiz"
	domainQuiz "easy-dictionary-server/domain/quiz"
	quizMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type quizRepository struct {
	db *database.Database
}

func NewQuizRepository(db *database.Database) domainQuiz.QuizRepository {
	return &quizRepository{db: db}
}

func (ur *quizRepository) CreateQuiz(c context.Context, quiz domainQuiz.Quiz) (int, error) {
	zap.S().Debugf("Create quiz")
	row, err := dbQuiz.CreateQuiz(ur.db, quizMapper.FromQuizDomain(&quiz))
	zap.S().Debugf("Quiz created with %d id", row)
	return row, err
}

func (ur *quizRepository) AddWordToQuiz(context context.Context, quizId int, wordId int) (int, error) {
	zap.S().Debugf("AddWordToQuiz to quiz %d", quizId)
	row, err := dbQuiz.CreateQuizWord(ur.db, quizMapper.FromQuizWordDomain(&domainQuiz.QuizWord{
		WordId: wordId,
	}, quizId))
	return row, err
}

func (ur *quizRepository) UpdateQuiz(context context.Context, quiz domainQuiz.Quiz) (int, error) {
	zap.S().Debugf("UpdateQuiz %s", quiz.Name)
	row, err := dbQuiz.UpdateQuiz(ur.db, quizMapper.FromQuizDomain(&quiz))
	return row, err
}

func (ur *quizRepository) DeleteQuizById(context context.Context, id int) (int64, error) {
	zap.S().Debugf("DeleteQuizById %d", id)
	rowsDeleted, errQuery := dbQuiz.DeleteQuizById(ur.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}

func (ur *quizRepository) DeleteWordFromQuizById(context context.Context, quizWord domainQuiz.QuizWord) (int64, error) {
	zap.S().Debugf("DeleteWordFromQuizById %d", quizWord.ID)
	rowsDeleted, errQuery := dbQuiz.DeleteQuizWordById(ur.db, quizWord.ID)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}

func (ur *quizRepository) GetAllResultsByAllQuiz(context context.Context, userId int) ([]*domainQuiz.QuizDetail, error) {
	zap.S().Debugf("GetAllResultsByAllQuiz %d", userId)
	quizDetailEntities, err := dbQuiz.GetAllResultsByAllQuiz(ur.db, userId)
	quizDetails := []*domainQuiz.QuizDetail{}
	for _, item := range *quizDetailEntities {
		quizDetails = append(quizDetails, quizMapper.ToQuizDetailDomain(&item))
	}
	return quizDetails, err
}

func (ur *quizRepository) GetAllResultsByQuizId(context context.Context, userId int, quizId int) (*domainQuiz.QuizItemDetail, error) {
	zap.S().Debugf("GetAllResultsByQuizId %d %d", userId, quizId)
	quizItemDetailEntity, err := dbQuiz.GetAllResultsByQuizId(ur.db, userId, quizId)
	quizItemDetail := quizMapper.ToQuizItemDetailDomain(quizItemDetailEntity)
	return quizItemDetail, err
}
