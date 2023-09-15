-- +goose Up
set timezone = 'UTC-3';

create table if not exists service_user
(
    id serial primary key,
    name    varchar not null,
    surname    varchar not null,
    patronymic varchar not null
);

create table if not exists user_document (
    id serial primary key,
    doc_type varchar not null,
    number varchar unique,

    user_id int not null,
    foreign key (user_id) references service_user(id)
);

create table if not exists tickets (
    id serial primary key,
    start_point varchar not null,
    end_point varchar not null,

    start_time timestamp not null,
    end_time timestamp not null,

    buy_time timestamp DEFAULT now(),

    company varchar not null,

    user_id int not null,
    foreign key (user_id) references service_user(id)
);

-- +goose Down
drop table if exists user_document;
drop table if exists tickets;
drop table if exists service_user;