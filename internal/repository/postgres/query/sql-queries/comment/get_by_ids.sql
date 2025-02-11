SELECT 
    c.id,
    c.text,
    c.author_id,
    c.post_id,
    c.created_at
FROM comments c
WHERE c.id = ANY($1)
