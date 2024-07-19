
-- Stripe schema
create schema if not exists stripe;

create table if not exists stripe.checkout_session (
    id text primary key,
    object jsonb not null default '{}',
    order_id uuid references app.order(id),
    created_at timestamptz not null default now()
);

