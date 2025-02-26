INSERT INTO comments (text, author_id, post_id, parent_id)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at
