-- +migrate Up
alter table product
add column email varchar(128);

update product
set email='ipoteka@pashabank.az';
