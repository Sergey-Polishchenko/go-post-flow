SELECT c.id, c.text, c.author_id, u.name AS author_name, c.created_at, c.post_id
FROM comments c
JOIN users u ON c.author_id = u.id
WHERE c.parent_id = $1
