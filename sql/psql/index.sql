-- список индексов
SELECT tablename, indexname, indexdef FROM pg_indexes WHERE schemaname = 'public';

-- использование индексов

SELECT * FROM pg_stat_all_indexes
WHERE schemaname = 'public' AND relname = 'product_search';

-- дубли

SELECT array_agg(indexname) AS indexes, replace(indexdef, indexname, '') AS defn
FROM pg_indexes
GROUP BY defn
HAVING count(*) > 1;

-- broken indexes
SELECT * FROM pg_class, pg_index WHERE pg_index.indisvalid = false AND pg_index.indexrelid = pg_class.oid;

