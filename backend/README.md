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

1. GET /health

2. POST /auth/signup
    Payload:
        - email     [required]
        - password  [required]

3. POST /auth/login
    Payload:
        - email     [required]
        - password  [required]
    Return:
        {message, token}

<!-- IGNORE THIS -->
4. GET /auth/profile
<!-- IT'S JUST FOR TESTING PURPOSES -->

5. GET /auth/reset-password
    Query Parameter
        - email


6. POST /auth/change-password
    Payload:
        - token
        - password
