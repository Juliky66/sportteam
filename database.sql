-- Создание таблиц
CREATE TABLE IF NOT EXISTS player_name (
    id SERIAL PRIMARY KEY,
    last_name   VARCHAR(50) NOT NULL,
    first_name  VARCHAR(50) NOT NULL,
    middle_name VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS city (
    id SERIAL PRIMARY KEY,
    city_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS player (
    id SERIAL PRIMARY KEY,
    player_name_id INT NOT NULL REFERENCES player_name(id),
    city_id        INT NOT NULL REFERENCES city(id),
    height_cm      INT NOT NULL,
    weight_kg      NUMERIC(5,2) NOT NULL
);

--  Данные для player_name

INSERT INTO player_name (last_name, first_name, middle_name) VALUES
('Иванов', 'Алексей', 'Сергеевич'),
('Петров', 'Дмитрий', 'Игоревич'),
('Сидоров', 'Никита', 'Андреевич'),
('Кузнецова', 'Анна', 'Владимировна'),
('Смирнов', 'Егор', 'Павлович'),
('Попова', 'Мария', 'Алексеевна'),
('Федоров', 'Илья', 'Олегович'),
('Крылова', 'Екатерина', 'Сергеевна'),
('Соколова', 'Ольга', 'Дмитриевна'),
('Морозов', 'Максим', 'Романович'),
('Васильев', 'Кирилл', 'Андреевич'),
('Новикова', 'Дарья', 'Викторовна');

--  Данные для city

INSERT INTO city (city_name) VALUES
('Москва'),
('Санкт-Петербург'),
('Новосибирск'),
('Екатеринбург'),
('Казань'),
('Нижний Новгород'),
('Красноярск'),
('Самара'),
('Ростов-на-Дону'),
('Омск'),
('Владивосток'),
('Уфа');

--  Данные для player

INSERT INTO player (player_name_id, city_id, height_cm, weight_kg) VALUES
(1, 1, 190, 85.0),
(2, 3, 185, 82.5),
(3, 2, 192, 90.0),
(4, 4, 178, 68.5),
(5, 5, 187, 80.0),
(6, 7, 172, 60.0),
(7, 6, 195, 95.2),
(8, 8, 180, 70.3),
(9, 9, 176, 62.4),
(10, 10, 193, 88.7),
(11, 11, 188, 83.1),
(12, 12, 170, 58.9);
