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
	filetype text not null
);
CREATE TABLE document_metadata(
	id UUID primary key unique,
	document_id UUID not null references documents(id) on delete cascade,
	locale text not null,
	title text not null,
	description text,
	UNIQUE(document_id, locale)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE document_translations;
DROP TABLE documents;
DROP TABLE users;

-- +goose StatementEnd
