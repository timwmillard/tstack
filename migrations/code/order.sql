
-- Order Update Trigger
drop trigger if exists order_changed_trigger on app.order;
drop function if exists app.order_changed;

create function app.order_changed() 
   returns trigger 
as $$
begin
    
    if new.status is distinct from old.status then
        case new.status
            when 'complete' then new.completed_on = now();
            when 'expired' then new.expired_on = now();
            else null;
        end case;
    end if;

    if tg_op = 'INSERT' then
        new.started_on = now();
    end if;

    return new;

end;
$$ language plpgsql;


create trigger order_changed_trigger
before insert or update on app.order
for each row
    execute function app.order_changed();

