create schema if not exists authentication;

create table if not exists authentication.companies
(
	id bigint not null,
	name text not null,
	allowedadmindevices int not null default(1),
	alloweduserdevices int not null default(10),
	primary key(id)
);

insert into authentication.companies (id, name)
values (1, 'company-1') on conflict(id) do nothing;

alter table authentication.companies
add column if not exists adminhashedpass text default null,
add column if not exists userhashedpass text default null;


