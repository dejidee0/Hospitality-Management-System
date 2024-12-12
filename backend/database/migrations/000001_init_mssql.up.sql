-- user schema
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    fullname VARCHAR(50),
    username VARCHAR(50),
    email VARCHAR(100) NOT NULL UNIQUE,
    phone_number VARCHAR(15),
    gender VARCHAR(10) DEFAULT '' CHECK (gender IN ('male', 'female', '')),
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME,
    role VARCHAR(20) DEFAULT 'user' CHECK (role IN ('user', 'agent', 'admin')),
    apply_to_be_agent BIT DEFAULT 0,
    approved_as_agent BIT DEFAULT 0,
    change_password_token VARCHAR(255)
);
 
-- hotel schema 
CREATE TABLE hotels(
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    property_type VARCHAR(50) NOT NULL, --Hotel & suite, Apartment, Resort, Guest House, etc not (summary in figma!)or say summary is replaced with type
    description VARCHAR(2048) NOT NULL,
    address VARCHAR(512) NOT NULL,
    city VARCHAR(50) NOT NULL,
    state VARCHAR(50) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    country VARCHAR(50) NOT NULL,
    part_of_group BIT NOT NULL,
    part_of_chain BIT NOT NULL,
    rating INT,
    images VARCHAR(2048), -- comma separated url endpoints, the first being the main one
    popular BIT DEFAULT 0,
    published BIT DEFAULT 0,
    agent_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (agent_id) REFERENCES users(id)
);

-- house rules schema
CREATE TABLE house_rules (
    id VARCHAR(36) PRIMARY KEY,
    hotel_id VARCHAR(36) NOT NULL,
    check_in VARCHAR(50) NOT NULL, -- e.g 'from 14:00 to 00:00'
    check_out VARCHAR(50) NOT NULL, -- e.g from 01:00 to 12:00
    allow_children BIT NOT NULL,
    allow_pets VARCHAR(10) DEFAULT 'no' CHECK(allow_pets IN ('yes', 'no', 'request')),
    cancel_booking VARCHAR(100),  -- when can guests cancel booking for free, e.g before 18:00 on the day of check-in
    cancel_booking_charge VARCHAR(100),
    currency VARCHAR(50),
    payment_method VARCHAR(50), -- comma separated 'credit card,cash,transfer'
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);

-- rooms schema
CREATE TABLE rooms (
    id VARCHAR(36) PRIMARY KEY,
    hotel_id VARCHAR(36) NOT NULL,
    images VARCHAR(2048), -- comma separated url, the first being the main one
    name VARCHAR(255) NOT NULL, -- name of room e.g 'Executive One-Bedroom Suite'
    type VARCHAR(255) NOT NULL, -- type of room e.g 'Double Room'
    total_number INT NOT NULL, --number of rooms of this type
    capacity INT NOT NULL, --number of guests that can stay in this room
    room_size VARCHAR(50) NOT NULL, --how big is the room e.g '30 square meter'
    allow_smoking BIT NOT NULL,
    single_bed INT DEFAULT 0,
    double_bed INT DEFAULT 0,
    king_sized INT DEFAULT 0,
    super_king_sized INT DEFAULT 0,
    price_per_night DECIMAL(10,2) DEFAULT 0.0,
    amenities VARCHAR(512), ----comma sepated amenities e.g "parking, swimming pool, restaurant, fitness centre, highspeed internet, dry cleaning"
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);


-- hotel bookings schema
CREATE TABLE hotel_bookings (
    id VARCHAR(36) PRIMARY KEY,
    booking_number INT IDENTITY(1,1),
    room_id VARCHAR(36) NOT NULL,
    guest_names VARCHAR(512) NOT NULL, --comma separated names in case of multiple guests
    phone_number VARCHAR(15) NOT NULL,
    email VARCHAR(100) NOT NULL,
    check_in DATETIME NOT NULL,
    check_out DATETIME NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    payment_status VARCHAR(20) DEFAULT 'pending' CHECK(payment_status IN ('pending', 'processing', 'paid', 'refunded')),
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME,
    rooms INT DEFAULT 1,
    special_requests VARCHAR(1024),
    payment_method VARCHAR(20),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

-- reviews schema
CREATE TABLE reviews (
    id VARCHAR(36) PRIMARY KEY,
    hotel_id VARCHAR(36),
    name VARCHAR(50), -- name of the person creating the review
    country VARCHAR(20), -- country of the person creating the review
    review_body VARCHAR(1024), -- the review itself
    created_at DATETIME DEFAULT GETDATE(),
    city VARCHAR(50) NOT NULL,
    cleanliness INT, -- rate 1-10
    location INT, -- rate 1-10
    amenities INT, -- rate 1-10
    services INT, -- rate 1-10
    rating AS ((cleanliness + location + amenities + services) / 4), -- computed average rating
    one_word VARCHAR(25), -- e.g., outstanding, poor, etc.
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
);

-- blogs schema
CREATE TABLE blogs (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(MAX),
    author VARCHAR(50),
    display_image VARCHAR(2048),
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME
);
