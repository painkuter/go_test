select * from priority_test;

-- priority |  id
-- 1	    | 1
-- 2	    | 2
-- 1	    | 3
-- 2	    | 4
-- 1	    | 5
-- 2	    | 6
-- 1	    | 7
-- 2	    | 8
-- 1	    | 9
-- 2	    | 10

-- priority |  id
-- ---------------
-- 2        |   8
-- 2        |   6
-- 2        |   4
-- 2        |   2
-- 1        |   7
-- 1        |   5
-- 1        |   3
-- 1        |   1




select * from priority_test where priority < 1e2 or (priority = 1e2 and id < 1e8) order by priority desc, id desc limit 3;
select * from priority_test where priority < 2 or (priority = 2 and id < 6) order by priority desc, id desc limit 3;
select * from priority_test where priority < 1 or (priority = 1 and id < 9) order by priority desc, id desc limit 3;
select * from priority_test where priority < 1 or (priority = 1 and id < 3) order by priority desc, id desc limit 3;

