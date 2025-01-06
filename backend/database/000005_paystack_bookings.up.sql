-- paystack bookings
CREATE TABLE paystack_bookings(
    id VARCHAR(36) PRIMARY KEY,
    booking_number INT,
    booking_id VARCHAR(36),
    amount DECIMAL(10,2),
    type VARCHAR(10) CHECK(type IN ('hotel', 'flight','event','car')),
    reference VARCHAR(256),
    access_code VARCHAR(256),
    payment_status VARCHAR(100)
);
