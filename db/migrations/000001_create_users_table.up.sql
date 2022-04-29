create table if not exists users (
    wallet_address varchar(64) primary key,
    username       varchar(128) unique,
    created_at     timestamp not null,
    updated_at     timestamp
);
