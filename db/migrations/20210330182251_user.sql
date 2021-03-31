
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table user
(
    id int not null  auto_increment,
    roleid int not null,
    email varchar(255) not null, 

    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table user;