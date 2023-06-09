# 176. 第二高的薪水
# 编写一个 SQL 查询，获取 Employee表中第二高的薪水（Salary）。
#
# +----+--------+
# | Id | Salary |
# +----+--------+
# | 1  | 100    |
# | 2  | 200    |
# | 3  | 300    |
# +----+--------+
# 例如上述Employee表，SQL查询应该返回200 作为第二高的薪水。如果不存在第二高的薪水，那么查询应返回 null。
#
# +---------------------+
# | SecondHighestSalary |
# +---------------------+
# | 200                 |
# +---------------------+

SELECT
    (SELECT DISTINCT Salary
    FROM Employee
    ORDER BY Salary DESC
    LIMIT 1 OFFSET 1)
AS SecondHighestSalary
