SELECT id, title, content, author_id, allow_comments
FROM posts
ORDER BY id
LIMIT $1 OFFSET $2
