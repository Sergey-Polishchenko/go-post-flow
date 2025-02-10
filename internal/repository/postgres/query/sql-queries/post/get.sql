SELECT id, title, content, author_id, allow_comments
FROM posts
WHERE id = $1
