# Hospitality-Management-System
## The Backend
Hotel Management System with Flight Booking and Event Ticketing

## Start the server
to start the server, you either run
```
go run main.go
```
or use the make command
``` make server ```

## Endpoints Documentation
NOTE: prefix every endpoint with `/api/v1

<!-- this endpoint checks if the server is running fine -->
1. GET /health


2. POST /auth/signup

endpoint creates a user


    Payload:

        - email     [required]
        - password  [required]

3. POST /auth/login

    Payload:

        - email     [required]
        - password  [required]
    Return:
        {message, token}

******** IGNORE THIS************************************

4. GET /auth/profile**********************************

*********** IT'S JUST FOR TESTING PURPOSES ************


5. GET /auth/reset-password

    this endpoint sends a reset token to the supplied email, token expires in 1hr

    Query Parameter:

        - email


6. POST /auth/change-password

    endpoint changes password to the newly supplied password, provided the token is valid

    Payload:

        - token
        - password

7. GET /hotels/index

    this gets the data required for the home page

    returns:

        data: {
			popular_hotels: list of popular hotels,
			trending_destinations: list of trending destinations,
			property_types:  list of property_types,
			blogs: list of recent travel articles
		}

8. GET /hotels/search

    Query Parameters:

        - state
    returns hotels from the supplied state with the neccessary details for the search page

9. GET /hotels/<hotel_id>

    returns details about the hotel with all the available rooms and their details

10. POST /hotels/booking

    Payload:

        room-id:            string   
	    guest-names:        string      --- comma separated names in case of multiple guests
        phone-numbers:      string      --- comma separated phones in case of multiple guests
        emails:             string      --- comma separated emails in case of multiple guests
        special-requests:   string
        payment-method:     string
        quantity:           int         --- number of rooms
        promo-code:         string      --- empty string '' if not applicable
        check-in            datetime    --- in this format "2006-01-02T15:00:00Z"
        check-out           datetime    --- in this format "yyyy-mm-ddThr:mm:ssZ"
        number-of-night     int   
    
    Response:
        {
            authorization-url: [string],
            booking-id: [string],
            booking-number: [int],
            email: [string],
            paystack-access-code: [string],
            total-amount: [float],
            tx-reference: [string]
        }

11. GET /hotels/booking/verify

        Query Parameters:

            -   reference
            -   booking_id
        Response body:
            {
                "status": "success"
                "message": [string],
                "data": {
                    "guest-name": [string],
                    "booking-number": [int],
                    "email": [string],
                    "check-in": [datetime],
                    "check-out": [datetime],
                    "rooms": [int],
                    "total-price": [float],
                    "room-name": [string],
                    "amenities": [string],
                    "price-per-night": [string],
                    "image": [string],
                    "hotel-name": [string],
                    "tax-rate":[string],
                    "discount": [float],
                }
            }

12. GET /events/index

        returns:
        data: {
            popular_events:  list of popular events,
			online_events:   list of events in the online format,
			music_events:   list of events in the music category,
			business_events: list of events in the business category,
		}

13. GET /events/search

    Query Parameter:

        -   state
    returns events in the supplied state with the neccessary details for the search page

14. GET /events/<event_id>

    returns details about the event with the provided id


15. POST /events/booking

    Payload:

        {
            event-id:           string   
            first-name:         string      
            last-name:          string      
            email:              string      
            payment-method:     string
            quantity:           int         --- number of tickets
            promo-code:         string      --- empty string '' [optional]
            total-amount        int   
        }  

    Response:

        {
            authorization-url: [string],
            event-booking-id: [string],
            booking-number: [int],
            email: [string],
            paystack-access-code: [string],
            total-amount: [float],
            tx-reference: [string]
        }


16. GET /events/booking/verify

        Query Parameters
        -   reference

        Response body:
        "status":  "success",
		"message": "payment is successful!",
		"data": {
                    "status": "success"
                    "message": [string],
                    "data": {
                        "first-name": [string],
                        "booking-number": [int],
                        "email": [string],                    
                        "quantity": [int],
                        "total-amount": [float],
                        "image": [string],
                        "event-name": [string],
                        "event-date":[string],
                        "venue": [float],
                        "image": [string]
                    }
            },



for verify booking for events, i may restructure verify booking hotel to accomodate all verify