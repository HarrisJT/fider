ALTER TABLE logs RENAME COLUMN created_on TO created_at;
ALTER TABLE post_subscribers RENAME COLUMN created_on TO created_at;
ALTER TABLE post_subscribers RENAME COLUMN updated_on TO updated_at;
ALTER TABLE user_providers RENAME COLUMN created_on TO created_at;
ALTER TABLE post_votes RENAME COLUMN created_on TO created_at;
ALTER TABLE oauth_providers RENAME COLUMN created_on TO created_at;
ALTER TABLE posts RENAME COLUMN created_on TO created_at;
ALTER TABLE uploads RENAME COLUMN created_on TO created_at;
ALTER TABLE tenants RENAME COLUMN created_on TO created_at;
ALTER TABLE post_tags RENAME COLUMN created_on TO created_at;
ALTER TABLE tags RENAME COLUMN created_on TO created_at;
ALTER TABLE notifications RENAME COLUMN created_on TO created_at;
ALTER TABLE notifications RENAME COLUMN updated_on TO updated_at;
ALTER TABLE users RENAME COLUMN created_on TO created_at;
ALTER TABLE comments RENAME COLUMN created_on TO created_at;
ALTER TABLE comments RENAME COLUMN edited_on TO edited_at;
ALTER TABLE email_verifications RENAME COLUMN created_on TO created_at;
ALTER TABLE email_verifications RENAME COLUMN verified_on TO verified_at;
ALTER TABLE email_verifications RENAME COLUMN expires_on TO expires_at;
