
-- Create tim login user
insert into auth.user (
    id,
    username,
    password,
    email,
    email_confirmed_at,
    phone,
    phone_confirmed_at)
values (
    '018ba511-12c1-75df-a64e-b14b97510251',
    'tim',
    '$2a$14$AZq5IlSoX6K5nC4Jb.z/v.DfgKxkRTMjRZRtIW9uTcNJvaQGe.sx.',
    null,
    null,
    null,
    null);

insert into app.admin (
    user_id,
    first_name,
    last_name
) values (
    '018ba511-12c1-75df-a64e-b14b97510251',
    'Tim',
    'Millard'
);

