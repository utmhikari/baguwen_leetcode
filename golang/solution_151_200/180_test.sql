# 180. 连续出现的数字
# 表：Logs
#
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | num         | varchar |
# +-------------+---------+
# id 是这个表的主键。
#
#
# 编写一个 SQL 查询，查找所有至少连续出现三次的数字。
#
# 返回的结果表中的数据可以按 任意顺序 排列。
# Logs 表：
# +----+-----+
# | Id | Num |
# +----+-----+
# | 1  | 1   |
# | 2  | 1   |
# | 3  | 1   |
# | 4  | 2   |
# | 5  | 1   |
# | 6  | 2   |
# | 7  | 2   |
# +----+-----+
#
# Result 表：
# +-----------------+
# | ConsecutiveNums |
# +-----------------+
# | 1               |
# +-----------------+
# 1 是唯一连续出现至少三次的数字。

SELECT DISTINCT L1.Num AS ConsecutiveNums
FROM Logs L1, Logs L2, Logs L3
WHERE L1.Num = L2.Num AND
      L2.Num = L3.Num AND
      L1.Id = L2.Id - 1 AND
      L2.Id = L3.Id - 1
