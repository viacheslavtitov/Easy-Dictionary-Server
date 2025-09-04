package db

// GetAllWordsByDictionaryQuery get query to get all words for dictionary table
// Params:
// - $1: dictionary id
// - $2: last id from latest page
// - $3: page size
func getAllWordsByDictionaryQuery() string {
	return `
WITH paged_words AS (
  SELECT id
  FROM word
  WHERE dictionary_id = $1
  ORDER BY id
  LIMIT $3 OFFSET $2
)
SELECT
  w.id            AS word_id,
  w.dictionary_id AS word_dictionary_id,
  w.original      AS word_original,
  w.phonetic      AS word_phonetic,
  w.type          AS word_type,

  t.id            AS translation_id,
  t.description   AS translation_description,
  t.translate     AS translation_text,

  tc.id           AS category_id,
  tc.name         AS category_name
FROM paged_words pw
JOIN word w ON w.id = pw.id
LEFT JOIN translation t  ON t.word_id = w.id
LEFT JOIN translation_category tc
       ON tc.id = t.category_id
      AND tc.dictionary_id = w.dictionary_id
ORDER BY w.id, t.id;`
}

// GetSearchWordsByDictionaryQuery get query to get all words for dictionary table
// Params:
// - $1: dictionary id
// - $2: search query string by original column
// - $3: last id from latest page
// - $4: page size
func getSearchWordsByDictionaryQuery() string {
	return `WITH paged_words AS (
  SELECT id
  FROM word
  WHERE dictionary_id = $1
  AND original ILIKE '%' || $2 || '%'
  ORDER BY id
  LIMIT $4 OFFSET $3
)
SELECT
  w.id            AS word_id,
  w.dictionary_id AS word_dictionary_id,
  w.original      AS word_original,
  w.phonetic      AS word_phonetic,
  w.type          AS word_type,

  t.id            AS translation_id,
  t.description   AS translation_description,
  t.translate     AS translation_text,

  tc.id           AS category_id,
  tc.name         AS category_name
FROM paged_words pw
JOIN word w ON w.id = pw.id
LEFT JOIN translation t  ON t.word_id = w.id
LEFT JOIN translation_category tc
       ON tc.id = t.category_id
      AND tc.dictionary_id = w.dictionary_id
ORDER BY w.id, t.id;`
}

// CreateWordAndReturnIdQuery get query to create word
// Params:
// - $1: original
// - $2: phonetic
// - $3: type
// - $4: dictionary id
func createWordAndReturnIdQuery() string {
	return `
INSERT INTO word (original, phonetic, type, dictionary_id)
VALUES ($1, $2, $3, $4)
RETURNING id;
`
}

// CreateWordQuery get query to create word
// Params:
// - $1: original
// - $2: phonetic
// - $3: type
// - $4: dictionary id
func createWordQuery() string {
	return `
INSERT INTO word (original, phonetic, type, dictionary_id)
VALUES ($1, $2, $3, $4);
`
}

// UpdateWordQuery get query to update word
// Params:
// - $1: original
// - $2: phonetic
// - $3: type
// - $4: word id
func updateWordQuery() string {
	return `
UPDATE word
SET 
    original = $1,
    phonetic = $2,
    type = $3,
WHERE id = $4
RETURNING id, original, phonetic, type, dictionary_id;`
}

// DeleteWordByIdQuery get query to delete word by id from word table
// Params:
// - $1: id
func deleteWordByIdQuery() string {
	return `DELETE FROM word WHERE id = $1`
}
