
-- insert sample data into rooms
INSERT INTO rooms (
    id, hotel_id, images, name, type, total_number, capacity, room_size, 
    allow_smoking, single_bed, double_bed, king_sized, super_king_sized, 
    price_per_night, amenities
) 
VALUES
-- Room 1
('room1-uuid-1234', '1bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room1.jpg,https://example.com/room1a.jpg', 
 'Executive Suite', 'Suite', 10, 3, '45 square meter', 
 0, 0, 1, 1, 0, 250.00, 
 'wifi, tv, minibar, bathtub, air conditioning'),
-- Room 2
('room2-uuid-1234', '1bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room2.jpg,https://example.com/room2a.jpg', 
 'Deluxe Room', 'Double Room', 15, 2, '30 square meter', 
 0, 0, 1, 0, 0, 180.00, 
 'wifi, tv, balcony, air conditioning'),
-- Room 3
('room3-uuid-1234', '2bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room3.jpg,https://example.com/room3a.jpg', 
 'Standard Room', 'Single Room', 20, 1, '25 square meter', 
 0, 1, 0, 0, 0, 100.00, 
 'wifi, tv, desk'),
-- Room 4
('room4-uuid-1234', '3bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room4.jpg,https://example.com/room4a.jpg', 
 'Family Suite', 'Suite', 5, 5, '70 square meter', 
 0, 2, 1, 1, 0, 350.00, 
 'wifi, tv, kitchen, sofa, balcony'),
-- Room 5
('room5-uuid-1234', '4bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room5.jpg,https://example.com/room5a.jpg', 
 'Queen Room', 'Double Room', 10, 2, '35 square meter', 
 1, 0, 1, 0, 0, 200.00, 
 'wifi, tv, bathtub, air conditioning'),
-- Room 6
('room6-uuid-1234', '5bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room6.jpg,https://example.com/room6a.jpg', 
 'Honeymoon Suite', 'Suite', 3, 2, '60 square meter', 
 0, 0, 0, 1, 1, 400.00, 
 'wifi, jacuzzi, tv, minibar, private pool'),
-- Room 7
('room7-uuid-1234', '6bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room7.jpg,https://example.com/room7a.jpg', 
 'Business Class Room', 'Double Room', 8, 2, '40 square meter', 
 0, 0, 1, 0, 0, 220.00, 
 'wifi, desk, printer, air conditioning'),
-- Room 8
('room8-uuid-1234', '7bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room8.jpg,https://example.com/room8a.jpg', 
 'Penthouse Suite', 'Suite', 2, 4, '120 square meter', 
 0, 0, 0, 2, 0, 1000.00, 
 'wifi, tv, kitchen, private terrace, pool'),
-- Room 9
('room9-uuid-1234', '8bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room9.jpg,https://example.com/room9a.jpg', 
 'Economy Room', 'Single Room', 25, 1, '20 square meter', 
 0, 1, 0, 0, 0, 80.00, 
 'wifi, tv'),
-- Room 10
('room10-uuid-1234', '9bcd1234-5678-90ab-cdef-1234567890ab', 
 'https://example.com/room10.jpg,https://example.com/room10a.jpg', 
 'Presidential Suite', 'Suite', 1, 6, '150 square meter', 
 0, 0, 0, 2, 2, 2000.00, 
 'wifi, private pool, tv, kitchen, gym');


-- isnert sample data into reviews
INSERT INTO reviews (
    id, hotel_id, name, country, review_body, created_at, city, cleanliness, location, amenities, services, one_word
)
VALUES
-- Review 1 for Hotel 1
('review1-uuid-1234', '1bcd1234-5678-90ab-cdef-1234567890ab', 
 'John Doe', 'USA', 'The hotel was excellent with clean rooms and amazing service.', 
 '2024-12-01 12:00:00', 'New York', 9, 10, 9, 10, 'outstanding'),
-- Review 2 for Hotel 1
('review2-uuid-1234', '1bcd1234-5678-90ab-cdef-1234567890ab', 
 'Jane Smith', 'Canada', 'Great location and comfortable stay, but the amenities could be improved.', 
 '2024-12-02 14:30:00', 'New York', 8, 9, 7, 8, 'great'),
-- Review 1 for Hotel 2
('review3-uuid-1234', '2bcd1234-5678-90ab-cdef-1234567890ab', 
 'Ali Khan', 'Pakistan', 'Affordable and clean. Highly recommended for budget travelers.', 
 '2024-11-28 10:15:00', 'Los Angeles', 8, 8, 7, 7, 'satisfactory'),
-- Review 2 for Hotel 2
('review4-uuid-1234', '2bcd1234-5678-90ab-cdef-1234567890ab', 
 'Maria Gonzalez', 'Mexico', 'Good value for money, but the location is a bit remote.', 
 '2024-11-30 17:45:00', 'Los Angeles', 7, 6, 8, 7, 'decent'),
-- Review 1 for Hotel 3
('review5-uuid-1234', '3bcd1234-5678-90ab-cdef-1234567890ab', 
 'Emma Brown', 'UK', 'Beautiful beachfront location with excellent amenities.', 
 '2024-12-03 09:00:00', 'Miami', 10, 10, 9, 10, 'amazing'),
-- Review 2 for Hotel 3
('review6-uuid-1234', '3bcd1234-5678-90ab-cdef-1234567890ab', 
 'Noah White', 'USA', 'A great getaway spot. The staff were very helpful.', 
 '2024-12-04 13:20:00', 'Miami', 9, 9, 8, 9, 'great'),
-- Review 1 for Hotel 4
('review7-uuid-1234', '4bcd1234-5678-90ab-cdef-1234567890ab', 
 'Liam Davis', 'Australia', 'Cozy and comfortable, but not many facilities available.', 
 '2024-11-29 11:30:00', 'Chicago', 8, 8, 6, 7, 'good'),
-- Review 2 for Hotel 4
('review8-uuid-1234', '4bcd1234-5678-90ab-cdef-1234567890ab', 
 'Sophia Wilson', 'UK', 'Affordable guest house with friendly staff.', 
 '2024-11-30 15:00:00', 'Chicago', 7, 7, 7, 8, 'nice'),
-- Review 1 for Hotel 5
('review9-uuid-1234', '5bcd1234-5678-90ab-cdef-1234567890ab', 
 'Olivia Martinez', 'Spain', 'Quiet retreat with beautiful views.', 
 '2024-12-01 18:45:00', 'Denver', 9, 9, 8, 9, 'peaceful'),
-- Review 2 for Hotel 5
('review10-uuid-1234', '5bcd1234-5678-90ab-cdef-1234567890ab', 
 'Ethan Taylor', 'USA', 'Great place to relax, but some amenities were not functioning.', 
 '2024-12-02 08:30:00', 'Denver', 8, 8, 7, 8, 'relaxing');


-- insert into house rules
INSERT INTO house_rules (
    id, hotel_id, check_in, check_out, allow_children, allow_pets, cancel_booking, cancel_booking_charge, currency, payment_method
)
VALUES
-- House Rules for Hotel 1
('rule1-uuid-1234', '1bcd1234-5678-90ab-cdef-1234567890ab', 
 'from 14:00 to 00:00', 'from 01:00 to 12:00', 
 1, 'no', 'before 18:00 on the day of check-in', '50% of the first night', 
 'USD', 'credit card,cash'),
-- House Rules for Hotel 2
('rule2-uuid-1234', '2bcd1234-5678-90ab-cdef-1234567890ab', 
 'from 15:00 to 23:00', 'from 02:00 to 11:00', 
 1, 'yes', '24 hours before check-in', 'one night stay charge', 
 'USD', 'credit card,transfer'),
-- House Rules for Hotel 3
('rule3-uuid-1234', '3bcd1234-5678-90ab-cdef-1234567890ab', 
 'from 13:00 to 22:00', 'from 03:00 to 10:00', 
 1, 'request', '48 hours before check-in', 'full booking amount', 
 'EUR', 'credit card,cash'),
-- House Rules for Hotel 4
('rule4-uuid-1234', '4bcd1234-5678-90ab-cdef-1234567890ab', 
 'from 12:00 to 21:00', 'from 04:00 to 11:30', 
 1, 'no', 'by 12:00 noon on check-in day', '50% of total booking amount', 
 'GBP', 'cash,credit card'),
-- House Rules for Hotel 5
('rule5-uuid-1234', '5bcd1234-5678-90ab-cdef-1234567890ab', 
 'from 16:00 to 00:00', 'from 02:00 to 12:00', 
 1, 'request', '7 days before check-in', 'non-refundable', 
 'USD', 'credit card,transfer');



-- insert oito blogs
INSERT INTO blogs (
    id, title, body, author, display_image, created_at, updated_at
)
VALUES
-- Blog 1
('blog1-uuid-1234', 'The Future of Travel: Exploring New Destinations', 
 'In this blog, we delve into the emerging travel trends and how new destinations are shaping the future of travel. From off-the-beaten-path locales to luxury escapes, the possibilities are endless.',
 'John Doe', 'https://example.com/images/future_of_travel.jpg', 
 '2024-12-01 10:00:00', '2024-12-01 10:00:00'),
-- Blog 2
('blog2-uuid-1234', 'Top 10 Budget-Friendly Destinations for 2025', 
 'Looking to travel without breaking the bank? Check out our top 10 destinations for budget-friendly getaways in 2025. These spots offer incredible experiences without the hefty price tag.',
 'Jane Smith', 'https://example.com/images/budget_destinations.jpg', 
 '2024-12-02 12:30:00', '2024-12-02 12:30:00'),
-- Blog 3
('blog3-uuid-1234', 'Ultimate Guide to Eco-Tourism', 
 'Eco-tourism is on the rise, and this guide will show you how to travel responsibly while preserving the environment. From sustainable accommodations to carbon offset programs, we cover it all.',
 'Ali Khan', 'https://example.com/images/eco_tourism_guide.jpg', 
 '2024-12-03 08:45:00', '2024-12-03 08:45:00'),
-- Blog 4
('blog4-uuid-1234', 'How to Plan the Perfect Family Vacation', 
 'Planning a family vacation can be stressful, but with the right tips and tricks, it can be an unforgettable experience. Discover the best family-friendly destinations and activities for all ages.',
 'Maria Gonzalez', 'https://example.com/images/family_vacation.jpg', 
 '2024-12-04 14:00:00', '2024-12-04 14:00:00'),
-- Blog 5
('blog5-uuid-1234', 'The Rise of Wellness Travel', 
 'Wellness travel has gained popularity in recent years, and its only expected to grow. In this blog, we explore wellness retreats, spas, and health-conscious travel destinations.',
 'Sophia Wilson', 'https://example.com/images/wellness_travel.jpg', 
 '2024-12-05 09:15:00', '2024-12-05 09:15:00');
