create keyspace servicediscovery
with replication = { 'class': 'SimpleStrategy', 'replication_factor': 2};

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
