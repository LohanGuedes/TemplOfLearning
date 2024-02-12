-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  username varchar UNIQUE NOT NULL,
  email varchar UNIQUE NOT NULL,
  password bytea NOT NULL,
  role varchar NOT NULL,
  verified bool DEFAULT false,
  created_at timestamp DEFAULT (NOW()),
  updated_at timestamp DEFAULT (NOW())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
