SELECT c.id
FROM comments c
JOIN users u ON c.author_id = u.id
WHERE c.parent_id = $1
