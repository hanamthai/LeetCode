https://leetcode.com/problems/duplicate-emails/

select email as Email from Person
group by email
having count(*) > 1

-- In this case, using COUNT(*) might be more optimized than COUNT(email) in the HAVING clause because:
-- COUNT(*) simply counts the number of rows and does not concern itself with the values of any specific column, making it generally faster to process.
-- When using COUNT(email), the database management system has to examine each value in the "email" column to count the number of non-null values. This can potentially increase the query execution time, especially if the "email" column is not designated as containing only non-null values.
-- Therefore, in some cases, using COUNT(*) can be more optimized and provide better performance. However, the effectiveness of this optimization also depends on the type of database management system and the specific conditions of the SQL statement and data.