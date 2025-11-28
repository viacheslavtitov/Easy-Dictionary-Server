package db

// GetAllWordsByDictionaryQuery get query to get all words for dictionary table
// Params:
// - $1: dictionary id
// - $2: last id from latest page
// - $3: page size
// - $4: created_from TIMESTAMP or NULL
// - $5: created_to TIMESTAMP or NULL
// - $6: word_types TEXT[] or NULL
// - $7: category_ids INT[] or NULL
// - $8: tag_ids INT[] or null
// - $9: query by original field (text, nullable/empty)
func getAllWordsByDictionaryQuery() string {
	return `
WITH page AS (
  SELECT id
  FROM word
  WHERE dictionary_id = $1
    AND id > $2               -- $2 = lastId (0 for first page)
    AND ($4::timestamp IS NULL OR created_at >= $4::timestamp) -- $4 created_from
    AND ($5::timestamp IS NULL OR created_at <= $5::timestamp) -- $5 created_to
    AND (
      $6::text[] IS NULL
      OR type = ANY($6::text[])          -- $6 word_types
    )
    AND (
      $9::text IS NULL
      OR original ILIKE '%' || $9::text || '%' -- $9
    )
  ORDER BY id
  LIMIT $3                    -- $3 = pageSize (+1, if you want to know has_more)
)
SELECT
  w.id            AS word_id,
  w.dictionary_id AS word_dictionary_id,
  w.original      AS word_original,
  w.phonetic      AS word_phonetic,
  w.type          AS word_type,
  w.created_at AS word_created_at,

  t.id            AS translation_id,
  t.description   AS translation_description,
  t.translate     AS translation_text,
  t.created_at AS translation_created_at,

  tc.id           AS category_id,
  tc.name         AS category_name,

  wt.id           AS tag_id,
  wt.name         AS tag_name
FROM word w
JOIN page p           ON p.id = w.id
LEFT JOIN translation t
       ON t.word_id = w.id
LEFT JOIN translation_category tc
       ON tc.id = t.category_id
      AND tc.dictionary_id = w.dictionary_id
LEFT JOIN word_tag_word wtw
       ON wtw.word_id = w.id
LEFT JOIN word_tag wt
       ON wt.id = wtw.word_tag_id
      AND wt.dictionary_id = w.dictionary_id
WHERE
  (
    -- 1) if all filters are null
    $6::text[] IS NULL
    AND $7::int[] IS NULL
    AND $8::int[] IS NULL
  )
  OR (
    -- 2) find in word_types
    $6::text[] IS NOT NULL
    AND w.type = ANY($6::text[])
  )
  OR (
    -- 3) find in category_ids
    $7::int[] IS NOT NULL
    AND tc.id = ANY($7::int[])
  )
  OR (
    -- 4) find in tag_ids
    $8::int[] IS NOT NULL
    AND wt.id = ANY($8::int[])
  )
ORDER BY w.id, t.id, wt.id;`
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
    type = $3
WHERE id = $4;`
}

// DeleteWordByIdQuery get query to delete word by id from word table
// Params:
// - $1: id
func deleteWordByIdQuery() string {
	return `DELETE FROM word WHERE id = $1`
}
