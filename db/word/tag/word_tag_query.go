package db

// GetAllWordTagsByDictionaryQuery get query to get all word tags for dictionary
// Params:
// - $1: dictionary id
func getAllWordTagsByDictionaryQuery() string {
	return `
SELECT 
    id AS id,
	dictionary_id AS dictionary_id,
	name AS name
FROM word_tag
WHERE dictionary_id = $1;`
}

// CreateWordTagQuery get query to create word tag
// Params:
// - $1: name
// - $2: dictionary id
func createWordTagQuery() string {
	return `
INSERT INTO word_tag (name, dictionary_id)
VALUES ($1, $2);
`
}

// UpdateWordTagQuery get query to update word tag
// Params:
// - $1: name
// - $2: word tag id
func updateWordTagQuery() string {
	return `
UPDATE word_tag
SET 
    name = $1
WHERE id = $2
RETURNING id, dictionary_id, name;`
}

// DeleteWordTagByIdQuery get query to delete word by id from word table
// Params:
// - $1: id
func deleteWordTagByIdQuery() string {
	return `DELETE FROM word_tag WHERE id = $1`
}
