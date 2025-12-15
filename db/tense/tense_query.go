package db

// GetAllDictionariesForUserQuery get query to get all dictionary tenses
// Params:
// - $1: dictionary id
func getAllTensesForDictionaryById() string {
	return `
SELECT 
    id AS id,
	name AS name
FROM tense
WHERE dictionary_id = $1;`
}

// CreateTenseQuery get query to create tense
// Params:
// - $1: dictionary id
// - $2: name
func CreateTenseQuery() string {
	return `
INSERT INTO tense (dictionary_id, name)
VALUES ($1, $2);
`
}

// UpdateTenseQuery get query to update tense
// Params:
// - $1: name
// - $2: id
func updateTenseQuery() string {
	return `
UPDATE tense
SET 
    name = $1
WHERE id = $2
RETURNING id, name;`
}

// DeleteTenseByIdQuery get query to delete tense by id
// Params:
// - $1: id
func deleteTenseByIdQuery() string {
	return `DELETE FROM tense WHERE id = $1`
}

// DeleteAllTensesByDictionaryIdQuery get query to delete tenses in dictionary
// Params:
// - $1: dictionary id
func DeleteAllTensesByDictionaryIdQuery() string {
	return `DELETE FROM tense WHERE dictionary_id = $1`
}

// BulkInsertTensesForDictionaryQuery get query to insert tenses for dictionary
// Params:
// - $1: dictionary id
// - $2: names
func BulkInsertTensesForDictionaryQuery() string {
	return `
INSERT INTO tense (dictionary_id, name)
SELECT $1, UNNEST($2::TEXT[]);`
}
