-- +goose Up
-- +goose StatementBegin
		CREATE TABLE IF NOT EXISTS services (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			price DECIMAL(10, 2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP
		);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  DROP TABLES SERVICES;
-- +goose StatementEnd
