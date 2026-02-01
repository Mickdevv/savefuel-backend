-- name: CreateRefreshToken :one
insert into refresh_tokens (token, user_id, expires_at)
values (
  $1,
  $2,
  $3
  )
  returning *;

-- name: GetRefreshToken :one
select * from refresh_tokens where token = $1;

-- name: RevokeRefreshToken :exec
update refresh_tokens set revoked_at = NOW() where token = $1;

-- PurgeRefreshTokens: Run to purge every token older than 30 days from database
