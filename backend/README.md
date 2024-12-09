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
    returns hotels that meet this state criteria
