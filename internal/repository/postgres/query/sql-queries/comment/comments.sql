SELECT c.id
FROM comments c
JOIN users u ON c.author_id = u.id
WHERE c.post_id = $1 AND c.parent_id IS NULL
