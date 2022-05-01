-- +migrate Up

ALTER TABLE client_request ALTER COLUMN name DROP NOT NULL;
ALTER TABLE client_request ALTER COLUMN surname DROP NOT NULL;
