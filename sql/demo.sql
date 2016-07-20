drop table if exists `user_info`;
create table `user_info`
(
    `id` bigint(20) unsigned not null auto_increment comment '主键',
    `nick_name` varchar(32) not null comment '昵称',
    `mobile` varchar(11) not null comment '手机号',
    `password` varchar(128) not null comment '密码',
    `taobao_id` varchar(32) not null  default '' comment '淘宝账号',
    `real_name` varchar(20) not null default '' comment '淘宝绑定的姓名',
    `ali_id` varchar(32) not null default '' comment '支付宝账号',
    `gmt_create` timestamp not null default current_timestamp comment '写入时间',
    primary key(`id`),
    index `i_mobile` (`mobile`),
    unique `u_mobile` (`mobile`)
)engine=InnoDB default charset=utf8mb4 comment='用户表';


drop table if exists `seller_info`;
create table `seller_info`
(
    `id` bigint(20) unsigned not null auto_increment comment '主键',
    `nick_name` varchar(32) not null comment '昵称',
    `mobile` varchar(11) not null comment '手机号',
    `password` varchar(128) not null comment '密码',
    `balance` bigint(20) not null comment '商家的余额',
    `gmt_create` timestamp not null default current_timestamp comment '写入时间',
    primary key(`id`),
    index `i_mobile` (`mobile`),
    unique `u_mobile` (`mobile`)
)engine=InnoDB default charset=utf8mb4 comment='商家信息表';


-- 后台人员表
drop table if exists `manager_info`;
create table `manager_info`
(
    `id` bigint(20) unsigned not null auto_increment comment '主键',
    `nick_name` varchar(32) not null default '' comment '后台人员昵称',
    `mobile` varchar(11) not null comment '手机号',
    `password` varchar(128) not null comment '密码',
    `gmt_create` timestamp not null default current_timestamp comment '写入时间',
    primary key(`id`),
    index `i_mobile` (`mobile`),
    unique `u_mobile` (`mobile`)
)engine=InnoDB default charset=utf8mb4 comment='后台人员表';




drop table if exists `order_info`;
create table `order_info`
(
    `id` bigint(20) unsigned not null auto_increment comment '主键',
    `user_id` bigint(20) unsigned not null comment '用户id',
    `goods_id` varchar(20) not null comment '商品id',
    `order_num` varchar(20) not null comment '订单号',
    `pay_money` bigint(20) not null comment '订单佣金',
    `buy_status` tinyint not null comment '购买状态',
    `receipt_status` tinyint not null comment '收货状态',
    `pay_status` tinyint not null comment '佣金支付状态',
    `gmt_create` timestamp not null default current_timestamp comment '写入时间',
    primary key(`id`),
    index `i_order_num` (`order_num`),
    unique `u_order_num` (`order_num`)
)engine=InnoDB default charset=utf8mb4 comment='订单信息表';
