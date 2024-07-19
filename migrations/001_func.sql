
-- Place holder, actual function is located in code.
create or replace function uuid_generate_v7()
returns uuid
as $$
begin
    return '00000000-0000-0000-0000-000000000000'::uuid;
end
$$
language plpgsql
volatile;

-- Place holder, actual function is located in code.
create or replace function uuid_generate_v8()
returns uuid
as $$
begin
    return '00000000-0000-0000-0000-000000000000'::uuid;
end
$$
language plpgsql
volatile;

-- update_at trigger
create or replace function update_at_trigger ()
    returns trigger
    as $$
begin
    new.updated_at = now();
    return new;
end;
$$
language 'plpgsql'

