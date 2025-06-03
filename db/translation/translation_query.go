package db

// GetAllTranslationsForWordQuery get query to get all translations for a word
// Params:
// - $1: word id
func getAllTranslationsForWordQuery() string {
	return `
SELECT 
    id AS id,
	word_id AS word_id,
	category_id AS category_id,
	translate AS translate,
	description AS description
FROM translation
WHERE word_id = $1;`
}

// CreateTranslationQuery get query to create translation for a word
// Params:
// - $1: word id
// - $2: category id
// - $3: translate
// - $4: description
func createTranslationQuery() string {
	return `
INSERT INTO translation (word_id, category_id, translate, description)
VALUES ($1, $2, $3, $4);
`
}

// UpdateTranslationQuery get query to update translation for a word
// Params:
// - $1: category id
// - $2: translate
// - $3: description
// - $4: id
func updateTranslationQuery() string {
	return `
UPDATE translation
SET 
    category_id = $1,
	translate = $2,
	description = $3
WHERE id = $4
RETURNING id, word_id, category_id, translate, description;`
}

// DeleteTranslationByIdQuery get query to delete translation from translation table
// Params:
// - $1: id
func deleteTranslationByIdQuery() string {
	return `DELETE FROM translation WHERE id = $1`
}
