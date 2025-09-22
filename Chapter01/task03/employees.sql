CREATE TABLE employees (
                           id INT AUTO_INCREMENT PRIMARY KEY,  -- 主键，自增，唯一标识员工
                           name VARCHAR(100) NOT NULL,         -- 员工姓名，非空，最长100字符
                           department VARCHAR(50),             -- 部门名称，最长50字符，允许为空
                           salary DECIMAL(10, 2) NOT NULL      -- 薪资，保留2位小数，非空（精度：最大99999999.99）
);

INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (1, '李芳', '研发部', 6818.88);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (2, '周伟', '销售部', 12940.12);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (3, '吴阳', '研发部', 18198.03);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (4, '刘强', '财务部', 13270.09);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (5, '杨阳', '销售部', 18386.61);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (6, '赵芳', '销售部', 5639.10);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (7, '张浩', '研发部', 4106.98);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (8, '杨静', '市场部', 7404.28);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (9, '杨敏', '销售部', 9324.18);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (10, '周敏', '人事部', 16936.30);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (11, '张敏', '人事部', 17021.41);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (12, '陈晓', '人事部', 9256.48);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (13, '王阳', '研发部', 9885.81);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (14, '吴芳', '市场部', 19298.07);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (15, '刘娜', '市场部', 15411.99);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (16, '周浩', '销售部', 17919.77);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (17, '李晓', '销售部', 4690.01);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (18, '周晓', '研发部', 9278.65);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (19, '李芳', '研发部', 12487.39);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (20, '王娜', '销售部', 10981.39);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (21, '陈静', '财务部', 18603.64);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (22, '刘敏', '人事部', 15452.91);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (23, '赵伟', '研发部', 9631.66);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (24, '李芳', '研发部', 9122.52);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (25, '陈娜', '技术部', 10127.24);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (26, '郑伟', '人事部', 18877.40);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (27, '王浩', '财务部', 15148.60);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (28, '吴敏', '人事部', 5032.22);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (29, '周阳', '销售部', 14943.32);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (30, '陈帆', '人事部', 6534.86);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (31, '杨浩', '研发部', 19000.45);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (32, '刘浩', '研发部', 15166.36);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (33, '杨娜', '财务部', 18470.60);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (34, '王晓', '市场部', 9574.18);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (35, '周强', '销售部', 7999.12);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (36, '周强', '研发部', 16635.67);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (37, '郑晓', '市场部', 13102.12);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (38, '吴晓', '技术部', 13318.27);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (39, '陈浩', '销售部', 18344.42);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (40, '刘芳', '财务部', 17522.37);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (41, '王敏', '销售部', 14773.28);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (42, '赵芳', '财务部', 18802.82);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (43, '赵伟', '人事部', 16100.68);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (44, '李静', '研发部', 16487.84);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (45, '吴晓', '销售部', 3265.67);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (46, '赵帆', '研发部', 6837.12);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (47, '郑敏', '销售部', 5692.30);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (48, '张强', '市场部', 16487.97);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (49, '刘敏', '财务部', 14982.80);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (50, '吴伟', '人事部', 18306.35);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (51, '张强', '研发部', 16388.59);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (52, '杨强', '财务部', 17031.54);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (53, '王帆', '技术部', 3602.42);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (54, '郑晓', '人事部', 15669.69);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (55, '王娜', '销售部', 14704.05);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (56, '李晓', '技术部', 10884.62);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (57, '周芳', '人事部', 16904.93);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (58, '李娜', '市场部', 14608.72);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (59, '陈帆', '研发部', 11025.34);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (60, '王敏', '技术部', 10264.47);
INSERT INTO `test`.`employees` (`id`, `name`, `department`, `salary`) VALUES (61, '张浩', '市场部', 12562.09);
