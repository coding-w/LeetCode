-- https://leetcode.cn/problems/employees-earning-more-than-their-managers/description/
select t1.name as Employee from Employee t1
inner join Employee t2 on t1.managerId is not null and t1.managerId = t2.id and t1.salary > t2.salary;