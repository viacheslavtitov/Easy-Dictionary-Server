package db

// GetAllWordsByDictionaryQuery get query to get all words for dictionary table
// Params:
// - $1: dictionary id
// - $2: last id from latest page
// - $3: page size
func getAllWordsByDictionaryQuery() string {
	return `
SELECT * FROM word
WHERE dictionary_id = $1 AND id > $2
ORDER BY id
LIMIT $3;`
}

// GetSearchWordsByDictionaryQuery get query to get all words for dictionary table
// Params:
// - $1: dictionary id
// - $2: search query string by original column
// - $3: last id from latest page
// - $4: page size
func getSearchWordsByDictionaryQuery() string {
	return `
SELECT * FROM word
WHERE dictionary_id = $1
  AND original ILIKE '%' || $2 || '%'
  AND id > $3
ORDER BY id
LIMIT $4;`
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
