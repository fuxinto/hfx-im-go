CREATE TABLE 'auth' (
    'id' bigint,
    'uid' varchar(255) not null,
    'type' smallint not null ,
    'identifier' varchar(255) not null,
    'credential' varchar(255) not null,
    PRIMARY KEY ('id')
);