-- Active: 1764308217441@@localhost@3306@bank_transfer
-- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
-- ● 要求 ：
--   ○ 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
--   ○ 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
--   ○ 如果余额不足，则回滚事务。

-- 创建数据库
CREATE DATABASE IF NOT EXISTS bank_transfer;
USE bank_transfer;

-- 创建账户表
CREATE TABLE accounts (
    id VARCHAR(50) PRIMARY KEY,
    balance DECIMAL(10,2) NOT NULL DEFAULT 0.00
);

-- 创建交易记录表
CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    from_account_id VARCHAR(50) NOT NULL,
    to_account_id VARCHAR(50) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    transaction_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 插入测试数据
INSERT INTO accounts (id, balance) VALUES 
('A', 500.00),
('B', 300.00),
('C', 100.00);

USE bank_transfer;

DELIMITER //

CREATE PROCEDURE transfer_money(
    IN from_account VARCHAR(50),
    IN to_account VARCHAR(50),
    IN transfer_amount DECIMAL(10,2)
)
BEGIN
    DECLARE current_balance DECIMAL(10,2);
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
        RESIGNAL;
    END;
    
    START TRANSACTION;
    
    -- 检查转出账户是否存在并锁定该行
    SELECT balance INTO current_balance 
    FROM accounts 
    WHERE id = from_account FOR UPDATE;
    
    -- 如果转出账户不存在，current_balance将为NULL
    IF current_balance IS NULL THEN
        ROLLBACK;
        SELECT '转出账户不存在' AS result;
    ELSEIF current_balance >= transfer_amount THEN
        -- 执行扣款
        UPDATE accounts 
        SET balance = balance - transfer_amount 
        WHERE id = from_account;
        
        -- 执行存款
        UPDATE accounts 
        SET balance = balance + transfer_amount 
        WHERE id = to_account;
        
        -- 记录交易
        INSERT INTO transactions (from_account_id, to_account_id, amount) 
        VALUES (from_account, to_account, transfer_amount);
        
        COMMIT;
        SELECT '转账成功' AS result;
    ELSE
        ROLLBACK;
        SELECT '余额不足，转账失败' AS result;
    END IF;
END //

DELIMITER ;


USE bank_transfer;

-- 场景1：正常转账测试
SELECT '=== 场景1：正常转账测试 ===' AS info;
-- 转账前查看余额
SELECT * FROM accounts;

-- 执行转账：A 向 B 转账 100
CALL transfer_money('A', 'B', 100);

-- 转账后查看余额和交易记录
SELECT * FROM accounts;
SELECT * FROM transactions;

-- 场景2：余额不足测试
SELECT '=== 场景2：余额不足测试 ===' AS info;
SELECT * FROM accounts WHERE id = 'C';
-- 尝试转账超过余额的金额
CALL transfer_money('C', 'A', 200);
-- 检查余额是否保持不变
SELECT * FROM accounts WHERE id = 'C';


-- 尝试向不存在的账户转账
CALL transfer_money('A', 'X', 50);