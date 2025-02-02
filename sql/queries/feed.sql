-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, name, url, user_id)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetFeeds :many
select f.*, u.name as user_name
from feeds f
         inner join users u
                    on f.user_id = u.id;

-- name: GetFeedByURL :one
select *
from feeds
where url = $1;

-- name: MarkFeedFetched :exec
update feeds
set updated_at      = $1,
    last_fetched_at =$2
where id = $3;

-- name: GetNextFeedToFetch :one
select *
from feeds
order by last_fetched_at
    nulls
    first, last_fetched_at
limit 1;
