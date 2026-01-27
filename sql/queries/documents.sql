-- name: CreateDocument :one
insert into documents (id, created_at, path, filename, filetype, locale, title, description, priority)
values (
	gen_random_UUID(),
	NOW(),
	$1, -- path
	$2, -- filename
	$3, -- filetype
	$4, -- locale
	$5, -- title
	$6, -- description
	$7 -- priority
	)
	returning *;

