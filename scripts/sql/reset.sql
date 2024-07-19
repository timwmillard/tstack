drop schema if exists auth cascade;
drop schema if exists app cascade;
drop schema if exists stripe cascade;

drop table if exists schema_version;

drop table if exists river_job;
drop table if exists river_leader;
drop table if exists river_queue;
drop table if exists river_migration;
drop type if exists river_job_state;

drop table if exists sessions;

