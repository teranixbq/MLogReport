BEGIN;
CREATE TYPE weekly_status AS ENUM (
	'pending',
	'approve'
);
COMMIT;