-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table role
(
    id int not null,
    alias varchar(255) not null,
    bame varchar(255) not null,

    primary key (id)
);