INSERT INTO users("name") VALUES
('vasya'),
('misha'),
('petya');

INSERT INTO events(Owner_id, Title, Text, StartTime, EndTime ) VALUES
(2, 'Phone call','call to vasya', '2020-03-26 09:00:00'::timestamp, '2020-03-26 10:30:00'::timestamp),
(1, 'Drink','Drink cup of water', '2020-03-26 12:00:00'::timestamp, '2020-03-26 12:30:00'::timestamp),
(2, 'Drink','Drink cup of water', '2020-03-26 12:00:00'::timestamp, '2020-03-26 12:30:00'::timestamp),
(3, 'Drink','Drink cup of water', '2020-03-26 12:00:00'::timestamp, '2020-03-26 12:30:00'::timestamp);
