CREATE TABLE project_roles (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default roles
INSERT INTO project_roles (name, description) VALUES
    ('owner', 'Full access to the project and can manage team members'),
    ('developer', 'Can contribute to the project and access most features'),
    ('tester', 'Can view and test the project');

-- Modify user_projects table to use role_id instead of role string
ALTER TABLE user_projects 
    DROP COLUMN role,
    ADD COLUMN role_id BIGINT NOT NULL,
    ADD CONSTRAINT fk_user_projects_role 
        FOREIGN KEY (role_id) 
        REFERENCES project_roles(id);

-- Add index for role lookups
CREATE INDEX idx_user_projects_role_id ON user_projects(role_id);
