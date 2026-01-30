-- name: CreateUser :one
insert into users (id, created_at, updated_at, email, password, email_verified)
values (
	gen_random_UUID(),
	NOW(),
	NOW(),
	$1,
	$2,
  false
	)
	returning id, created_at, updated_at, email, email_verified;


-- name: GetUserForAuth :one 
select id, created_at, updated_at, email, email_verified, password from users where email = $1;

-- name: GetUserByEmail :one 
select id, created_at, updated_at, email, email_verified from users where email = $1;

-- name: GetUserById :one 
select id, created_at, updated_at, email, email_verified from users where id = $1;

-- name: GetUsers :many
select id, created_at, updated_at, email, email_verified from users; 

-- name: UpdateUser :one
update users set email =$2, email_verified = $3, password = $4, updated_at = NOW() where id = $1 returning id, created_at, updated_at, email, email_verified;

-- name: Deleteuser :exec
delete from users where id = $1;
