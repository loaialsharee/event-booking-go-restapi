# event-booking-go-restapi
An event booking REST API project powered by Golang.

# Project Breakdown

## API Routes
This part shows the REST API endpoints that are included in this Golang backend server:

|Method                |Route      |Description                          |Auth                         |
|--------|----------------|---------------|-----------------------------|
|`GET`| /events | Get a list of available events | 
|`GET`| /events/{id} |Get an available list by ID  |
|`POST`| /events | Create a new event | `Auth Required` 
|`PUT`| /events/{id} |Update an event | `Auth Required` `Admin access only`
|`DELETE`| /events/{id} |Delete an event|`Auth Required` `Admin access only`
|`POST`| /signup |Create a new user|
|`POST`| /login  |Authenticate user| `Initiates Auth Token`
|`POST`| /events/{id}/register |Register user for event| `Auth Required`
|`DELETE`| /events/{id}/register |Cancel event registration |`Auth Required`

## Visual Flow Chart 

```mermaid
graph LR

A((Auth)) --> B(GET /events)
A((Auth)) --> C(GET /events/id)
A((Auth)) --> D(POST /events)
A((Auth)) -- Admin Access --> E(PUT /events/id)
A((Auth)) -- Admin Access --> F(DELETE /events/id)
A((Auth)) --> I(POST /events/id/register)
A((Auth)) --> J(DELETE /events/id/register)

