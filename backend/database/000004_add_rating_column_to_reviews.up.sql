

ALTER TABLE reviews ADD rating AS (cleanliness + location + amenities + services) / 4 PERSISTED; -- computed average rating