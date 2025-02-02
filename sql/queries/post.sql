-- name: CreatePost :one
insert into posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning *;

-- name: GetPostsForUser :many
select p.*, f.name as feed_name
from posts p
         inner join feed_follow ff on ff.feed_id = p.feed_id
         inner join feeds f on f.id = p.feed_id
where ff.user_id = $1
order by p.published_at desc
limit $2;