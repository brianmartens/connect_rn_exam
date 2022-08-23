create schema if not exists test;
use test;

create table user_info
(
    id         bigint unsigned auto_increment primary key,
    first_name varchar(255) not null,
    last_name  varchar(255) null,
    city       varchar(255) not null,
    zip_code   int          null
);

create table password_history
(
    password     varchar(255)                         null,
    date_changed timestamp  default CURRENT_TIMESTAMP null,
    active       tinyint(1) default false             null,
    user_id      bigint unsigned                      null,
    constraint user_id_constraint
        foreign key (user_id) references user_info (id)
            on delete cascade
);

select *
from password_history
where active = true;

start transaction;

update password_history set active = false where active = true;

insert into password_history (user_id, password, active)
values (1, 'password', true);

commit;