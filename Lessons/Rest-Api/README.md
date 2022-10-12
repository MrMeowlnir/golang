# Rest-Api tutorial project

# user-service

# Rest Api

## 1. Starting with the Null Style
## 2. Client-Server
## 3. Stateless
## 4. Cache
## 5. Uniform Interface
## 6. Layered System
## 7. Code-On-Demand
## 8. Style Derivation Summary

#endpoints
### GET /users -- list of users -- 200(OK), 404, 500
### GET /users/:id -- user by id -- 200, 404, 500
### POST /users/:id -- create user -- 204(NO CONTENT), 4xx(ANY ERRORS)   server shutdown etc.
### PUT /users/:id -- fully update user -- 204/200  New user
### PATCH /users/:id -- partially update user -- 204/200 New data in user
### DELETE /users/:id -- delete user by id -- 204, 404, 400 
