
-- query := `INSERT INTO event_bookings (id, event_id, firstname, lastname, 
	-- email, total_price, quantity, promocode_used, payment_method, paystack_access_code, paystack_reference) OUTPUT INSERTED.booking_number 
	-- VALUES(@id, @event_id, @firstname, @lastname, @email, @total_price, @quantity, 
	-- @promocode_used, @payment_method, @access_code, @reference);`
-- event bookings schema
CREATE TABLE event_bookings (
    id VARCHAR(36) PRIMARY KEY,
    booking_number INT IDENTITY(1,1),
    event_id VARCHAR(36) NOT NULL,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    promocode_used VARCHAR(50),
    payment_status VARCHAR(20) DEFAULT 'pending' CHECK(payment_status IN ('pending', 'processing', 'paid', 'refunded')),
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME,
    quantity INT DEFAULT 1,
    payment_method VARCHAR(20),
    paystack_reference VARCHAR(256),
    paystack_access_code VARCHAR(256),
    FOREIGN KEY (event_id) REFERENCES events(id)
);
