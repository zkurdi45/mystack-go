-- internal/migrations/sql/000001_create_users_and_admins_tables.up.sql
CREATE TABLE admins (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending_approval', -- e.g., pending_approval, active, suspended
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- We'll create an index on the user status for faster lookups
CREATE INDEX ON users (status);
