CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE,
    token_hash TEXT UNIQUE NOT NULL -- Storing hash instead of value. Attacker could steal value and use it to impersonate user.
);