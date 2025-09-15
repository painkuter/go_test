SELECT table,
    formatReadableSize(sum(bytes)) as size,
    sum(rows)                      as rows,
    count()                        as parts
FROM system.parts
WHERE active
GROUP BY table
ORDER BY sum(bytes) DESC;