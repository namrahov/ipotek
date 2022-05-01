-- +migrate Up
create table if not exists client_question_request
(
    id           bigserial    not null primary key,
    product_id   bigserial    not null,
    name         varchar(32)  not null,
    surname      varchar(32)  not null,
    phone        varchar(32)  not null,
    created_at   timestamp    not null default now(),
    updated_at   timestamp    not null default current_timestamp
);
