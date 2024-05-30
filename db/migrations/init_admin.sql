CREATE USER netrunnerflix_admin WITH PASSWORD 'admin';

GRANT CONNECT ON DATABASE netrunnerflix TO netrunnerflix_admin;
GRANT USAGE ON SCHEMA public TO netrunnerflix_admin;

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO netrunnerflix_admin;
