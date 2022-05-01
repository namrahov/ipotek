-- +migrate Up
create table if not exists client_request
(
    id            bigserial    not null primary key,
    product_id    bigserial    not null,
    name          varchar(32)  not null,
    surname       varchar(32)  not null,
    pin_code      varchar(16),
    phone         varchar(32)  not null,
    date_of_birth varchar(16),
    created_at    timestamp    not null default now(),
    updated_at    timestamp    not null default current_timestamp
);
