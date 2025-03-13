-- +goose Up
-- +goose StatementBegin
CRAETE TABLE bookings (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES users(id),
  timeslot_id INT NOT NULL UNIQUE REFERENCES timeslots(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bookings;
-- +goose StatementEnd
