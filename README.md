# Golang-gin-realworld-example-app
This codebase was created to demonstrate a fully fledged fullstack application built with Golang/Gin including CRUD operations, authentication, routing, pagination, and more.

## APIs
- Users and Authentication
    - `POST /user/login`: Existing user login
    - `POST /users`: Register a new user
    - `GET /user`: Get current user
    - `PUT /user`: Update current user
 
- Profiles
    - `GET /profiles/{username}`: Get a profile
    - `POST /profiles/{username}/follow`: Follow a user
    - `DELETE /profiles/{username}/follow`: Unfollow a user

- Articles
    - `GET /articles/feed`: Get recent articles from users you follow
    - `GET /articles`: Get recent articles globally
    - `POST /articles`: Create an article
    - `GET /articles/{slug}`: Get an article
    - `PUT /articles/{slug}`: Update an article
    - `DELETE /articles/{slug}`: Delete an article
 
- Comments
    - `GET /articles/{slug}/comments`: Get comments for an article
    - `POST /articles/{slug}/comments`: Create a comment for an article
    - `DELETE /articles/{slug}/comments/{id}`: Delete a comment for an article
 
- Favorites
    - `POST /articles/{slug}/favorite`: Favorite an article
    - `DELETE /articles/{slug}/favorite`: Unfavorite an article