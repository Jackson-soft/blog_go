create table `user_info`
(
    `id` bigint(20) unsigned not null auto_increment comment '主键',
    `mobile` varchar(11) not null comment '手机号',
    `password` varchar(128) not null comment '密码',
    primary key(`id`)
)engine=InnoDB default charset=utf8 comment='用户表';