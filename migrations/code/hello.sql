
create or replace function greet(name text)
returns text as
$$
begin

  return 'Hello ' || name;

end;
$$ language plpgsql;

