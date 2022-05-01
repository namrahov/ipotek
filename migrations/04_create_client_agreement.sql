-- +migrate Up
create table if not exists client_agreement
(
    id              bigserial      not null primary key,
    client_id       bigserial      not null,
    agreement_text  varchar(800),
    is_signed       boolean,
    sign_date       varchar(16),
    created_at      timestamp      not null default now(),
    updated_at      timestamp      not null default current_timestamp
);
