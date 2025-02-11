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

-- name: MarkFeedFetched :one
update feeds
set updated_at      = now(),
    last_fetched_at =now()
where id = $1
returning *;

-- name: GetNextFeedToFetch :one
select *
from feeds
order by last_fetched_at
    nulls
    first
limit 1;
