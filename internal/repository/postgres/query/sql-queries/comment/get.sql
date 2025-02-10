SELECT c.id, c.text, c.author_id, u.name AS author_name, c.post_id, c.created_at
FROM comments c
JOIN users u ON c.author_id = u.id
WHERE c.id = $1
