# 输入:
# Courses table:
# +---------+----------+
# | student | class    |
# +---------+----------+
# | A       | Math     |
# | B       | English  |
# | C       | Math     |
# | D       | Biology  |
# | E       | Math     |
# | F       | Computer |
# | G       | Math     |
# | H       | Math     |
# | I       | Math     |
# +---------+----------+
# 输出:
# +---------+
# | class   |
# +---------+
# | Math    |
# +---------+
# 解释:
# -数学课有6个学生，所以我们包括它。
# -英语课有1名学生，所以我们不包括它。
# -生物课有1名学生，所以我们不包括它。
# -计算机课有1个学生，所以我们不包括它。


select class from Courses
group by class having count(student) >= 5
;