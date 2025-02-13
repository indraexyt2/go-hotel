                                     # User Management Service Documentation

## Overview
This document outlines the API endpoints for the User Management Service (UMS). The service provides user authentication, registration, and profile management functionalities.

## Base URL
All API endpoints are prefixed with`/api/ums/v1`

## Authentication
Most endpoints require authentication using JWT tokens. Include the token in the Authorization header:
```
Authorization: <token>
```

## Endpoints

### Register
Creates a new user account.

**POST** `/api/ums/v1/register`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "username": "indra_123",
  "password": "indra_password",
  "email": "indra@example.com",
  "name": "Indra Wijaya",
  "phone": "+6281234567890",
  "address": "Jl. Merdeka No. 45",
  "role": "guest"
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "photo_profile": "path_url",
    "username": "indra_123",
    "email": "indra@example.com",
    "name": "Indra Wansyah Baru",
    "phone": "+6281234567890",
    "address": "Jl. Merdeka No. 45",
    "role": "guest",
    "is_verified": "false"
  }
}
```

### Email Verification
Verifies user email using a token.

**PUT** `/api/user/v1/email-verification/:token`

**Response:**
```json
{
  "message": "success"
}
```

### Resend Email Verification
Resends verification email to user.

**GET** `/api/user/v1/email-verification`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "email": "xxxx@example.com"
}
```

**Response:**
```json
{
  "message": "success"
}
```

### Login
Authenticates user and returns access tokens.

**POST** `/api/user/v1/login`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "username": "indra_123",
  "password": "indra_password"
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "user_id": 3,
    "full_name": "Indra Wansyah",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Refresh Token
Generates new access token using refresh token.

**PUT** `/api/user/v1/refresh-token`

**Headers:**
```
Authorization: <refresh_token>
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Get User
Retrieves authenticated user's profile.

**GET** `/api/ums/v1/user`

**Headers:**
```
Authorization: <token>
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "username": "indra_123",
    "photo_profile": "path_url",
    "email": "indra@example.com",
    "name": "Indra Wansyah",
    "phone": "+6281234567890",
    "address": "Jl. Merdeka No. 45",
    "role": "guest",
    "is_verified": "true"
  }
}
```

### Get Users
Retrieves list of all users (admin only).

**GET** `/api/ums/v1/users`

**Headers:**
```
Authorization: <token>
```

**Response:**
```json
{
  "message": "success",
  "data": [
    {
      "id": 1,
      "photo_profile": "path_url",
      "username": "indra_123",
      "email": "indra@example.com",
      "name": "Indra Wijaya",
      "phone": "+6281234567890",
      "address": "Jl. Merdeka No. 45",
      "role": "guest",
      "is_verified": "false"
    },
    {
      "id": 2,
      "photo_profile": "path_url",
      "username": "budi_456",
      "email": "budi@example.com",
      "name": "Budi Santoso",
      "phone": "+6289876543210",
      "address": "Jl. Baru No. 99",
      "role": "admin",
      "is_verified": "true"
    }
  ]
}
```

### Update Profile
Updates user profile information.

**PUT** `/api/user/v1/profile`

**Headers:**
```
Authorization: <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Indra Wansyah Baru",
  "phone": "+6289876543210",
  "address": "Jl. Baru No. 99"
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "photo_profile": "path_url",
    "username": "indra_123",
    "email": "indra@example.com",
    "name": "Indra Wansyah Baru",
    "phone": "+6281234567890",
    "address": "Jl. Merdeka No. 45",
    "role": "guest",
    "is_verified": "true"
  }
}
```

### Logout
Invalidates user's access token.

**DELETE** `/api/user/v1/logout`

**Headers:**
```
Authorization: <token>
```

**Response:**
```json
{
  "message": "success"
}
```

### Google OAuth
Endpoints for Google OAuth authentication.

#### Login
**GET** `/api/user/v1/auth/google/login`

#### Callback
**GET** `/api/user/v1/auth/google/callback`

**Response:**
```json
{
  "message": "success"
}
```