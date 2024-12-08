PRAGMA foreign_keys=OFF;

-- users schema
CREATE TABLE users(
    id VARCHAR(36) PRIMARY KEY,
    fullname VARCHAR(50) DEFAULT "",
    username VARCHAR(50) DEFAULT "",
    email VARCHAR(100) NOT NULL UNIQUE,
    phone_number VARCHAR(15) DEFAULT "",
    gender VARCHAR(10) DEFAULT "" CHECK(gender IN ('male', 'female', '')),
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(20) DEFAULT "user" CHECK(role IN ('user', 'agent', 'admin')),
    apply_to_be_agent BOOLEAN DEFAULT 0,
    approved_as_agent BOOLEAN DEFAULT 0,
    change_password_token VARCHAR(255) -- a jwt token with email as a claim and expiry time, this token is sent as email to be added to the new pasword form
);

-- hotel schema
CREATE TABLE hotels(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    property_type VARCHAR(50) NOT NULL DEFAULT "guest house", --Hotel & suite, Apartment, Resort, Guest House, etc not (summary in figma!)or say summary is replaced with type
    description TEXT NOT NULL,
    address TEXT NOT NULL,
    city VARCHAR(50) NOT NULL,
    state VARCHAR(50) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    country VARCHAR(50) NOT NULL,
    part_of_group BOOLEAN NOT NULL,
    part_of_chain BOOLEAN NOT NULL,
    rating INTEGER DEFAULT 5,
    images TEXT, -- comma separated url endpoints, the first being the main one
    popular BOOLEAN DEFAULT 0,
    published BOOLEAN DEFAULT 0,
    agent_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (agent_id) REFERENCES users(id)
);
 
INSERT INTO hotels (id, name, description, address, city, state, postal_code, country, part_of_group, part_of_chain, popular, agent_id) 
VALUES 
('1','Eko Hotel','hdghdgdyd hdhdhdhdf','23, gddgdend gdttdd', 'Eko-isolo','lagos','1234','nigeria',0,1,1,'23'),
('2','Eko Hotel','hdghdgdyd hdhdhdhdf','23, gddgdend gdttdd', 'Eko-isolo','lagos','1234','nigeria',0,1,1,'23'),
('3','Eko Hotel','hdghdgdyd hdhdhdhdf','23, gddgdend gdttdd', 'Eko-isolo','lagos','1234','nigeria',0,1,1,'23'),
('4','Eko Hotel','hdghdgdyd hdhdhdhdf','23, gddgdend gdttdd', 'Eko-isolo','lagos','1234','nigeria',0,1,1,'23'),
('5','Eko Hotel','hdghdgdyd hdhdhdhdf','23, gddgdend gdttdd', 'Eko-isolo','lagos','1234','nigeria',0,1,1,'23'),
('6','Eko Hotel','hdghdgdyd hdhdhdhdf','23, gddgdend gdttdd', 'Eko-isolo','lagos','1234','nigeria',0,1,1,'23'),
('7','Eko Hotel','hdghdgdyd hdhdhdhdf','23, gddgdend gdttdd', 'Eko-isolo','lagos','1234','nigeria',0,1,1,'23'),
('8','Eko Hotel','hdghdgdyd hdhdhdhdf','23, gddgdend gdttdd', 'Eko-isolo','lagos','1234','nigeria',0,1,1,'23');



-- house rules schema
CREATE TABLE house_rules (
    id VARCHAR(36) PRIMARY KEY,
    hotel_id VARCHAR(36) NOT NULL,
    check_in VARCHAR(50) NOT NULL, -- e.g 'from 14:00 to 00:00'
    check_out VARCHAR(50) NOT NULL, -- e.g from 01:00 to 12:00
    allow_children BOOLEAN NOT NULL,
    allow_pets VARCHAR(10) DEFAULT "no" CHECK(allow_pets IN ('yes', 'no', 'request')),
    cancel_booking VARCHAR(100) NOT NULL,  -- when can guests cancel booking for free, e.g before 18:00 on the day of check-in
    cancel_booking_charge VARCHAR(100) NOT NULL,
    currency VARCHAR(50) NOT NULL,
    payment_method VARCHAR(50) NOT NULL, -- comma separated 'credit card,cash,transfer'
    FOREIGN KEY (hotel_id) REFERENCES hotels (id)
);

-- rooms schema
CREATE TABLE rooms (
    id VARCHAR(36) PRIMARY KEY,
    hotel_id VARCHAR(36) NOT NULL,
    images TEXT NOT NULL, -- comma separated url, the first being the main one
    name VARCHAR(255) NOT NULL, -- name of room e.g 'Executive One-Bedroom Suite'
    type VARCHAR(255) NOT NULL, -- type of room e.g 'Double Room'
    total_number INTEGER NOT NULL, --number of rooms of this type
    capacity INTEGER NOT NULL, --number of guests that can stay in this room
    room_size VARCHAR(50) NOT NULL, --how big is the room e.g '30 square meter'
    allow_smoking BOOLEAN NOT NULL,
    single_bed INTEGER DEFAULT 0,
    double_bed INTEGER DEFAULT 0,
    king_sized INTEGER DEFAULT 0,
    super_king_sized INTEGER DEFAULT 0,
    price_per_night REAL DEFAULT 0.0,
    amenities TEXT, ----comma sepated amenities e.g "parking, swimming pool, restaurant, fitness centre, highspeed internet, dry cleaning"
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);

-- hotel bookings schema
CREATE TABLE hotel_bookings (
    id VARCHAR(36), -- this will be changed in sql server since i cant have two primary keys
    booking_number INTEGER PRIMARY KEY AUTOINCREMENT,
    room_id VARCHAR(36) NOT NULL,
    guest_names VARCHAR(100) NOT NULL, --comma separated names in case of multiple guests
    phone_number VARCHAR(15) NOT NULL,
    email VARCHAR(100) NOT NULL,
    check_in DATETIME NOT NULL,
    check_out DATETIME NOT NULL,
    price REAL NOT NULL,
    payment_status VARCHAR(20) DEFAULT 'pending' CHECK(payment_status IN ('pending', 'processing', 'paid', 'refunded')),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    rooms INTEGER DEFAULT 1,
    special_requests TEXT DEFAULT "",
    payment_method VARCHAR(20) NOT NULL,
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

-- reviews schema
CREATE TABLE reviews (
    id VARCHAR(36) PRIMARY KEY,
    hotel_id VARCHAR(36),
    name VARCHAR(20), -- name of the person creating the review
    country VARCHAR(20), -- country of the person creating the review
    review_body TEXT, -- the review itself
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    city VARCHAR(50) NOT NULL,
    cleanliness INTEGER DEFAULT 5, -- rate 1-10
    location INTEGER DEFAULT 5, -- rate 1-10
    amenities INTEGER DEFAULT 5, -- rate 1-10
    services INTEGER DEFAULT 5, -- rate 1-10
    one_word VARCHAR(25) DEFAULT 'outstanding', -- e.g., outstanding, poor, etc.
    FOREIGN KEY (hotel_id) REFERENCES hotels (id)
);
INSERT INTO reviews (id, hotel_id, city) VALUES('1','1','lagos');
INSERT INTO reviews (id, hotel_id, city) VALUES('2','1','Osun');
INSERT INTO reviews (id, hotel_id, city) VALUES('3','1','lagos');
INSERT INTO reviews (id, hotel_id, city) VALUES('4','1','kano');
INSERT INTO reviews (id, hotel_id, city) VALUES('5','2','lagos');
INSERT INTO reviews (id, hotel_id, city) VALUES('6','2','Abuja');
INSERT INTO reviews (id, hotel_id, city) VALUES('7','3','Ogun');
INSERT INTO reviews (id, hotel_id, city) VALUES('8','5','lagos');
INSERT INTO reviews (id, hotel_id, city) VALUES('9','6','Eko-isolo');
INSERT INTO reviews (id, hotel_id, city) VALUES('10','4','lagos');

-- blogs schema
CREATE TABLE blogs (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body TEXT DEFAULT '',
    author VARCHAR(50) DEFAULT '',
    display_image VARCHAR(50) DEFAULT '',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- insert dummy data to table for testing purposes
INSERT INTO blogs (id, title, display_image) VALUES('1','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('2','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('3','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('4','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('5','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('6','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('7','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('8','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('9','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('10','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('11','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('12','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('13','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
INSERT INTO blogs (id, title, display_image) VALUES('14','An Echanting Spanish Summer Await: Discover the Many Charm of Valencia','default_blog.jpeg');
