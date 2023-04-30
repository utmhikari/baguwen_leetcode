
# 197. 上升的温度
# 编写一个 SQL 查询，来查找与之前（昨天的）日期相比温度更高的所有日期的 id 。


SELECT w1.id AS id FROM Weather w1, Weather w2
WHERE dateDiff(w1.recordDate, w2.recordDate) = 1 AND w1.Temperature > w2.Temperature