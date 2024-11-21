-- https://leetcode.cn/problems/rank-scores/
select score, DENSE_RANK() over(order by score desc) as rank from Scores;