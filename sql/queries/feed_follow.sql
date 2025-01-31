-- name: CreateFeedFollow :one
with inserted_feed_follow as (
    insert into feed_follow (id, created_at, updated_at, user_id, feed_id)
        values ($1, $2, $3, $4, $5)
        returning *)
select iff.*, f.name as feed_name, u.name as user_name
from inserted_feed_follow iff
         inner join feeds f on iff.feed_id = f.id
         inner join users u on iff.user_id = u.id;

-- name: GetFeedFollowsForUser :many
select ff.*, f.name as feed_name, u.name as user_name
from feed_follow ff
         inner join feeds f on ff.feed_id = f.id
         inner join users u on ff.user_id = u.id
where ff.user_id = $1;