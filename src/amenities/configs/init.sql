CREATE TABLE IF NOT EXISTS amenities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    start_time TIME,
    end_time TIME
);

INSERT INTO amenities (name, description, start_time, end_time) VALUES
('Pool', 'Indoor pool', '06:00:00', '22:00:00'),
('Gym', '18/7 access gym', '05:00:00', '23:00:00'),
('Breakfast', 'Continental Breakfast', '08:00:00', '20:00:00'),
('Bar', 'Now with Ice Cream!', '07:00:00', '21:00:00');
