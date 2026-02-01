-- +goose Up
-- +goose StatementBegin

CREATE TABLE users(
	id UUID primary key unique,
	created_at timestamp not null,
	updated_at timestamp not null,
	email text not null unique,
  email_verified bool not null default false,
	password TEXT not null default 'unset'
);

CREATE TABLE refresh_tokens(
  token TEXT primary key unique,
  revoked_at timestamp default null,
  created_at timestamp not null default NOW(),
  expires_at timestamp not null,
  user_id UUID not null,
  foreign key(user_id) references users(id) on delete cascade
);

CREATE TABLE document_categories(
  id uuid primary key unique,
  name text not null,
	created_at timestamp not null,
	updated_at timestamp not null,
  active bool not null default true
  
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
	priority int not null default 0,
  active bool not null default true,
  category_id uuid not null,
  hits int default 0,
  foreign key(category_id) references document_categories(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE document_categories;
DROP TABLE documents;
DROP TABLE users;
DROP TABLE refresh_tokens;

-- +goose StatementEnd
