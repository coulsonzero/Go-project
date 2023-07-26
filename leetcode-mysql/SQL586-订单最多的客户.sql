# 输入:
# Orders 表:
# +--------------+-----------------+
# | order_number | customer_number |
# +--------------+-----------------+
# | 1            | 1               |
# | 2            | 2               |
# | 3            | 3               |
# | 4            | 3               |
# +--------------+-----------------+
# 输出:
# +-----------------+
# | customer_number |
# +-----------------+
# | 3               |
# +-----------------+
# 解释:
# customer_number 为 '3' 的顾客有两个订单，比顾客 '1' 或者 '2' 都要多，因为他们只有一个订单。
# 所以结果是该顾客的 customer_number ，也就是 3 。

select customer_number from Orders
group by customer_number
order by count(customer_number) desc
limit 1
