package db

// getAllTranslationCategoriesForUserQuery get query to get all translation categories from all dictionaries
// Params:
// - $1: user id
func getAllTranslationCategoriesForUserQuery() string {
	return `
SELECT tc.*
FROM translation_category tc
JOIN dictionary d ON tc.dictionary_id = d.id
WHERE d.user_id = $1;
`
}

// CreateUserTranslationCategoryQuery get query to create translation category
// Params:
// - $1: name
// - $2: user id
// - $3: dictionary id
func createUserTranslationCategoryQuery() string {
	return `
INSERT INTO translation_category (name, user_id, dictionary_id)
VALUES ($1, $2, $3);
`
}

// UpdateUserTranslationCategoryQuery get query to update translation category
// Params:
// - $1: name
// - $2: id
func updateUserTranslationCategoryQuery() string {
	return `
UPDATE translation_category
SET 
    name = $1
WHERE id = $2
RETURNING id, name, user_id, dictionary_id;`
}

// DeleteUserLanguageByIdQuery get query to delete translation category by id from language table
// Params:
// - $1: id
func deleteUserTranslationCategoryByIdQuery() string {
	return `DELETE FROM translation_category WHERE id = $1`
}
