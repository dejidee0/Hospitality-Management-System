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

<!-- this endpoint checks if the server is running fine -->
1. GET /v1/health

endpoint creates a user, expects [email: string, password: string]
2. POST /v1/auth/signup
    Payload:
        - email     [required]
        - password  [required]

3. POST /v1/auth/login
    Payload:
        - email     [required]
        - password  [required]
    Return:
        {message, token}

<!-- IGNORE THIS -->
4. GET /v1/auth/profile
<!-- IT'S JUST FOR TESTING PURPOSES -->

<!-- this endpoint sends a reset token to the supplied email, token expires in 1hr -->
5. GET /v1/auth/reset-password
    Query Parameter
        - email

<!-- endpoint changes password to the newly supplied password, provided the token is valid -->
6. POST /v1/auth/change-password
    Payload:
        - token
        - password

<!-- this gets the data required for the home page -->
7. GET /v1/hotels/index
    returns:
        data: {
			popular_hotels: list of popular hotels,
			trending_destinations: list of trending destinations,
			property_types:  list of property_types,
			blogs: list of recent travel articles
		}

8. GET /v1/hotels/search?state=
    returns hotels from the supplied state with the neccessary details for the search page

9. GET /v1/hotels/<hotel_id>
    returns details about the hotel with all the available rooms and their details

10. POST /v1/hotels/booking
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
        
