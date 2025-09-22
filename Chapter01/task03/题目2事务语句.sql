-- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
-- （包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
-- 要求 ：
-- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
-- 在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
-- 如果余额不足，则回滚事务。

CREATE TABLE accounts
(
    id      INT AUTO_INCREMENT PRIMARY KEY,
    name    VARCHAR(50)    NOT NULL,
    balance DECIMAL(10, 2) NOT NULL
);

CREATE TABLE transactions
(
    id              INT AUTO_INCREMENT PRIMARY KEY,
    from_account_id INT,
    to_account_id   INT,
    amount          DECIMAL(10, 2) NOT NULL
);

INSERT INTO `test`.`accounts` (`id`, `name`, `balance`)
VALUES (1, 'a', 500.00);
INSERT INTO `test`.`accounts` (`id`, `name`, `balance`)
VALUES (2, 'b', 100.00);


DELIMITER $$

CREATE PROCEDURE TransferMoney(
    IN from_acc_id INT,
    IN to_acc_id INT,
    IN transfer_amount DECIMAL(10,2)
)
BEGIN -- 开始事务
    SET autocommit = 0; -- 关闭自动提交
    IF (SELECT balance FROM accounts WHERE id = from_acc_id) >= transfer_amount THEN
        -- 账户A的余额充足，可以进行转账
        UPDATE accounts
        SET balance = balance - transfer_amount
        WHERE id = from_acc_id;

        UPDATE accounts
        SET balance = balance + transfer_amount
        WHERE id = to_acc_id;

        INSERT INTO transactions
        VALUES (NULL, from_acc_id, to_acc_id, transfer_amount);

        COMMIT; -- 提交事务
        SELECT '转账成功' AS result;
    ELSE
        -- 账户A的余额不足，回滚事务
        ROLLBACK;
        SELECT '余额不足，转账失败' AS result;
    END IF;
END  $$

DELIMITER ;

CALL TransferMoney(1, 2, 100);
