
-- +migrate Up
CREATE TABLE user (
    id          serial,
    name        varchar(256),
    is_admin    bool 
);

-- +migrate Down
DROP TABLE user;
