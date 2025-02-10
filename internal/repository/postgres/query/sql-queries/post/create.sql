INSERT INTO posts (title, content, author_id, allow_comments)
VALUES ($1, $2, $3, $4)
RETURNING id
