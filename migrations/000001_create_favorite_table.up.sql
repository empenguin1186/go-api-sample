create table if not exists favorite (
    tweet_id char(12),
    registered_at timestamp with time zone default CURRENT_TIMESTAMP not null,
    primary key (tweet_id)
);

insert into favorite values (
'testtest0001',
'2022-02-25 00:00:00+0900'
),(
'testtest0002',
'2022-02-26 00:00:00+0900'
),
(
'testtest0003',
'2022-02-27 00:00:00+0900'
);