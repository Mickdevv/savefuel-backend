-- name: CreateUser :one
insert into users (id, created_at, updated_at, email, password)
values (
	gen_random_UUID(),
	NOW(),
	NOW(),
	$1,
	$2
	)
	returning *;
