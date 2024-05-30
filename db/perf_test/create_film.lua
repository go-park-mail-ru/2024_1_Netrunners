wrk.headers["Content-Type"] = "application/json"
wrk.method = "POST"
wrk.path = "/api/films/add"

wrk.body = '{\
    "filmData" : {\
        "title": "Форсажжжжиесть",\
	    "preview": "https://upload.wikimedia.org/wikipedia/ru/3/3b/The_Fast_and_the_Furious.jpg",\
	    "director": "Роб Коин",\
	    "data": "Его зовут Брайан, но это не точно, и он — фанат турбин и нитроускорителей. Он пытается попасть в автобанду легендарного Доминика Торетто, чемпиона опасных и незаконных уличных гонок. Брайан также полицейский, и его задание — втереться в доверие к Торетто, подозреваемому в причастности к дерзким грабежам грузовиков, совершаемым прямо на ходу.",\
	    "ageLimit": 16,\
	    "duration": 106,\
	    "publishedAt": "2001-06-22T00:00:00Z",\
	    "genres": ["боевик", "смешарик", "триллер", "криминал"],\
        "link": "https://daimnefilm.hb.ru-msk.vkcs.cloud/%D0%A4%D0%BE%D1%80%D1%81%D0%B0%D0%B6%20%282001%29%C2%A0%E2%80%94%20%D1%80%D1%83%D1%81%D1%81%D0%BA%D0%B8%D0%B9%20%D1%82%D1%80%D0%B5%D0%B9%D0%BB%D0%B5%D1%80.mp4"\
    },\
    "actors": [\
        {\
            "name": "Вин Масло",\
            "avatar": "https://upload.wikimedia.org/wikipedia/commons/thumb/8/83/Vin_Diesel_by_Gage_Skidmore_2.jpg/1920px-Vin_Diesel_by_Gage_Skidmore_2.jpg",\
            "birthday": "1967-07-18T00:00:00Z",\
            "career": "Актер, продюсер, режиссер, сценарист",\
            "height": 182,\
            "birthPlace": "Нью-Йорк, Нью-Йорк, США",\
            "spouse": "Палома Хименес (с 2007 г.)"\
        },\
        {\
            "name": "Потолок Уокер",\
            "avatar": "https://upload.wikimedia.org/wikipedia/commons/9/91/PaulWalkerEdit-1.jpg",\
            "birthday": "1973-09-12T00:00:00Z",\
            "career": "Актер, продюсер, модель",\
            "height": 187,\
            "birthPlace": "Глендейл, Калифорния, США",\
            "spouse": "Ребекка МакБрайд (с 1998 г.)"\
        },\
        {\
            "name": "Мишень Родригес",\
            "avatar": "https://upload.wikimedia.org/wikipedia/commons/7/78/Michelle_Rodriguez_Cannes_2015.jpg",\
            "birthday": "1978-07-12T00:00:00Z",\
            "career": "Актриса",\
            "height": 165,\
            "birthPlace": "Сан-Антонио, Техас, США",\
            "spouse": "Одинока"\
        },\
        {\
            "name": "Аааа аааа",\
            "avatar": "https://upload.wikimedia.org/wikipedia/commons/9/90/Jordana_Brewster_at_PaleyFest_2013_2.jpg",\
            "birthday": "1980-04-26T00:00:00Z",\
            "career": "Актриса, модель",\
            "height": 170,\
            "birthPlace": "Панама-Сити, Флорида, США",\
            "spouse": "Эндрю Форм (с 2007 г.)"\
        }\
    ],\
    "directorToAdd": {\
        "name": "Роб Коин",\
        "avatar": "https://upload.wikimedia.org/wikipedia/commons/4/41/US_Navy_040618-N-6817C-090_Director_Rob_Cohen_visits_with_Commanding_Officer%2C_USS_Abraham_Lincoln_%28CVN_72%29%2C_Capt._Kendall_L._Card%2C_on_the_bridge_after_the_completion_of_filming%2C_the_upcoming_motion_picture_Stealth_%28cropped%29.JPG",\
        "birthday": "1959-03-12T00:00:00Z"\
    }\
}'
