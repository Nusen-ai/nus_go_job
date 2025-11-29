-- 创建数据库和表
CREATE DATABASE IF NOT EXISTS book_store;
USE book_store;

-- 创建books表
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 插入测试数据
INSERT INTO books (title, author, price) VALUES
('Go语言编程', '张三', 45.00),
('深入理解计算机系统', '李四', 89.00),
('算法导论', '王五', 128.00),
('数据库系统概念', '赵六', 75.50),
('Clean Code', 'Robert Martin', 65.00),
('设计模式', '四人组', 95.00),
('JavaScript权威指南', 'David Flanagan', 120.00),
('Python编程', 'Mark Lutz', 55.00),
('Java核心技术', 'Cay Horstmann', 88.00),
('C++ Primer', 'Stanley Lippman', 110.00);