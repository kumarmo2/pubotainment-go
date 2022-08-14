create keyspace servicediscovery
with replication = { 'class': 'SimpleStrategy', 'replication_factor': 2};

create keyspace songs
with replication = { 'class': 'SimpleStrategy', 'replication_factor': 1};

create table servicediscovery.serverinstances (
    id text,
    ips list<text>,
    registeredon timestamp,
    primary key(id)
);


create table servicediscovery.connectionservermap (
connectionid text,
serverid text,
lastpinged timestamp,
companyid bigint,
primary key((companyid), connectionid)
); 


create table songs.inventory_main(
    companyid bigint ,
    id bigint,
    name text,
    createon timestamp,
    modifiedon timestamp,
    primary key((companyid), id)
) with clustering order by (id desc);

create table songs.inventory_sorted_by_name(
    companyid bigint,
    name text,
    id bigint,
    primary key((companyid), name)
) with clustering order by(name asc);

