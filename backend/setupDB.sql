-- Enable foreign key constraints
PRAGMA foreign_keys = ON;

-- CREATE users table
CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(255) PRIMARY KEY,
    firstname VARCHAR(100) DEFAULT "",
    lastname VARCHAR(100) DEFAULT "",
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(20) DEFAULT "user" CHECK(role IN ('user', 'admin', 'partner')),
    change_password_token VARCHAR(255) -- a jwt token with email as a claim and expiry time, this token is sent as email to be added to the new pasword form
);


-- CREATE hotels table
CREATE TABLE IF NOT EXISTS hotels(
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



-- CREATE reviews table
CREATE TABLE IF NOT EXISTS reviews (
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


-- CREATE hotel_bookings table
CREATE TABLE IF NOT EXISTS hotel_bookings (
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


-- CREATE blogs table
CREATE TABLE IF NOT EXISTS blogs (
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255),
    author VARCHAR(50),
    display_image VARCHAR(50),
    created_at DATE DEFAULT (DATE('now')),
    updated_at DATE DEFAULT (DATE('now'))
);


