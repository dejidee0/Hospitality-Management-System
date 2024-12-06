-- populate hotels table


INSERT INTO hotels (id, name, city, state, price_per_night, popular)
VALUES
    ("1", "Eko Hotel", "Eko-isolo", "lagos", 3300.30, 1),
    ("2", "Maropo Hotel", "Maropo-", "Osun", 6000.30, 0),
    ("3", "Jigoo Hotel", "Eko-isolo", "lagos", 300.50, 1),
    ("4", "Meko Hotel", "Eko-isolo", "Kano", 2200.10, 1),
    ("5", "Ekonm Hotel", "Eko-isolo", "lagos", 5300.30, 1),
    ("6", "Aristosko Hotel", "Eko-isolo", "Abuja", 36300.30, 1),
    ("7", "Eko Hotel", "Eko-isolo", "Ogun", 330340.30, 0),
    ("8", "pemEko Hotel", "Eko-isolo", "lagos", 33002.30, 0),
    ("9", "Ekomdhd Hotel", "Eko-isolo", "lagos", 9001.30, 0),
    ("10", "Elko Hotel", "Eko-isolo", "lagos", 23300.30, 1);


INSERT INTO reviews (id, hotel_id, city)
VALUES
    ("1", "1", "lagos"),
    ("2", "1", "Osun"),
    ("3", "1", "lagos"),
    ("4", "1", "kano"),
    ("5", "2", "lagos"),
    ("6", "2", "Abuja"),
    ("7", "3", "Ogun"),
    ("8", "5", "lagos"),
    ("9", "6", "Eko-isolo"),
    ("10", "4", "lagos");