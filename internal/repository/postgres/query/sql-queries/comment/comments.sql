SELECT c.id, c.text, c.author_id, u.name AS author_name, c.created_at, c.post_id
FROM comments c
JOIN users u ON c.author_id = u.id
WHERE c.post_id = $1 AND c.parent_id IS NULL
