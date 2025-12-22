package db

// CreateWordTenseAndReturnIdQuery get query to create tense for word
// Params:
// - $1: tense id
// - $2: word id
// - $3: original
// - $4: phonetic
func CreateWordTenseAndReturnIdQuery() string {
	return `
INSERT INTO word_tense (tense_id, word_id, original, phonetic)
VALUES ($1, $2, $3, $4)
RETURNING id;
`
}

// GetAllWordTensesByQuery get query to get all tenses for word
// Params:
// - $1: word id
func getAllWordTensesByWordQuery() string {
	return `
SELECT
    wt.id AS id,
	wt.tense_id AS tense_id,
	wt.word_id AS word_id,
	wt.original AS original,
	wt.phonetic AS phonetic
FROM word_tense AS wt
WHERE wt.word_id = $1;`
}

// UpdateWordTenseQuery get query to update word tense
// Params:
// - $1: original
// - $2: phonetic
// - $3: word tense id
func updateWordTenseQuery() string {
	return `
UPDATE word_tense
SET 
    original = $1,
    phonetic = $2
WHERE id = $3;`
}

// DeleteWordTenseByIdQuery get query to delete word tense by id from word_tense table
// Params:
// - $1: id
func deleteWordTenseByIdQuery() string {
	return `DELETE FROM word_tense WHERE id = $1`
}
