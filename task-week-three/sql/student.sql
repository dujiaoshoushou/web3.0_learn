# 题目1：基本CRUD操作
# 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
# 要求 ：
# 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
# 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
# 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
# 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
# 题目2：事务语句
# 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
# 要求 ：
# 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。


create database gorm_example;

create table students(
    id bigint primary key auto_increment,
    name varchar(255),
    age int8,
    grade varchar(255)
);



insert into students(name,age,grade) values ('张三',23,'三年级'),('李四',12,'一年级'),('王五',13,'二年级');

select * from students where age > 18;

update students set grade='四年级' where name = '张三';

delete from students where age <15;

create table accounts(
id bigint primary key auto_increment,
    name varchar(255),
    balance float8
);

create table transactions(
    id bigint primary key auto_increment,
    name varchar(2555),
    from_account_id bigint,
    to_account_id bigint,
    amount float8,
    foreign key (from_account_id) references accounts(id),
    foreign key (to_account_id) references accounts(id)
);

INSERT INTO accounts (name, balance) VALUES
    ('A', 500.00),
    ('B', 300.00);


START TRANSACTION;  -- 开始事务

-- 获取账户A余额并锁定行（避免并发修改）
SELECT balance INTO @current_balance FROM accounts WHERE name = 'A' FOR UPDATE;

-- 检查余额是否充足
IF @current_balance >= 100 THEN
    -- 账户A扣款
    UPDATE accounts SET balance = balance - 100 WHERE name = 'A';
    -- 账户B加款
    UPDATE accounts SET balance = balance + 100 WHERE name = 'B';
    -- 记录交易（获取实际账户ID）
    INSERT INTO transactions (from_account_id, to_account_id, amount)
    SELECT
        (SELECT id FROM accounts WHERE name = 'A') AS from_id,
        (SELECT id FROM accounts WHERE name = 'B') AS to_id,
        100;
    COMMIT;  -- 提交事务
    SELECT '转账成功' AS result;
ELSE
    ROLLBACK;  -- 回滚事务
    select '余额不足，转账失败';  -- 抛出错误[7,9](@ref)
END IF;