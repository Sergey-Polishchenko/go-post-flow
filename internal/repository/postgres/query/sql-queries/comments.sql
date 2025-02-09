-- name: CreateComment
INSERT INTO comments (text, author_id, post_id, parent_id)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetComments
SELECT id, text, author_id, created_at
FROM comments
WHERE post_id = $1 AND parent_id IS NULL;
