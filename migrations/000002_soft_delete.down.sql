ALTER TABLE todoapp.users DROP COLUMN IF EXISTS deleted_at;
ALTER TABLE todoapp.tasks DROP COLUMN IF EXISTS deleted_at;