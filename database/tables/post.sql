create table if not exists post
(
    post_id      uuid primary key,
    author       uuid                     not null,
    title        varchar(255)             not null,
    slug         varchar(50)              not null,
    published    bool                     not null,
    created_at   timestamp with time zone not null,
    updated_at   timestamp with time zone,
    published_at timestamp with time zone,
    foreign key (author) references "user" (user_id)
);