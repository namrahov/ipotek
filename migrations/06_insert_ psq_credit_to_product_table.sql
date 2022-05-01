-- +migrate Up

insert into product(name, name_az, status, email)
values ('PSQ', 'Paşa Bank Kredit Əməkdaşları', 'ACTIVE', 'PASHA_GROUP_CREDIT_MEMBERS');
