-- 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。

-- 创建测试数据库和表
CREATE DATABASE student_test;
USE student_test;

-- 创建students表
CREATE TABLE students (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT,
    grade VARCHAR(50)
);

-- 插入一条新记录，学生姓名为"张三"，年龄为20，年级为"三年级"。
INSERT INTO students (name, age, grade) 
VALUES ('张三', 20, '三年级');



-- 查询所有年龄大于18岁的学生信息。
SELECT * FROM students 
WHERE age > 14;

-- 更新张三的年级为"四年级"。
UPDATE students 
SET grade = '四年级' 
WHERE name = '张三';

-- 删除所有年龄小于15岁的学生记录。
DELETE FROM students 
WHERE age < 15;
