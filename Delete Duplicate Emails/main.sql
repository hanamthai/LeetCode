https://leetcode.com/problems/delete-duplicate-emails/

delete from Person p1
using Person p2
where p1.id > p2.id and p1.email = p2.email;
select * from Person

delete from Person
where id in (
    select id from (
        select id, row_number() over(partition by email order by id asc) as row_num
        from Person
    ) p1 where p1.row_num > 1
)