-- owner
SELECT u.usename
FROM pg_database d
         JOIN pg_user u ON (d.datdba = u.usesysid)
WHERE d.datname = (SELECT current_database());

-- grant
grant all privileges on database dbname to user_name;
grant SELECT ON ALL TABLES IN SCHEMA public to user_name;
GRANT USAGE ON SCHEMA public TO my_user;


-- selected table privileges
SELECT grantee, privilege_type
FROM information_schema.role_table_grants
WHERE table_name='table_name';

-- sequences
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO "user-name";
