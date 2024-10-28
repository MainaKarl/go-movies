-- +goose Up
-- +goose StatementBegin
CREATE TABLE movies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tmdb_id INT UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    release_date DATE,
    language VARCHAR(10),
    popularity DECIMAL(5, 2),
    vote_average DECIMAL(3, 1),
    vote_count INT,
    poster_path VARCHAR(255),
    backdrop_path VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
