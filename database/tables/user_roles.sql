create table if not exists user_roles
(
    user_id      uuid not null,
    role_id      int  not null,
    granted_date timestamp with time zone,
    primary key (user_id, role_id),
    foreign key (user_id) references "user" (user_id),
    foreign key (role_id) references role (role_id)
);