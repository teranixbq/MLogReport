BEGIN;
CREATE TYPE role_type AS ENUM (
	'admin',
	'advisor'
);
COMMIT;