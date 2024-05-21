# User Management System with FastHTTP

This repository implements a user management system using the FastHTTP framework. Below is a detailed explanation of the available endpoints, including user login management and user management routes.

## Endpoints

### User Login Management Routes

#### POST `/api/v1/login`

This endpoint handles user login.

**Handler:** `s.userSigninHandler.LoginIn`

**Description:** This route processes user login requests, verifying credentials and returning a token or session information upon successful authentication.

### User Management Routes

#### GET `/api/v1/users`

This endpoint retrieves a list of users.

**Middleware:**

- `logging.ErrorLoggingMiddleware`: Logs errors that occur during the request.
- `authentication.AuthorizationMiddleware("Administrator")`: Authorizes only users with the "Administrator" role.

**Handler:** `s.userHandler.List`

**Description:** This route returns a list of all users in the system. Only administrators can access this route.

#### POST `/api/v1/users`

This endpoint creates a new user.

**Handler:** `s.userHandler.Create`

**Description:** This route handles the creation of a new user. The request should include user details such as name, email, and password.

#### GET `/api/v1/users/{id}`

This endpoint retrieves a user by ID.

**Handler:** `s.userHandler.Find`

**Description:** This route returns the details of a specific user identified by their ID.

#### PUT `/api/v1/users/{id}`

This endpoint updates a user's information.

**Handler:** `s.userHandler.Update`

**Description:** This route allows for updating user details such as name, email, and role. The user is identified by their ID.

#### DELETE `/api/v1/users/{id}`

This endpoint deletes a user by ID.

**Handler:** `s.userHandler.Remove`

**Description:** This route handles the deletion of a specific user identified by their ID.

#### POST `/api/v1/users/search`

This endpoint searches for users based on specific criteria.

**Handler:** `s.userHandler.Search`

**Description:** This route allows for searching users by various criteria such as name, email, or role. The search criteria should be included in the request body.

## Middleware

### Error Logging Middleware

**Middleware:** `logging.ErrorLoggingMiddleware`

**Description:** This middleware logs any errors that occur during the handling of a request. It is used to track and diagnose issues in the application.

### Authorization Middleware

**Middleware:** `authentication.AuthorizationMiddleware("Administrator")`

**Description:** This middleware ensures that only users with the "Administrator" role can access certain routes. It checks the user's role and either allows the request to proceed or denies access.

## How to Run

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   cd <repository-directory>
