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

// GetAllWordTagsByDictionaryQuery get query to get all word tags for dictionary
// Params:
// - $1: word id
func getAllWordTagsByWordQuery() string {
	return `
SELECT
    wt.id AS id,
	wt.dictionary_id AS dictionary_id,
	wt.name AS name

	wtw.word_id AS word_id,
	wtw.id AS word_tag_word_id,
FROM word_tag AS wt
JOIN word_tag_word AS wtw
  ON wt.id = wtw.word_tag_id
WHERE wtw.word_id = $1;`
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

// AddWordTagToWordQuery get query to create word tag
// Params:
// - $1: word tag id
// - $2: word id
func AddWordTagToWordQuery() string {
	return `
INSERT INTO word_tag_word (word_tag_id, word_id)
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
