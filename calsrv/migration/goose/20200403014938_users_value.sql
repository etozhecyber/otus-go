-- +goose Up
INSERT INTO users("name") VALUES
('vasya'),
('misha'),
('petya');

-- +goose Down
