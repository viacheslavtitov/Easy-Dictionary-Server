package db

// GetAllDictionariesForUserQuery get query to get all user dictionaries from dictionary table
func getAllDictionariesForUserQuery() string {
	return `
SELECT 
    id AS id,
	user_id AS user_id,
	dialect AS dialect,
	lang_from_id AS lang_from_id,
	lang_to_id AS lang_to_id
FROM dictionary
WHERE user_id = $1;`
}

// CreateUserDictionaryQuery get query to create dictionary
// Params:
// - $1: dialect
// - $2: language from id
// - $3: language to id
// - $4: user id
func createUserDictionaryQuery() string {
	return `
INSERT INTO dictionary (dialect, lang_from_id, lang_to_id, user_id)
VALUES ($1, $2, $3, $4);
`
}

// UpdateUserDictionaryQuery get query to update dictionary
// Params:
// - $1: dialect
// - $2: language from id
// - $3: language to id
// - $4: dictionary id
func updateUserDictionaryQuery() string {
	return `
UPDATE dictionary
SET 
    dialect = $1,
    lang_from_id = $2
	lang_to_id = $3
WHERE id = $4
RETURNING id, dialect, lang_from_id, lang_to_id;`
}

// DeleteUserDictionaryByIdQuery get query to delete dictionary by id from dictionary table
// Params:
// - $1: id
func deleteUserDictionaryByIdQuery() string {
	return `DELETE FROM dictionary WHERE id = $1`
}
