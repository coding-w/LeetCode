-- https://leetcode.cn/problems/department-top-three-salaries/description/
with t1 as (
    select name, DENSE_RANK() OVER (PARTITION BY departmentId order by salary desc) as rank, departmentId, salary from Employee
)
select t2.name as "Department", t1.name as "Employee", t1.salary as "Salary" from t1
inner join Department t2 on t1.rank <= 3 and t1.departmentId = id;