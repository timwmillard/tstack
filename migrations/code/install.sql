
-- update schema_version set code_version = 1;

create schema if not exists app;

{{ template "uuid_v7_v8.sql" . }}
{{ template "order.sql" . }}
{{ template "hello.sql" . }}

