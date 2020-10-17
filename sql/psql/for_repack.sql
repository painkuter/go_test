------------------------------------------------------------------------------
WITH data AS (
    SELECT current_database() AS current_database,
           sub.nspname AS schemaname,
           sub.tblname,
           sub.idxname,
           sub.bs * sub.relpages::bigint::numeric AS real_size,
           sub.bs * (sub.relpages::double precision - sub.est_pages)::bigint::numeric AS extra_size,
           100::double precision * (sub.relpages::double precision - sub.est_pages) / sub.relpages::double precision AS extra_ratio,
           sub.fillfactor,
           sub.bs::double precision * (sub.relpages::double precision - sub.est_pages_ff) AS bloat_size,
           100::double precision * (sub.relpages::double precision - sub.est_pages_ff) / sub.relpages::double precision AS bloat_ratio,
           sub.is_na
    FROM ( SELECT COALESCE(1::double precision + ceil(s2.reltuples / floor((s2.bs - s2.pageopqdata::numeric - s2.pagehdr::numeric)::double precision / (4::numeric + s2.nulldatahdrwidth)::double precision)), 0::double precision) AS est_pages,
                  COALESCE(1::double precision + ceil(s2.reltuples / floor(((s2.bs - s2.pageopqdata::numeric - s2.pagehdr::numeric) * s2.fillfactor::numeric)::double precision / (100::double precision * (4::numeric + s2.nulldatahdrwidth)::double precision))), 0::double precision) AS est_pages_ff,
                  s2.bs,
                  s2.nspname,
                  s2.table_oid,
                  s2.tblname,
                  s2.idxname,
                  s2.relpages,
                  s2.fillfactor,
                  s2.is_na
           FROM ( SELECT s1.maxalign,
                         s1.bs,
                         s1.nspname,
                         s1.tblname,
                         s1.idxname,
                         s1.reltuples,
                         s1.relpages,
                         s1.relam,
                         s1.table_oid,
                         s1.fillfactor,
                         ((s1.index_tuple_hdr_bm + s1.maxalign -
                           CASE
                               WHEN (s1.index_tuple_hdr_bm % s1.maxalign) = 0 THEN s1.maxalign
                               ELSE s1.index_tuple_hdr_bm % s1.maxalign
                               END)::double precision + s1.nulldatawidth + s1.maxalign::double precision -
                          CASE
                              WHEN s1.nulldatawidth = 0::double precision THEN 0
                              WHEN (s1.nulldatawidth::integer % s1.maxalign) = 0 THEN s1.maxalign
                              ELSE s1.nulldatawidth::integer % s1.maxalign
                              END::double precision)::numeric AS nulldatahdrwidth,
                         s1.pagehdr,
                         s1.pageopqdata,
                         s1.is_na
                  FROM ( SELECT i.nspname,
                                i.tblname,
                                i.idxname,
                                i.reltuples,
                                i.relpages,
                                i.relam,
                                a.attrelid AS table_oid,
                                current_setting('block_size'::text)::numeric AS bs,
                                i.fillfactor,
                                CASE
                                    WHEN version() ~ 'mingw32'::text OR version() ~ '64-bit|x86_64|ppc64|ia64|amd64'::text THEN 8
                                    ELSE 4
                                    END AS maxalign,
                                24 AS pagehdr,
                                16 AS pageopqdata,
                                CASE
                                    WHEN max(COALESCE(s.null_frac, 0::real)) = 0::double precision THEN 2
                                    ELSE 2 + (32 + 8 - 1) / 8
                                    END AS index_tuple_hdr_bm,
                                sum((1::double precision - COALESCE(s.null_frac, 0::real)) * COALESCE(s.avg_width, 1024)::double precision) AS nulldatawidth,
                                max(
                                        CASE
                                            WHEN a.atttypid = 'name'::regtype::oid THEN 1
                                            ELSE 0
                                            END) > 0 AS is_na
                         FROM pg_attribute a
                                  JOIN ( SELECT pg_namespace.nspname,
                                                tbl.relname AS tblname,
                                                idx.relname AS idxname,
                                                idx.reltuples,
                                                idx.relpages,
                                                idx.relam,
                                                pg_index.indrelid,
                                                pg_index.indexrelid,
                                                pg_index.indkey::smallint[] AS attnum,
                                                COALESCE("substring"(array_to_string(idx.reloptions, ' '::text), 'fillfactor=([0-9]+)'::text)::smallint::integer, 90) AS fillfactor
                                         FROM pg_index
                                                  JOIN pg_class idx ON idx.oid = pg_index.indexrelid
                                                  JOIN pg_class tbl ON tbl.oid = pg_index.indrelid
                                                  JOIN pg_namespace ON pg_namespace.oid = idx.relnamespace
                                         WHERE pg_index.indisvalid AND tbl.relkind = 'r'::"char" AND idx.relpages > 0) i ON a.attrelid = i.indexrelid
                                  JOIN pg_stats s ON s.schemaname = i.nspname AND (s.tablename = i.tblname AND s.attname::text = pg_get_indexdef(a.attrelid, a.attnum::integer, true) OR s.tablename = i.idxname AND s.attname = a.attname)
                                  JOIN pg_type t ON a.atttypid = t.oid
                         WHERE a.attnum > 0
                         GROUP BY i.nspname, i.tblname, i.idxname, i.reltuples, i.relpages, i.relam, a.attrelid, (current_setting('block_size'::text)::numeric), i.fillfactor) s1) s2
                    JOIN pg_am am ON s2.relam = am.oid
           WHERE am.amname = 'btree'::name) sub
)
SELECT data.current_database,
       data.schemaname,
       data.tblname,
       data.idxname,
       data.real_size,
       pg_size_pretty(data.real_size) AS real_size_pretty,
       data.extra_size,
       pg_size_pretty(data.extra_size) AS extra_size_pretty,
       data.extra_ratio AS "extra_ratio, %",
       data.bloat_size,
       pg_size_pretty(data.bloat_size::numeric) AS bloat_size_pretty,
       data.bloat_ratio AS "bloat_ratio, %",
       data.fillfactor,
       data.is_na,
       data.real_size::double precision - data.bloat_size AS live_data_size
FROM data
ORDER BY data.bloat_size DESC;

------------------------------------------------------------------------------
------------------------------------------------------------------------------

WITH constants AS (
    -- define some constants for sizes of things
    -- for reference down the query and easy maintenance
    SELECT current_setting('block_size')::numeric AS bs, 23 AS hdr, 8 AS ma
),
     no_stats AS (
         -- screen out table who have attributes
         -- which dont have stats, such as JSON
         SELECT table_schema, table_name,
                n_live_tup::numeric as est_rows,
                pg_table_size(relid)::numeric as table_size
         FROM information_schema.columns
                  JOIN pg_stat_user_tables as psut
                       ON table_schema = psut.schemaname
                           AND table_name = psut.relname
                  LEFT OUTER JOIN pg_stats
                                  ON table_schema = pg_stats.schemaname
                                      AND table_name = pg_stats.tablename
                                      AND column_name = attname
         WHERE attname IS NULL
           AND table_schema NOT IN ('pg_catalog', 'information_schema')
         GROUP BY table_schema, table_name, relid, n_live_tup
     ),
     null_headers AS (
         -- calculate null header sizes
         -- omitting tables which dont have complete stats
         -- and attributes which aren't visible
         SELECT
                 hdr+1+(sum(case when null_frac <> 0 THEN 1 else 0 END)/8) as nullhdr,
                 SUM((1-null_frac)*avg_width) as datawidth,
                 MAX(null_frac) as maxfracsum,
                 schemaname,
                 tablename,
                 hdr, ma, bs
         FROM pg_stats CROSS JOIN constants
                       LEFT OUTER JOIN no_stats
                                       ON schemaname = no_stats.table_schema
                                           AND tablename = no_stats.table_name
         WHERE schemaname NOT IN ('pg_catalog', 'information_schema')
           AND no_stats.table_name IS NULL
           AND EXISTS ( SELECT 1
                        FROM information_schema.columns
                        WHERE schemaname = columns.table_schema
                          AND tablename = columns.table_name )
         GROUP BY schemaname, tablename, hdr, ma, bs
     ),
     data_headers AS (
         -- estimate header and row size
         SELECT
             ma, bs, hdr, schemaname, tablename,
             (datawidth+(hdr+ma-(case when hdr%ma=0 THEN ma ELSE hdr%ma END)))::numeric AS datahdr,
             (maxfracsum*(nullhdr+ma-(case when nullhdr%ma=0 THEN ma ELSE nullhdr%ma END))) AS nullhdr2
         FROM null_headers
     ),
     table_estimates AS (
         -- make estimates of how large the table should be
         -- based on row and page size
         SELECT schemaname, tablename, bs,
                reltuples::numeric as est_rows, relpages * bs as table_bytes,
                CEIL((reltuples*
                      (datahdr + nullhdr2 + 4 + ma -
                       (CASE WHEN datahdr%ma=0
                                 THEN ma ELSE datahdr%ma END)
                          )/(bs-20))) * bs AS expected_bytes,
                reltoastrelid
         FROM data_headers
                  JOIN pg_class ON tablename = relname
                  JOIN pg_namespace ON relnamespace = pg_namespace.oid
             AND schemaname = nspname
         WHERE pg_class.relkind = 'r'
     ),
     estimates_with_toast AS (
         -- add in estimated TOAST table sizes
         -- estimate based on 4 toast tuples per page because we dont have
         -- anything better.  also append the no_data tables
         SELECT schemaname, tablename,
                TRUE as can_estimate,
                est_rows,
                table_bytes + ( coalesce(toast.relpages, 0) * bs ) as table_bytes,
                expected_bytes + ( ceil( coalesce(toast.reltuples, 0) / 4 ) * bs ) as expected_bytes
         FROM table_estimates LEFT OUTER JOIN pg_class as toast
                                              ON table_estimates.reltoastrelid = toast.oid
                                                  AND toast.relkind = 't'
     ),
     table_estimates_plus AS (
-- add some extra metadata to the table data
-- and calculations to be reused
-- including whether we cant estimate it
-- or whether we think it might be compressed
         SELECT current_database() as databasename,
                schemaname, tablename, can_estimate,
                est_rows,
                CASE WHEN table_bytes > 0
                         THEN table_bytes::NUMERIC
                     ELSE NULL::NUMERIC END
                                   AS table_bytes,
                CASE WHEN expected_bytes > 0
                         THEN expected_bytes::NUMERIC
                     ELSE NULL::NUMERIC END
                                   AS expected_bytes,
                CASE WHEN expected_bytes > 0 AND table_bytes > 0
                    AND expected_bytes <= table_bytes
                         THEN (table_bytes - expected_bytes)::NUMERIC
                     ELSE 0::NUMERIC END AS bloat_bytes
         FROM estimates_with_toast
         UNION ALL
         SELECT current_database() as databasename,
                table_schema, table_name, FALSE,
                est_rows, table_size,
                NULL::NUMERIC, NULL::NUMERIC
         FROM no_stats
     ),
     bloat_data AS (
         -- do final math calculations and formatting
         select current_database() as databasename,
                schemaname, tablename, can_estimate,
                table_bytes, round(table_bytes/(1024^2)::NUMERIC,3) as table_mb,
                expected_bytes, round(expected_bytes/(1024^2)::NUMERIC,3) as expected_mb,
                round(bloat_bytes*100/table_bytes) as pct_bloat,
                round(bloat_bytes/(1024::NUMERIC^2),2) as mb_bloat,
                table_bytes, expected_bytes, est_rows
         FROM table_estimates_plus
     )
-- filter output for bloated tables
SELECT databasename, schemaname, tablename,
       can_estimate,
       est_rows,
       pct_bloat, mb_bloat,
       table_mb
FROM bloat_data
-- this where clause defines which tables actually appear
-- in the bloat chart
-- example below filters for tables which are either 50%
-- bloated and more than 20mb in size, or more than 25%
-- bloated and more than 1GB in size
WHERE ( pct_bloat >= 50 AND mb_bloat >= 20 )
   OR ( pct_bloat >= 25 AND mb_bloat >= 1000 )
ORDER BY pct_bloat DESC;

------------------------------------------------------------------------------
------------------------------------------------------------------------------

SELECT current_database() AS current_database,
       sml.schemaname,
       sml.tablename,
       round(
               CASE
                   WHEN sml.otta = 0::double precision THEN 0.0::double precision
                   ELSE sml.relpages::double precision / sml.otta
                   END::numeric, 1) AS tbloat,
       CASE
           WHEN sml.relpages::double precision < sml.otta THEN 0::numeric
           ELSE sml.bs * (sml.relpages::double precision - sml.otta)::bigint::numeric
           END AS wastedbytes,
       sml.iname,
       round(
               CASE
                   WHEN sml.iotta = 0::double precision OR sml.ipages = 0 THEN 0.0::double precision
                   ELSE sml.ipages::double precision / sml.iotta
                   END::numeric, 1) AS ibloat,
       CASE
           WHEN sml.ipages::double precision < sml.iotta THEN 0::double precision
           ELSE sml.bs::double precision * (sml.ipages::double precision - sml.iotta)
           END AS wastedibytes
FROM ( SELECT rs.schemaname,
              rs.tablename,
              cc.reltuples,
              cc.relpages,
              rs.bs,
              ceil(cc.reltuples * ((rs.datahdr + rs.ma::numeric -
                                    CASE
                                        WHEN (rs.datahdr % rs.ma::numeric) = 0::numeric THEN rs.ma::numeric
                                        ELSE rs.datahdr % rs.ma::numeric
                                        END)::double precision + rs.nullhdr2 + 4::double precision) / (rs.bs::double precision - 20::double precision)) AS otta,
              COALESCE(c2.relname, '?'::name) AS iname,
              COALESCE(c2.reltuples, 0::real) AS ituples,
              COALESCE(c2.relpages, 0) AS ipages,
              COALESCE(ceil(c2.reltuples * (rs.datahdr - 12::numeric)::double precision / (rs.bs::double precision - 20::double precision)), 0::double precision) AS iotta
       FROM ( SELECT foo.ma,
                     foo.bs,
                     foo.schemaname,
                     foo.tablename,
                     (foo.datawidth + (foo.hdr + foo.ma -
                                       CASE
                                           WHEN (foo.hdr % foo.ma) = 0 THEN foo.ma
                                           ELSE foo.hdr % foo.ma
                                           END)::double precision)::numeric AS datahdr,
                     foo.maxfracsum * (foo.nullhdr + foo.ma -
                                       CASE
                                           WHEN (foo.nullhdr % foo.ma::bigint) = 0 THEN foo.ma::bigint
                                           ELSE foo.nullhdr % foo.ma::bigint
                                           END)::double precision AS nullhdr2
              FROM ( SELECT s.schemaname,
                            s.tablename,
                            constants.hdr,
                            constants.ma,
                            constants.bs,
                            sum((1::double precision - s.null_frac) * s.avg_width::double precision) AS datawidth,
                            max(s.null_frac) AS maxfracsum,
                            constants.hdr + (( SELECT 1 + count(*) / 8
                                               FROM pg_stats s2
                                               WHERE s2.null_frac <> 0::double precision AND s2.schemaname = s.schemaname AND s2.tablename = s.tablename)) AS nullhdr
                     FROM pg_stats s,
                          ( SELECT ( SELECT current_setting('block_size'::text)::numeric AS current_setting) AS bs,
                                   CASE
                                       WHEN "substring"(foo_1.v, 12, 3) = ANY (ARRAY['8.0'::text, '8.1'::text, '8.2'::text]) THEN 27
                                       ELSE 23
                                       END AS hdr,
                                   CASE
                                       WHEN foo_1.v ~ 'mingw32'::text THEN 8
                                       ELSE 4
                                       END AS ma
                            FROM ( SELECT version() AS v) foo_1) constants
                     GROUP BY s.schemaname, s.tablename, constants.hdr, constants.ma, constants.bs) foo) rs
                JOIN pg_class cc ON cc.relname = rs.tablename
                JOIN pg_namespace nn ON cc.relnamespace = nn.oid AND nn.nspname = rs.schemaname AND nn.nspname <> 'information_schema'::name
                LEFT JOIN pg_index i ON i.indrelid = cc.oid
                LEFT JOIN pg_class c2 ON c2.oid = i.indexrelid) sml
ORDER BY (
             CASE
                 WHEN sml.relpages::double precision < sml.otta THEN 0::numeric
                 ELSE sml.bs * (sml.relpages::double precision - sml.otta)::bigint::numeric
                 END) DESC;
