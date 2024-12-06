PRAGMA foreign_keys=OFF;
-- BEGIN TRANSACTION;
CREATE TABLE users(
    id VARCHAR(255) PRIMARY KEY,
    firstname VARCHAR(100) DEFAULT "",
    lastname VARCHAR(100) DEFAULT "",
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(20) DEFAULT "user" CHECK(role IN ('user', 'admin', 'partner')),
    change_password_token VARCHAR(255) -- a jwt token with email as a claim and expiry time, this token is sent as email to be added to the new pasword form
);
CREATE TABLE hotels(
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(50) NOT NULL,
    state VARCHAR(50) NOT NULL,
    address VARCHAR(100),
    type VARCHAR(50) NOT NULL DEFAULT "guest house", --Hotel & suite, Apartment, Resort, Guest House, etc
    price_per_night REAL DEFAULT 0.0,
    amenities TEXT DEFAULT "no smoking, parking, swimming pool", --comma sepated amenities
    rating INTEGER DEFAULT 5,
    images TEXT DEFAULT "default_hotel.jpeg",  -- comma separated url endpoints, the first being the main one
    popular BOOLEAN DEFAULT 0
);
INSERT INTO hotels VALUES('1','Eko Hotel','Eko-isolo','lagos',NULL,'guest house',3300.3000000000001818,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',1);
INSERT INTO hotels VALUES('2','Maropo Hotel','Maropo-','Osun',NULL,'guest house',6000.300000000000182,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',0);
INSERT INTO hotels VALUES('3','Jigoo Hotel','Eko-isolo','lagos',NULL,'guest house',300.50000000000000001,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',1);
INSERT INTO hotels VALUES('4','Meko Hotel','Eko-isolo','Kano',NULL,'guest house',2200.099999999999909,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',1);
INSERT INTO hotels VALUES('5','Ekonm Hotel','Eko-isolo','lagos',NULL,'guest house',5300.3000000000001818,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',1);
INSERT INTO hotels VALUES('6','Aristosko Hotel','Eko-isolo','Abuja',NULL,'guest house',36300.300000000002909,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',1);
INSERT INTO hotels VALUES('7','Eko Hotel','Eko-isolo','Ogun',NULL,'guest house',330340.29999999998835,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',0);
INSERT INTO hotels VALUES('8','pemEko Hotel','Eko-isolo','lagos',NULL,'guest house',33002.300000000002911,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',0);
INSERT INTO hotels VALUES('9','Ekomdhd Hotel','Eko-isolo','lagos',NULL,'guest house',9001.2999999999992722,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',0);
INSERT INTO hotels VALUES('10','Elko Hotel','Eko-isolo','lagos',NULL,'guest house',23300.299999999999271,'no smoking, parking, swimming pool',5,'default_hotel.jpeg',1);
CREATE TABLE reviews (
    id VARCHAR(255) PRIMARY KEY,
    hotel_id VARCHAR(255),
    name VARCHAR(50), -- name of the person creating the review
    country VARCHAR(50), -- country of the person creating the review
    created_at DATE DEFAULT (DATE('now')),
    city VARCHAR(50) NOT NULL,
    cleanliness INTEGER DEFAULT 5, -- rate 1-10
    location INTEGER DEFAULT 5, -- rate 1-10
    amenities INTEGER DEFAULT 5, -- rate 1-10
    services INTEGER DEFAULT 5, -- rate 1-10
    one_word VARCHAR(25) DEFAULT 'outstanding', -- e.g., outstanding, poor, etc.
    FOREIGN KEY (hotel_id) REFERENCES hotels (id)
);
INSERT INTO reviews VALUES('1','1',NULL,NULL,'2024-12-06','lagos',5,5,5,5,'outstanding');
INSERT INTO reviews VALUES('2','1',NULL,NULL,'2024-12-06','Osun',5,5,5,5,'outstanding');
INSERT INTO reviews VALUES('3','1',NULL,NULL,'2024-12-06','lagos',5,5,5,5,'outstanding');
INSERT INTO reviews VALUES('4','1',NULL,NULL,'2024-12-06','kano',5,5,5,5,'outstanding');
INSERT INTO reviews VALUES('5','2',NULL,NULL,'2024-12-06','lagos',5,5,5,5,'outstanding');
INSERT INTO reviews VALUES('6','2',NULL,NULL,'2024-12-06','Abuja',5,5,5,5,'outstanding');
INSERT INTO reviews VALUES('7','3',NULL,NULL,'2024-12-06','Ogun',5,5,5,5,'outstanding');
INSERT INTO reviews VALUES('8','5',NULL,NULL,'2024-12-06','lagos',5,5,5,5,'outstanding');
INSERT INTO reviews VALUES('9','6',NULL,NULL,'2024-12-06','Eko-isolo',5,5,5,5,'outstanding');
INSERT INTO reviews VALUES('10','4',NULL,NULL,'2024-12-06','lagos',5,5,5,5,'outstanding');
CREATE TABLE hotel_bookings (
    id VARCHAR(255) PRIMARY KEY,
    hotel_id VARCHAR(255),
    created_at DATE DEFAULT (DATE('now')),
    effective_date DATE,
    rooms INTEGER DEFAULT 1,
    guests TEXT, -- in this format [[firstname,lastname,email,phone],[firstname,lastname,email,phone]]
    special_requests TEXT DEFAULT "",
    status VARCHAR(20) DEFAULT "pending" CHECK(status IN ('pending', 'processing', 'paid', 'used')),
    FOREIGN KEY (hotel_id) REFERENCES hotels (id)
);
CREATE TABLE blogs (
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255),
    author VARCHAR(50),
    display_image VARCHAR(50),
    created_at DATE DEFAULT (DATE('now')),
    updated_at DATE DEFAULT (DATE('now'))
);
INSERT INTO blogs VALUES('1','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('2','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('3','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('4','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('5','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('6','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('7','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('8','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('9','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('10','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('11','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('12','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('13','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
INSERT INTO blogs VALUES('14','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia',NULL,'default_blog.jpeg','2024-12-06','2024-12-06');
-- COMMIT;
