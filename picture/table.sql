create table ztcbase.info
(
    id   bigint auto_increment
        primary key,
    name varchar(255) null,
    note varchar(255) null,
    path varchar(255) null comment '文件保存路径',
    unix bigint       null comment 'unix时间戳'
)
    comment '相册';