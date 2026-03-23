CREATE TABLE IF NOT EXISTS tasks (
                                     id SERIAL PRIMARY KEY,
                                     type TEXT NOT NULL,
                                     payload TEXT,
                                     status TEXT NOT NULL,
                                     result TEXT,
                                     error TEXT,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);