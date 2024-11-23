-- https://leetcode.cn/problems/delete-duplicate-emails/
delete from Person where id not in (select min(id) as id from Person group by email)