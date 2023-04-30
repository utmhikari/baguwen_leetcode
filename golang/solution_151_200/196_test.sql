


-- 196. 删除重复的电子邮箱
-- 编写一个 SQL 查询，来删除 Person 表中所有重复的电子邮箱，重复的邮箱里只保留 Id 最小 的那个。


DELETE p1 FROM Person p1, Person p2 WHERE p1.Email = p2.Email AND p1.ID > p2.ID
