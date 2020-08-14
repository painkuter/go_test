-- неиспользуемые индексы

SELECT * FROM pg_stat_all_indexes
WHERE schemaname = 'public' AND relname = 'product_search';

-- дубли

SELECT array_agg(indexname) AS indexes, replace(indexdef, indexname, '') AS defn
FROM pg_indexes
GROUP BY defn
HAVING count(*) > 1;
