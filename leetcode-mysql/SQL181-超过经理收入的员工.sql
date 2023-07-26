# 输入:
# Employee 表:
# +----+-------+--------+-----------+
# | id | name  | salary | managerId |
# +----+-------+--------+-----------+
# | 1  | Joe   | 70000  | 3         |
# | 2  | Henry | 80000  | 4         |
# | 3  | Sam   | 60000  | Null      |
# | 4  | Max   | 90000  | Null      |
# +----+-------+--------+-----------+
# 输出:
# +----------+
# | Employee |
# +----------+
# | Joe      |
# +----------+
# 解释: Joe 是唯一挣得比经理多的雇员。


select e.Name as Employee
from employee e
where salary > (select salary from employee where Id = e.ManagerId)