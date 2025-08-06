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
// - $2: dictionary id
func updateUserDictionaryQuery() string {
	return `
UPDATE dictionary
SET 
    dialect = $1
WHERE id = $2
RETURNING id, dialect, lang_from_id, lang_to_id;`
}

// DeleteUserDictionaryByIdQuery get query to delete dictionary by id from dictionary table
// Params:
// - $1: id
func deleteUserDictionaryByIdQuery() string {
	return `DELETE FROM dictionary WHERE id = $1`
}

// GetAllDictionariesWithShortInfoForUserQuery get query to get dictionaries with entities from dictionary and related tables
// Params:
// - $1: user uuid
func getAllDictionariesWithShortInfoForUserQuery() string {
	return `SELECT
  d.id AS dictionary_id,
  d.dialect,
  d.lang_from_id,
  d.lang_to_id,

  lf.id AS lang_from_id,
  lf.name AS lang_from_name,
  lf.code AS lang_from_code,

  lt.id AS lang_to_id,
  lt.name AS lang_to_name,
  lt.code AS lang_to_code,

  COUNT(DISTINCT w.id) AS word_count,
  COUNT(DISTINCT wt.id) AS word_tag_count,
  COUNT(DISTINCT q.id) AS quiz_count

FROM dictionary d
LEFT JOIN language lf ON d.lang_from_id = lf.id
LEFT JOIN language lt ON d.lang_to_id = lt.id
LEFT JOIN word w ON w.dictionary_id = d.id
LEFT JOIN word_tag wt ON wt.dictionary_id = d.id
LEFT JOIN quiz q ON q.dictionary_id = d.id
WHERE d.user_id = $1
GROUP BY d.id, d.dialect, d.lang_from_id, d.lang_to_id,
         lf.id, lf.name, lf.code,
         lt.id, lt.name, lt.code
ORDER BY d.id;
`
}
