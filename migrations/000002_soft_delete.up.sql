ALTER TABLE todoapp.users ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE todoapp.tasks ADD COLUMN deleted_at TIMESTAMP;