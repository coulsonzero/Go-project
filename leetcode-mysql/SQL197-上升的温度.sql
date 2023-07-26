# 输入：
# Weather 表：
# +----+------------+-------------+
# | id | recordDate | Temperature |
# +----+------------+-------------+
# | 1  | 2015-01-01 | 10          |
# | 2  | 2015-01-02 | 25          |
# | 3  | 2015-01-03 | 20          |
# | 4  | 2015-01-04 | 30          |
# +----+------------+-------------+
# 输出：
# +----+
# | id |
# +----+
# | 2  |
# | 4  |
# +----+
# 解释：
# 2015-01-02 的温度比前一天高（10 -> 25）
# 2015-01-04 的温度比前一天高（20 -> 30）


select a.id
from weather as a cross join weather as b
on datediff(a.recordDate, b.recordDate) = 1
where a.Temperature > b.Temperature;
