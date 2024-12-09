-- https://leetcode.cn/problems/trips-and-users/description/
with t2 as (
    select id, status,client_id, request_at from Trips where client_id not in (select users_id from Users where role = 'client' and banned = 'Yes') and driver_id not in (select users_id from Users where role = 'driver' and banned = 'Yes') and request_at in ('2013-10-01','2013-10-02','2013-10-03')
)
select request_at as "Day", round(1.0*sum((status != 'completed')::int)/count(*),2) as "Cancellation Rate"  from t2 group by request_at