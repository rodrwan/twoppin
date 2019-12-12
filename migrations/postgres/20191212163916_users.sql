-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id          uuid primary key default gen_random_uuid(),

  username        varchar(30) not null,
  email       varchar(100) not null,

  created_at  timestamptz default now(),
  updated_at  timestamptz default now(),
  deleted_at  timestamptz
);

create trigger update_users_updated_at
before update on users for each row execute procedure update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
