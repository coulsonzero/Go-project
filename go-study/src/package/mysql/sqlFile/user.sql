create table 'Users'(
    'id' int(10) NOT NULL AUTO_INCREMENT,
    'username' varchar(30) NOT NULL DEFAULT '' COMMENT '用户名',
    'password' varchar(20) NOT NULL DEFAULT '' COMMENT '用户密码',
    'class' char(1) NOT NULL COMMENT '级别',
    'hire_date' date NOT NULL COMMENT '上岗时间',
    PRIMARY KEY('id')
) ENGINE=InnoDB CHARAULT=utf8 COMMENT='员工信息';