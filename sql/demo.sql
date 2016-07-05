drop table if exists `user_info`;
create table `user_info`
(
    `id` bigint(20) unsigned not null auto_increment comment '主键',
    `nick_name` varchar(32) not null comment '昵称',
    `mobile` varchar(11) not null comment '手机号',
    `password` varchar(128) not null comment '密码',
    `gmt_create` timestamp not null default current_timestamp comment '写入时间',
    primary key(`id`),
    index `i_mobile` (`mobile`),
    unique `u_mobile` (`mobile`)
)engine=InnoDB default charset=utf8mb4 comment='用户表';
