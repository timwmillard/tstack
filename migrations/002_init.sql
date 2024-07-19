--  alter table schema_version add column code_version integer not null default 0;

create extension if not exists "uuid-ossp";

create schema if not exists auth;

create table auth.user (
    id uuid primary key default uuid_generate_v7(),
    username text not null unique check (username <> ''),
    password text not null check (password <> ''),
    email text unique,
    email_confirmed_at timestamptz,
    phone text unique,
    phone_confirmed_at timestamptz
);

create schema if not exists app;

create table app.admin (
    user_id uuid primary key default uuid_generate_v7() references auth.user(id),
    first_name text not null default '',
    last_name text not null default '',
    created_at timestamptz not null default now()
);

create table app.customer (
    id uuid primary key default uuid_generate_v7(),
    user_id uuid references auth.user(id),
    first_name text not null default '',
    last_name text not null default '',
    email text not null default '',
    created_at timestamptz not null default now()
);

create type app.order_status as enum (
    'pending',
    'processing',
    'complete',
    'expired'
);

create table app.order (
    id uuid primary key default uuid_generate_v7(),
    customer_id uuid references app.customer(id) on delete restrict,
    status app.order_status not null default 'pending',
    started_on timestamptz not null default now(),
    completed_on timestamptz,
    expired_on timestamptz
);


create table app.order_item (
    id uuid primary key default uuid_generate_v7(),
    order_id uuid not null references app.order(id) on delete cascade,
    description text not null,
    price numeric not null,
    qty integer not null default 1
);

-- Session pgxstore
-- https://github.com/alexedwards/scs/tree/master/pgxstore
create table sessions (
    token text primary key,
    data bytea not null,
    expiry timestamptz not null
);

create index sessions_expiry_idx on sessions (expiry);

