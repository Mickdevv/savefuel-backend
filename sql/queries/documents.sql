-- name: CreateDocument :one
insert into documents (id, created_at, path, filename, filetype, locale, title, description, priority, active, category_id, hits)
values (
	gen_random_UUID(),
	NOW(),
	$1, -- path
	$2, -- filename
	$3, -- filetype
	$4, -- locale
	$5, -- title
	$6, -- description
	$7, -- priority
  $8, -- active
  $9, -- category_id
  0 -- hits
	)
	returning *;

-- name: UpdateDocument :one
update documents set  locale = $2, title = $3, description =$4, priority = $5, active = $6, category_id = $7 where id = $1 returning *;

-- name: GetDocuments :many
select id, created_at, path, filename, filetype, category_id, locale, title, description, priority, active from documents;

-- name: GetDocument :one
select id, created_at, path, filename, filetype, category_id, locale, title, description, priority, active from documents where id = $1;

-- name: DeleteDocument :exec
delete from documents where id = $1;

