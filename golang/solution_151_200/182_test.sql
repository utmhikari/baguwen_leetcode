# 182. 查找重复的电子邮箱
# 编写一个 SQL 查询，查找Person 表中所有重复的电子邮箱。
#
# 示例：
#
# +----+---------+
# | Id | Email   |
# +----+---------+
# | 1  | a@b.com |
# | 2  | c@d.com |
# | 3  | a@b.com |
# +----+---------+
# 根据以上输入，你的查询应返回以下结果：
#
# +---------+
# | Email   |
# +---------+
# | a@b.com |
# +---------+

SELECT DISTINCT P1.Email AS Email
FROM Person P1, Person P2
WHERE P1.Email = P2.Email AND P1.Id != P2.Id
