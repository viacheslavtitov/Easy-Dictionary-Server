package db

// GetAllLanguagesForUserQuery get query to get all user languages from user_languages table
func getAllLanguagesForUserQuery() string {
	return `
SELECT 
    id AS id,
	name AS name,
	code AS code,
FROM user_languages
WHERE id = $1;`
}

// CreateUserLanguageQuery get query to create language
// Params:
// - $1: name
// - $2: code
// - $3: user id
func createUserLanguageQuery() string {
	return `
INSERT INTO user_languages (name, code, user_id)
VALUES ($1, $2, $3);
`
}

// UpdateUserLanguageQuery get query to update language
// Params:
// - $1: name
// - $2: code
// - $3: language id
func updateUserLanguageQuery() string {
	return `
UPDATE user_languages
SET 
    name = $1,
    code = $2
WHERE id = $3
RETURNING id, code, name;`
}

// DeleteUserLanguageByIdQuery get query to delete language by id from user_languages table
// Params:
// - $1: id
func deleteUserLanguageByIdQuery() string {
	return `DELETE FROM user_languages WHERE id = $1`
}
