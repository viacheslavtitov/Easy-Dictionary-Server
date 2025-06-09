package mapper

import (
	dbQuiz "easy-dictionary-server/db/quiz"
	domainDictionary "easy-dictionary-server/domain/dictionary"
	domainQuiz "easy-dictionary-server/domain/quiz"
)

// Domain
func ToQuizDomain(q *dbQuiz.QuizEntity) *domainQuiz.Quiz {
	return &domainQuiz.Quiz{
		ID:           q.ID,
		DictionaryId: q.DictionaryId,
		Name:         q.Name,
		Time:         q.Time,
	}
}

func ToQuizDetailDomain(q *dbQuiz.QuizDetailEntity) *domainQuiz.QuizDetail {
	return &domainQuiz.QuizDetail{
		QuizItem: domainQuiz.Quiz{
			ID:           q.QuizItem.ID,
			DictionaryId: q.QuizItem.DictionaryId,
			Name:         q.QuizItem.Name,
			Time:         q.QuizItem.Time,
		},
		DictionaryItem: domainDictionary.Dictionary{
			ID:         q.DictionaryItem.ID,
			Dialect:    q.DictionaryItem.Dialect,
			LangFromId: q.DictionaryItem.LangFromId,
			LangToId:   q.DictionaryItem.LangToId,
		},
		QuizWordCount:       q.QuizWordCount,
		QuizWordResultCount: q.QuizWordCount,
		QuizResultCount:     q.QuizResultCount,
	}
}

func ToQuizItemDetailDomain(q *dbQuiz.QuizItemDetailEntity) *domainQuiz.QuizItemDetail {
	return &domainQuiz.QuizItemDetail{
		QuizItem: domainQuiz.Quiz{
			ID:           q.QuizItem.ID,
			DictionaryId: q.QuizItem.DictionaryId,
			Name:         q.QuizItem.Name,
			Time:         q.QuizItem.Time,
		},
		DictionaryItem: domainDictionary.Dictionary{
			ID:         q.DictionaryItem.ID,
			Dialect:    q.DictionaryItem.Dialect,
			LangFromId: q.DictionaryItem.LangFromId,
			LangToId:   q.DictionaryItem.LangToId,
		},
		QuizWords:        ToQuizWordsDomain(q.QuizWords),
		QuizResult:       ToQuizResultDomain(q.QuizResult),
		QuizWordsResults: ToQuizWordsResultDomain(q.QuizWordsResults),
	}
}

func ToQuizWordDomain(q *dbQuiz.QuizWordEntity) *domainQuiz.QuizWord {
	return &domainQuiz.QuizWord{
		ID:     q.ID,
		WordId: q.WordId,
	}
}

func ToQuizWordsDomain(q *[]dbQuiz.QuizWordEntity) *[]domainQuiz.QuizWord {
	if q == nil {
		return &[]domainQuiz.QuizWord{}
	}
	quizWords := make([]domainQuiz.QuizWord, len(*q))
	for i, item := range *q {
		quizWords[i] = *ToQuizWordDomain(&item)
	}
	return &quizWords
}

func ToQuizWordResultDomain(q *dbQuiz.QuizWordResultEntity) *domainQuiz.QuizWordResult {
	return &domainQuiz.QuizWordResult{
		ID:           q.ID,
		WordId:       q.WordId,
		QuizResultId: q.QuizResultId,
		Answer:       q.Answer,
	}
}

func ToQuizWordsResultDomain(q *[]dbQuiz.QuizWordResultEntity) *[]domainQuiz.QuizWordResult {
	if q == nil {
		return &[]domainQuiz.QuizWordResult{}
	}
	quizWordsResult := make([]domainQuiz.QuizWordResult, len(*q))
	for i, item := range *q {
		quizWordsResult[i] = *ToQuizWordResultDomain(&item)
	}
	return &quizWordsResult
}

func ToQuizResultDomain(q *dbQuiz.QuizResultEntity) *domainQuiz.QuizResult {
	return &domainQuiz.QuizResult{
		ID:     q.ID,
		WordId: q.WordId,
		Time:   q.Time,
	}
}

// DB
func FromQuizDomain(q *domainQuiz.Quiz) *dbQuiz.QuizEntity {
	return &dbQuiz.QuizEntity{
		ID:           q.ID,
		DictionaryId: q.DictionaryId,
		Name:         q.Name,
		Time:         q.Time,
	}
}

func FromQuizWordDomain(q *domainQuiz.QuizWord, quizId int) *dbQuiz.QuizWordEntity {
	return &dbQuiz.QuizWordEntity{
		ID:     q.ID,
		QuizId: quizId,
		WordId: q.WordId,
	}
}
