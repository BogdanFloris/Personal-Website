create type post_part_type as enum ('text', 'media', 'code');

create table if not exists post_part
(
    post_id uuid           not null,
    number  int            not null,
    type    post_part_type not null,
    data    bytea          not null,
    primary key (post_id, number),
    foreign key (post_id) references post (post_id)
);