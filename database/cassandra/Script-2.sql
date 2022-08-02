create keyspace servicediscovery
with replication = { 'class': 'SimpleStrategy', 'replication_factor': 2}

describe servicediscovery;

use servicediscovery;

create table servicediscovery.serverinstances (
id text,
ips list<text>,
registeredon timestamp,
primary key(id)
);


create table servicediscovery.deviceconnectionmap (
deviceid text,
serverid text,
lastpinged timestamp,
companyid bigint,
primary key((companyid), deviceid, lastpinged)
) with clustering order by (deviceid asc, lastpinged desc); 



drop table servicediscovery.serverinstances;

insert into servicediscovery.serverinstances
(id, ips) values ('id1', ['sdfsdfdgdfg fdgsdfgdgsdfsdfsdfsdff'])

select * from servicediscovery.serverinstances;
truncate servicediscovery.serverinstances;