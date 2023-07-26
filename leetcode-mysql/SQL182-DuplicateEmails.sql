# Input:
# Person table:
# +----+---------+
# | id | email   |
# +----+---------+
# | 1  | a@b.com |
# | 2  | c@d.com |
# | 3  | a@b.com |
# +----+---------+
# Output:
# +---------+
# | Email   |
# +---------+
# | a@b.com |
# +---------+
# Explanation: a@b.com is repeated two times.

SELECT Email from Person
group by Email having count(Email) > 1;
