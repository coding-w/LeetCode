-- https://leetcode.cn/problems/duplicate-emails/description/
select email as Email from Person group by email having count(*) > 1;
