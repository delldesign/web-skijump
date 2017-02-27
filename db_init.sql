-- Roles
--

CREATE USER testuser
  WITH PASSWORD 'md599e5ea7a6f7c3269995cba3927fd0093';

--
-- Database creation
--

CREATE DATABASE testdb
  WITH OWNER testuser;

--
-- Access rights
--

REVOKE ALL ON DATABASE testdb FROM PUBLIC;
GRANT ALL ON DATABASE testdb TO testuser;

\connect testdb

REVOKE ALL ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO testuser;

GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO testuser;