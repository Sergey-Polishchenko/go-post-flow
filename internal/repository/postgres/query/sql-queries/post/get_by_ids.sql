SELECT 
    p.id,
    p.title,
    p.content,
    p.allow_comments,
    u.id as author_id
FROM posts p
JOIN users u ON p.author_id = u.id
WHERE p.id = ANY($1)
