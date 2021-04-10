create table if not exists role
(
    role_id   serial primary key,
    role_name varchar(50) unique
)