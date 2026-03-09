-- ----------------------------
-- 基础数据字典
-- ----------------------------
INSERT INTO `sys_dict_type` (`dict_code`, `dict_name`, `sort`, `remark`) VALUES
('USER_STATUS', '用户状态', 1, '系统用户的启用/禁用状态'),
('ORDER_TYPE', '订单类型', 2, '订单的业务类型分类');

INSERT INTO `sys_dict_item` (`dict_code`, `item_value`, `item_label`, `sort`, `remark`) VALUES
('USER_STATUS', '0', '禁用', 1, '用户账号禁用，无法登录'),
('USER_STATUS', '1', '启用', 2, '用户账号正常，可登录'),
('ORDER_TYPE', '01', '普通订单', 1, '常规商品购买订单'),
('ORDER_TYPE', '02', '秒杀订单', 2, '限时秒杀活动订单'),
('ORDER_TYPE', '03', '拼团订单', 3, '多人拼团订单');

-- ----------------------------
-- 智管系统核心业务表DDL脚本
-- 适用版本：MySQL 8.0+
-- 字符集：utf8mb4（支持中文/Emoji）
-- 存储引擎：InnoDB（支持事务/外键）
-- 执行顺序：property → tenant → contract → bill → operation_log
-- ----------------------------

-- 1. 房源信息表
CREATE TABLE IF NOT EXISTS `hrs_property` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '房源ID（主键）',
  `property_code` VARCHAR(30) NOT NULL COMMENT '房源编号（唯一标识）',
  `property_title` VARCHAR(100) NOT NULL COMMENT '房源标题',
  `owner_name` VARCHAR(50) NOT NULL COMMENT '产权人姓名',
  `property_cert_no` VARCHAR(50) NOT NULL COMMENT '产权证号',
  `address` VARCHAR(255) NOT NULL COMMENT '房屋地址（精确至门牌号）',
  `area` DECIMAL(10,2) NOT NULL COMMENT '建筑面积（单位：㎡）',
  `room_type` VARCHAR(20) NOT NULL COMMENT '户型（如"3室2厅1卫"）',
  `orientation` VARCHAR(10) NOT NULL COMMENT '朝向',
  `decoration` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '装修程度（0：毛坯，1：简装，2：精装）',
  `room_count` INT(2) NOT NULL COMMENT '房间数量',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '房源状态（0：待租，1：已租，2：维护中）',
  `price` DECIMAL(10,2) NOT NULL COMMENT '租金价格',
  `description` TEXT DEFAULT NULL COMMENT '房源描述',
  `creator_id` INT(11) NOT NULL COMMENT '创建人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_property_code` (`property_code`),
  UNIQUE KEY `uk_property_cert_no` (`property_cert_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='房源信息表';

-- 2. 房源状态变更表
CREATE TABLE IF NOT EXISTS `hrs_property_status_log` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '日志ID（主键）',
  `property_id` INT(11) NOT NULL COMMENT '房源ID',
  `old_status` TINYINT(1) NOT NULL COMMENT '原状态',
  `new_status` TINYINT(1) NOT NULL COMMENT '新状态',
  `change_reason` VARCHAR(255) NOT NULL COMMENT '变更原因',
  `operator_id` INT(11) NOT NULL COMMENT '操作人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_property_id` (`property_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='房源状态变更表';

-- 3. 房源图片表
CREATE TABLE IF NOT EXISTS `hrs_property_image` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '图片ID（主键）',
  `property_id` INT(11) NOT NULL COMMENT '房源ID',
  `image_url` VARCHAR(500) NOT NULL COMMENT '图片URL',
  `image_type` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '图片类型（0：室内图，1：室外图，2：户型图）',
  `sort_order` INT(2) NOT NULL DEFAULT 0 COMMENT '排序顺序',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_property_id` (`property_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='房源图片表';

-- 4. 房源标签表
CREATE TABLE IF NOT EXISTS `hrs_property_tag` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '标签ID（主键）',
  `property_id` INT(11) NOT NULL COMMENT '房源ID',
  `tag_name` VARCHAR(50) NOT NULL COMMENT '标签名称',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_property_id` (`property_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='房源标签表';

-- 5. 房源带看记录表
CREATE TABLE IF NOT EXISTS `hrs_property_viewing` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '带看记录ID（主键）',
  `property_id` INT(11) NOT NULL COMMENT '房源ID',
  `tenant_id` INT(11) DEFAULT NULL COMMENT '租户ID',
  `viewer_name` VARCHAR(50) NOT NULL COMMENT '看房人姓名',
  `viewer_phone` VARCHAR(20) NOT NULL COMMENT '看房人电话',
  `view_time` DATETIME NOT NULL COMMENT '带看时间',
  `feedback` TEXT DEFAULT NULL COMMENT '租户反馈',
  `next_plan` VARCHAR(255) DEFAULT NULL COMMENT '后续计划',
  `agent_id` INT(11) NOT NULL COMMENT '经纪人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_property_id` (`property_id`),
  KEY `idx_tenant_id` (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='房源带看记录表';

-- 2. 租户信息表
CREATE TABLE IF NOT EXISTS `hrs_tenant` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '租户ID（主键）',
  `tenant_code` VARCHAR(30) NOT NULL COMMENT '租户编号（唯一标识）',
  `name` VARCHAR(50) NOT NULL COMMENT '租户姓名',
  `gender` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '性别（0：男，1：女）',
  `age` INT(3) NOT NULL COMMENT '年龄',
  `id_card` VARCHAR(18) NOT NULL COMMENT '身份证号（唯一）',
  `phone` VARCHAR(20) NOT NULL COMMENT '联系电话',
  `phone2` VARCHAR(20) DEFAULT NULL COMMENT '备用联系电话',
  `email` VARCHAR(100) DEFAULT NULL COMMENT '电子邮箱',
  `occupation` VARCHAR(100) DEFAULT NULL COMMENT '职业',
  `family_members` INT(2) NOT NULL DEFAULT 1 COMMENT '家庭人口',
  `credit_score` INT(3) DEFAULT NULL COMMENT '信用评分',
  `level` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '客户等级（0：普通，1：VIP）',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '租户状态（0：禁用，1：正常）',
  `creator_id` INT(11) NOT NULL COMMENT '创建人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_tenant_code` (`tenant_code`),
  UNIQUE KEY `uk_tenant_id_card` (`id_card`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='租户信息表';

-- 3. 租户资质文件表
CREATE TABLE IF NOT EXISTS `hrs_tenant_qualification` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '资质文件ID（主键）',
  `tenant_id` INT(11) NOT NULL COMMENT '租户ID',
  `file_type` TINYINT(1) NOT NULL COMMENT '文件类型（0：身份证正面，1：身份证反面，2：工作证明，3：征信报告，4：营业执照，5：法人身份证明）',
  `file_url` VARCHAR(500) NOT NULL COMMENT '文件URL',
  `verify_status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '核验状态（0：待核验，1：核验通过，2：核验失败）',
  `verify_result` VARCHAR(255) DEFAULT NULL COMMENT '核验结果',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='租户资质文件表';

-- 4. 租户跟进表
CREATE TABLE IF NOT EXISTS `hrs_tenant_followup` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '跟进记录ID（主键）',
  `tenant_id` INT(11) NOT NULL COMMENT '租户ID',
  `stage` TINYINT(1) NOT NULL COMMENT '跟进阶段（0：需求确认，1：带看安排，2：合同洽谈，3：入住维护，4：退租结算）',
  `content` TEXT NOT NULL COMMENT '跟进内容',
  `next_action` VARCHAR(255) DEFAULT NULL COMMENT '下一步行动',
  `next_time` DATETIME DEFAULT NULL COMMENT '下次跟进时间',
  `agent_id` INT(11) NOT NULL COMMENT '经纪人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='租户跟进表';

-- 5. 租户沟通记录表
CREATE TABLE IF NOT EXISTS `hrs_tenant_communication` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '沟通记录ID（主键）',
  `tenant_id` INT(11) NOT NULL COMMENT '租户ID',
  `communication_type` TINYINT(1) NOT NULL COMMENT '沟通类型（0：电话，1：微信，2：短信，3：面谈）',
  `communication_time` DATETIME NOT NULL COMMENT '沟通时间',
  `content` TEXT NOT NULL COMMENT '沟通内容',
  `duration` INT(5) DEFAULT NULL COMMENT '沟通时长（秒）',
  `recording_url` VARCHAR(500) DEFAULT NULL COMMENT '录音URL',
  `agent_id` INT(11) NOT NULL COMMENT '经纪人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_tenant_id` (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='租户沟通记录表';

-- 6. 合同模板表
CREATE TABLE IF NOT EXISTS `hrs_contract_template` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '模板ID（主键）',
  `template_name` VARCHAR(100) NOT NULL COMMENT '模板名称',
  `template_content` TEXT NOT NULL COMMENT '模板内容',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态（0：禁用，1：启用）',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合同模板表';

-- 7. 合同表
CREATE TABLE IF NOT EXISTS `hrs_contract` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '合同ID（主键）',
  `contract_code` VARCHAR(30) NOT NULL COMMENT '合同编号（唯一标识）',
  `property_id` INT(11) NOT NULL COMMENT '房源ID',
  `tenant_id` INT(11) NOT NULL COMMENT '租户ID',
  `template_id` INT(11) NOT NULL COMMENT '模板ID',
  `contract_name` VARCHAR(100) NOT NULL COMMENT '合同名称',
  `start_date` DATETIME NOT NULL COMMENT '开始日期',
  `end_date` DATETIME NOT NULL COMMENT '结束日期',
  `rent_amount` DECIMAL(10,2) NOT NULL COMMENT '租金金额',
  `deposit_amount` DECIMAL(10,2) NOT NULL COMMENT '押金金额',
  `payment_method` TINYINT(1) NOT NULL COMMENT '付款方式（0：月付，1：季付，2：半年付，3：年付）',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '合同状态（0：待签署，1：已签署，2：履行中，3：已到期，4：已终止）',
  `contract_file_url` VARCHAR(500) DEFAULT NULL COMMENT '合同文件URL',
  `creator_id` INT(11) NOT NULL COMMENT '创建人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_contract_code` (`contract_code`),
  KEY `idx_property_id` (`property_id`),
  KEY `idx_tenant_id` (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合同表';

-- 8. 合同条款表
CREATE TABLE IF NOT EXISTS `hrs_contract_clause` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '条款ID（主键）',
  `contract_id` INT(11) NOT NULL COMMENT '合同ID',
  `clause_type` VARCHAR(50) NOT NULL COMMENT '条款类型',
  `clause_content` TEXT NOT NULL COMMENT '条款内容',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_contract_id` (`contract_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合同条款表';

-- 9. 合同签章表
CREATE TABLE IF NOT EXISTS `hrs_contract_signature` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '签章ID（主键）',
  `contract_id` INT(11) NOT NULL COMMENT '合同ID',
  `signer_type` TINYINT(1) NOT NULL COMMENT '签署人类型（0：房东，1：租户，2：中介）',
  `signer_id` INT(11) NOT NULL COMMENT '签署人ID',
  `signer_name` VARCHAR(50) NOT NULL COMMENT '签署人姓名',
  `signature_url` VARCHAR(500) NOT NULL COMMENT '签章URL',
  `sign_time` DATETIME NOT NULL COMMENT '签署时间',
  PRIMARY KEY (`id`),
  KEY `idx_contract_id` (`contract_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合同签章表';

-- 10. 合同状态变更表
CREATE TABLE IF NOT EXISTS `hrs_contract_status_log` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '日志ID（主键）',
  `contract_id` INT(11) NOT NULL COMMENT '合同ID',
  `old_status` TINYINT(1) NOT NULL COMMENT '原状态',
  `new_status` TINYINT(1) NOT NULL COMMENT '新状态',
  `change_reason` VARCHAR(255) NOT NULL COMMENT '变更原因',
  `operator_id` INT(11) NOT NULL COMMENT '操作人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_contract_id` (`contract_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合同状态变更表';

-- 11. 合同风险预警表
CREATE TABLE IF NOT EXISTS `hrs_contract_risk` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '预警ID（主键）',
  `contract_id` INT(11) NOT NULL COMMENT '合同ID',
  `risk_type` TINYINT(1) NOT NULL COMMENT '风险类型（0：到期预警，1：欠租预警，2：违约预警）',
  `risk_level` TINYINT(1) NOT NULL COMMENT '风险等级（0：低，1：中，2：高）',
  `risk_message` VARCHAR(255) NOT NULL COMMENT '风险信息',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '处理状态（0：未处理，1：已处理）',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_contract_id` (`contract_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合同风险预警表';

-- 12. 账单表
CREATE TABLE IF NOT EXISTS `hrs_bill` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '账单ID（主键）',
  `bill_code` VARCHAR(30) NOT NULL COMMENT '账单编号（唯一标识）',
  `contract_id` INT(11) NOT NULL COMMENT '合同ID',
  `tenant_id` INT(11) NOT NULL COMMENT '租户ID',
  `property_id` INT(11) NOT NULL COMMENT '房源ID',
  `bill_type` TINYINT(1) NOT NULL COMMENT '账单类型（0：租金，1：押金，2：水电费，3：物业费，4：其他）',
  `amount` DECIMAL(10,2) NOT NULL COMMENT '账单金额',
  `due_date` DATETIME NOT NULL COMMENT '到期日期',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '账单状态（0：待支付，1：已支付，2：逾期，3：部分支付）',
  `paid_amount` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '已支付金额',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '账单描述',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_bill_code` (`bill_code`),
  KEY `idx_contract_id` (`contract_id`),
  KEY `idx_tenant_id` (`tenant_id`),
  KEY `idx_property_id` (`property_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账单表';

-- 13. 支付记录表
CREATE TABLE IF NOT EXISTS `hrs_payment` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '支付记录ID（主键）',
  `payment_code` VARCHAR(30) NOT NULL COMMENT '支付编号（唯一标识）',
  `bill_id` INT(11) NOT NULL COMMENT '账单ID',
  `tenant_id` INT(11) NOT NULL COMMENT '租户ID',
  `amount` DECIMAL(10,2) NOT NULL COMMENT '支付金额',
  `payment_method` TINYINT(1) NOT NULL COMMENT '支付方式（0：微信，1：支付宝，2：银行卡，3：现金）',
  `payment_time` DATETIME NOT NULL COMMENT '支付时间',
  `transaction_id` VARCHAR(100) NOT NULL COMMENT '交易单号',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '支付状态（0：失败，1：成功）',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_payment_code` (`payment_code`),
  KEY `idx_bill_id` (`bill_id`),
  KEY `idx_tenant_id` (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付记录表';

-- 14. 财务报表表
CREATE TABLE IF NOT EXISTS `hrs_financial_report` (
  `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '报表ID（主键）',
  `report_name` VARCHAR(100) NOT NULL COMMENT '报表名称',
  `report_type` TINYINT(1) NOT NULL COMMENT '报表类型（0：日报，1：周报，2：月报，3：季报，4：年报）',
  `start_date` DATETIME NOT NULL COMMENT '开始日期',
  `end_date` DATETIME NOT NULL COMMENT '结束日期',
  `total_income` DECIMAL(12,2) NOT NULL COMMENT '总收入',
  `total_expense` DECIMAL(12,2) NOT NULL COMMENT '总支出',
  `net_income` DECIMAL(12,2) NOT NULL COMMENT '净收入',
  `report_url` VARCHAR(500) DEFAULT NULL COMMENT '报表文件URL',
  `creator_id` INT(11) NOT NULL COMMENT '创建人ID',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='财务报表表';