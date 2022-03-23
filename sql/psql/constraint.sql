-- ограничение по пересечению интервалов time_from, time_to

CREATE EXTENSION btree_gist;

alter table delivery_method_timetable add constraint dmt_time_from_time_to_non_intersec
    EXCLUDE USING gist (
    delivery_method_id WITH =,
    tsrange(time_from, time_to) WITH &&
    );

-- constraints list

SELECT con.*
FROM pg_catalog.pg_constraint con
         INNER JOIN pg_catalog.pg_class rel
                    ON rel.oid = con.conrelid
         INNER JOIN pg_catalog.pg_namespace nsp
                    ON nsp.oid = connamespace
WHERE nsp.nspname = '<schema>'
  AND rel.relname = '<table_name>';