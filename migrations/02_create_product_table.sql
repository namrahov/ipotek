-- +migrate Up
create table if not exists product
(
    id           bigserial     not null primary key,
    name         varchar(32)   not null,
    name_az      varchar(64)   not null,
    status       varchar(16)   not null,
    description  varchar(1024),
    created_at   timestamp     not null default now(),
    updated_at   timestamp     not null default current_timestamp
);

insert into product(name, name_az, status)
values ('ExtractiveDomesticMortgage', 'Çıxarışlı daxili ipoteka', 'ACTIVE'),
       ('UnsecuredDomesticMortgage', 'Çıxarışsız daxili ipoteka', 'ACTIVE'),
       ('MidaMortgage', 'Mida ipotekası', 'ACTIVE'),
       ('StateMortgage', 'Dövlət ipotekası', 'ACTIVE'),
       ('PreferentialStateMortgage', 'Güzəştli dövlət ipotekası', 'ACTIVE');
