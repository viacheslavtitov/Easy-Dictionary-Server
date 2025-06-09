package db

// GetAllResultsByAllQuiz get query to get all results from quiz table
// Params:
// - $1: user id
func getAllResultsByAllQuiz() string {
	return `
SELECT
    q.id,
    q.name,
    q.time,
	d.id AS dictionary_id,
    d.dialect AS dictionary_dialect,
    d.lang_from_id,
    d.lang_to_id,
    COUNT(DISTINCT qw.id) AS quiz_word_count,
    COUNT(DISTINCT qr.id) AS quiz_result_count,
    COUNT(DISTINCT qwr.id) AS quiz_word_result_count
FROM quiz q
JOIN dictionary d ON d.id = q.dictionary_id
LEFT JOIN quiz_word qw ON q.id = qw.quiz_id
LEFT JOIN quiz_result qr ON q.id = qr.quiz_id
LEFT JOIN quiz_word_result qwr ON qwr.quiz_result_id = qr.id
WHERE d.user_id = $1
GROUP BY q.id, q.name, q.time, d.id, d.dialect, d.lang_from_id, d.lang_to_id
ORDER BY q.id DESC;`
}

// GetAllResultsByQuizId get query to get all results from quiz table
// Params:
// - $1: quiz id
func getAllResultsByQuizId() string {
	return `
SELECT
	q.id,
    q.name,
    q.time,
    d.id AS dictionary_id,
    d.dialect AS dictionary_dialect,
    d.lang_from_id,
    d.lang_to_id,
	qw.id AS quiz_word_id,
	qw.word_id AS quiz_word_entity_id,
	qr.id AS quiz_result_id,
	qr.word_id AS quiz_result_word_id,
	qr.time AS quiz_result_time,
	qwr.id AS quiz_word_result_id,
	qwr.word_id AS quiz_word_entity_result_id,
	qwr.quiz_result_id AS quiz_result_entity_result_id,
	qwr.answer AS answer
FROM quiz q
JOIN dictionary d ON d.id = q.dictionary_id
LEFT JOIN quiz_word qw ON q.id = qw.quiz_id
LEFT JOIN quiz_result qr ON q.id = qr.quiz_id
LEFT JOIN quiz_word_result qwr ON qwr.quiz_result_id = qr.id
WHERE q.id = $1
ORDER BY q.id DESC;`
}

// CreateQuizQuery get query to create quiz
// Params:
// - $1: dictionary id
// - $2: name
// - $3: time
// Return:
// - id: created quiz id
func createQuizQuery() string {
	return `
WITH new_quiz AS (
    INSERT INTO quiz (dictionary_id, name, time)
    VALUES ($1, $2, $3)
    RETURNING id
)
SELECT id FROM new_quiz;
`
}

// UpdateQuizQuery get query to update user
// Params:
// - $1: quiz id
// - $2: name
// - $3: time
func updateQuizQuery() string {
	return `
UPDATE quiz
SET 
    name = $2,
    time = $3
WHERE id = $1
RETURNING id;`
}

// CreateQuizWordQuery get query to create quiz word
// Params:
// - $1: quiz id
// - $2: word id
// Return:
// - id: created quiz word id
func createQuizWordQuery() string {
	return `
WITH new_quiz_word AS (
    INSERT INTO quiz_word (quiz_id, word_id)
    VALUES ($1, $2)
    RETURNING id
)
SELECT id FROM new_quiz_word;
`
}

// CreateQuizWordResultQuery get query to create quiz word result
// Params:
// - $1: quiz result id
// - $2: word id
// - $3: answer
// Return:
// - id: created quiz word id
func createQuizWordResultQuery() string {
	return `
WITH new_quiz_word_result AS (
    INSERT INTO quiz_word_result (quiz_result_id, word_id, answer)
    VALUES ($1, $2, $3)
    RETURNING id
)
SELECT id FROM new_quiz_word_result;
`
}

// DeleteQuizByIdQuery get query to delete quiz by id from quiz table
// Params:
// - $1: id
func deleteQuizByIdQuery() string {
	return `DELETE FROM quiz WHERE id = $1`
}

// DeleteQuizByIdQuery get query to delete quiz word by id from quiz word table
// Params:
// - $1: id
func deleteQuizWordByIdQuery() string {
	return `DELETE FROM quiz_word WHERE id = $1`
}
