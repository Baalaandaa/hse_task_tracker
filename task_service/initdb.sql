create table "user"
(
    id           uuid primary key default gen_random_uuid(),
    login        varchar not null,
    password     varchar not null,
    "name"       varchar,
    surname      varchar,
    birthday     timestamp,
    email        varchar,
    phone_number varchar
)