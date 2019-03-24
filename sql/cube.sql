-- Database: currency
-- DROP DATABASE currency;
create database currency
with owner = postgres
encoding = 'utf8'
tablespace = pg_default
lc_collate = 'c'
lc_ctype = 'c'
connection limit = -1;
grant all on database currency to postgres;

-- Table: Cube
-- drop table if exists cube cascade;
create table cube
(
  id        serial,
  currency  character(3)  not null,
  rate      numeric(10,5) not null,
  rate_time timestamp     not null,
  constraint cube_pkey primary key (id)
)
  with
(
  oids=false
);
alter table cube
  owner to postgres;

-- Create a unique index for cube columns currency, rate and rate_time to prevent duplicates.
create unique index cube_crr_index
on  cube
  (
    currency,
    rate,
    rate_time
  );

-- Modify the table to add unique index constraint.
alter table cube
  add constraint cube_crr_index
    unique using index cube_crr_index;

