-- name: CreateFeedFollow :one
insert into feed_follow (id, created_at, updated_at, user_id, feed_id)
values ($1,$2,$3,$4,$5)
returning *;

-- name: GetFeedFollowsForUser :many
select f.name from feed_follow ff
         inner join feeds f
         on  ff.feed_id = f.id
where ff.user_id = $1;