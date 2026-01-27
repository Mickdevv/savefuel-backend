-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
	id UUID primary key unique,
	created_at timestamp not null,
	updated_at timestamp not null,
	email text not null unique,
	password TEXT not null default 'unset'
);

CREATE TABLE documents(
	id UUID primary key unique,
	created_at timestamp not null default NOW(),
	path TEXT NOT NULL,
	filename text not null,
	filetype text not null,
	locale text not null,
	title text not null,
	description text not null,
	priority smallint 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE documents;
DROP TABLE users;

-- +goose StatementEnd
