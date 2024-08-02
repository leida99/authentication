# authentication
Go+mysql完成一个认证系统

## 初始化
cd myapp
go mod init myapp
go mod tidy
go run .


# mysql命令
## 登录mysql
mysql -u root -p
## 查看所有的数据库
SHOW DATABASES;
## 查看所有表
use mysql;
show tables;
## 创建数据库，指定数据库的字符集和排序规则
CREATE DATABASE personal_info_web
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;
## 清空表
TRUNCATE TABLE table_name;
## 删除表
DROP TABLE table_name;
## 唯一约束
id_card VARCHAR(20) NOT NULL UNIQUE,