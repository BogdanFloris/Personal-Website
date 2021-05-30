begin;

create type user_role as enum ('admin', 'normal');

create table if not exists "user"
(
    user_id       uuid primary key         not null,
    username      varchar(50)              not null unique,
    password_hash varchar(60)              not null,
    role          user_role                not null,
    created_on    timestamp with time zone not null,
    last_login    timestamp with time zone
);

create unique index if not exists user_idx on "user"(username);

commit;