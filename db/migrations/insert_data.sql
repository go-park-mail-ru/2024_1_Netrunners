INSERT INTO actor (name, avatar, career, height, birth_place, genres, spouse) VALUES ('Роберт Де Ниро', 'https://upload.wikimedia.org/wikipedia/commons/5/58/Robert_De_Niro_Cannes_2016.jpg', 'Американский актёр, продюсер и режиссёр', 165, 'Гринвич-Виллидж, Манхэттен, Нью-Йорк, Нью-Йорк, США', 'Криминал, Драма, Боевик', 'Женат на Грейс Хайтауэр');
INSERT INTO actor (name, avatar, career, height, birth_place, genres, spouse) VALUES ('Рэймонд Аллен Лиотта', 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Ray_Liotta_%288140672892%29.jpg/274px-Ray_Liotta_%288140672892%29.jpg', 'Американский актёр кино и озвучивания', 183, 'Ньюарк, Нью-Джерси', 'Криминал, Драма', 'Помолвлен Джейси Ниттоло');
INSERT INTO actor (name, avatar, career, height, birth_place, genres, spouse) VALUES ('Джозеф Фрэнк Пеши', 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/37/JoePesci-2009.jpg/274px-JoePesci-2009.jpg', 'Американский актёр, комик и певец', 156, 'Ньюарк, Нью-Джерси, США', 'Криминал, Комедия, Драма', 'Холост');

INSERT INTO director (name, avatar) VALUES ('Мартин Чарльз Скорсезе', 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/ce/Martin_Scorsese_MFF_2023.jpg/800px-Martin_Scorsese_MFF_2023.jpg');

INSERT INTO film (title, data, banner, s3_link, director, age_limit, duration) VALUES ('Славные парни', 'История о Генри Хилле — начинающем гангстере, занимающемся грабежами вместе с подельниками Джими Конвеем и Томми Де Вито, которые с легкостью убивают любого, кто встаёт у них на пути.', 'https://avatars.mds.yandex.net/get-kinopoisk-image/1900788/9d56c458-1c44-4da0-b718-2899ccbf6b5b/300x', '', 1, 16, 127);

INSERT INTO comment (text, score, author, film) VALUES ('один из любимых моих фильмов', 5, 1, 1);
INSERT INTO comment (text, score, author, film) VALUES ('проходняк)))', 5, 1, 1);
INSERT INTO comment (text, score, author, film) VALUES ('классика, что тут еще говорить', 5, 1, 1);

INSERT INTO film_actor (film, actor) VALUES (1, 1);
INSERT INTO film_actor (film, actor) VALUES (1, 2);
INSERT INTO film_actor (film, actor) VALUES (1, 3);

