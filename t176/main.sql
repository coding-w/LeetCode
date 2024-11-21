-- https://leetcode.cn/problems/second-highest-salary/description/
select (select distinct salary from Employee order by salary desc limit 1,1) as SecondHighestSalary
