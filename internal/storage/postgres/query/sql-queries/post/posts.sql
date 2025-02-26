SELECT 
    p.id,
    p.title,
    p.content,
    p.author_id,
    p.allow_comments
FROM posts p
ORDER BY p.id DESC
LIMIT $1 OFFSET $2
