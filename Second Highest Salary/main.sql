https://leetcode.com/problems/second-highest-salary/

select coalesce(
    (select salary from Employee
    order by salary desc
    limit 1 offset 1),
    null
) SecondHighestSalary