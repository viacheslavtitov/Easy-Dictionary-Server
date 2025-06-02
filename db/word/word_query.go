package db

// getAllWordsByDictionaryQuery get query to get all words for dictionary table
// Params:
// - $1: dictionary id
func getAllWordsByDictionaryQuery() string {
	return `
SELECT 
    id AS id,
	original AS original,
	phonetic AS phonetic,
	dictionary_id AS dictionary_id,
	type AS type,
	category_id AS category_id
FROM word
WHERE dictionary_id = $1;`
}

// CreateWordQuery get query to create word
// Params:
// - $1: original
// - $2: phonetic
// - $3: type
// - $4: category id
// - $5: dictionary id
func createWordQuery() string {
	return `
INSERT INTO word (original, phonetic, type, category_id, dictionary_id)
VALUES ($1, $2, $3, $4, $5);
`
}

// UpdateWordQuery get query to update word
// Params:
// - $1: original
// - $2: phonetic
// - $3: type
// - $4: category id
// - $5: word id
func updateWordQuery() string {
	return `
UPDATE word
SET 
    original = $1,
    phonetic = $2,
    type = $3,
    category_id = $4
WHERE id = $5
RETURNING id, original, phonetic, type, category_id, dictionary_id;`
}

// DeleteWordByIdQuery get query to delete word by id from word table
// Params:
// - $1: id
func deleteWordByIdQuery() string {
	return `DELETE FROM word WHERE id = $1`
}
