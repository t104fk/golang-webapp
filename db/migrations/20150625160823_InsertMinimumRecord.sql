
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO users VALUES (1, 'takasing', NOW(), NOW());
INSERT INTO articles VALUES (1, 'たいたいたいとる', 'ほんほんほんぶん国国', NOW(), NOW());

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM users WHERE id = 1;
DELETE FROM articles WHERE id = 1;

