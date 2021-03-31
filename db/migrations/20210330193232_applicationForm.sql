
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table applicationForm
(
    id int not null auto_increment,
    name varchar(255) not null,
    brithday datetime not null,
    studentId varchar(255) not null,
    college varchar(255) not null,
    profession varchar(255) not null,
    phoneNumber varchar(255) not null,
    email varchar(255) not null,
    qq varchar(255) not null,
    gender varchar(255) not null,
    selfIntroduction varchar(255) not null,
    blog varchar(255) null,
    github varchar(255) null,
    skills varchar(255) null,
    elseskills varchar(255) null,
    expections varchar(255) null,


    created_at datetime,
    updated_at datetime,
    deleted_at datetime,

    primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table applicationForm
