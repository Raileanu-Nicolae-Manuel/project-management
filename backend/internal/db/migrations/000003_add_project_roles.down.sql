-- Remove the foreign key and index first
ALTER TABLE user_projects 
    DROP FOREIGN KEY fk_user_projects_role,
    DROP INDEX idx_user_projects_role_id;

-- Revert the column change
ALTER TABLE user_projects 
    DROP COLUMN role_id,
    ADD COLUMN role VARCHAR(50) NOT NULL DEFAULT 'member';

-- Drop the roles table
DROP TABLE IF EXISTS project_roles;
