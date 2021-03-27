
-- +migrate Up
CREATE TABLE account (
    id          serial,
    name        varchar(256),
    is_admin    bool 
);

-- +migrate Down
DROP TABLE account;
