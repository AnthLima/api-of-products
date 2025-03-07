BEGIN;

CREATE TABLE profiles (
    id UUID PRIMARY KEY,
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    nickname VARCHAR(255) NOT NULL,
    profile_image VARCHAR(255),
    user_type VARCHAR(50) NOT NULL CHECK (user_type IN ('consumer', 'seller')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

COMMIT;