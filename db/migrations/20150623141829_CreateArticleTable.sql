
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE articles(
  id serial PRIMARY KEY,
  title varchar(255) NOT NULL,
  content text,
  created_at TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00',
  updated_at TIMESTAMP NOT NULL DEFAULT '1970-01-01 00:00:00'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE articles;

