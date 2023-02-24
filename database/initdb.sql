CREATE USER gouser WITH PASSWORD 'yourniceownpwd';
GRANT CONNECT ON DATABASE golangdb TO gouser;
GRANT USAGE ON SCHEMA public TO gouser;
create table todo
(
    id serial not null,
    title text not null,
    is_done boolean default false not null,
    date_created timestamp default now() not null,
    id_creator int default 1 not null,
    date_last_modification timestamp,
    id_last_modifier int
);

comment on table todo is 'simple todo table to make some db tests';

create unique index todo_id_uindex on todo (id);

alter table todo  add constraint todo_pk primary key (id);

insert into todo (title, is_done) values ( 'learn golang language', false);
insert into todo (title, is_done) values ( 'install ubuntu gnu/linux on my computer', true);