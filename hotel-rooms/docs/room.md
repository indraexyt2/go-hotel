# Room Service Documentation

## Overview
This document outlines the API endpoints for the Room Service. The service provides functionality for managing hotel rooms, room types, features, and photos.

## Base URL
All API endpoints are prefixed with `/api/room/v1`

## Endpoints

### Room Types

#### List All Room Types
**GET** `/room-types`

**Headers:**
```
Content-Type: application/json
```

**Response:**
```json
{
  "message": "success",
  "data": [
    {
      "id": 1,
      "name": "Deluxe Room",
      "description": "A luxurious room with a king-sized bed and ocean view.",
      "price_per_night": 400.000,
      "capacity": 2,
      "total_rooms": 10
    }
  ]
}
```

#### Get Room Type Details
**GET** `/room-types/{id}`

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "name": "Deluxe Room",
    "description": "A luxurious room with a king-sized bed and ocean view.",
    "price_per_night": 400.000,
    "capacity": 2,
    "total_rooms": 10
  }
}
```

#### Create Room Type
**POST** `/room-types`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Family Room",
  "description": "A spacious room suitable for families.",
  "price_per_night": 200.00,
  "capacity": 4,
  "total_rooms": 5
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 3,
    "name": "Family Room",
    "description": "A spacious room suitable for families.",
    "price_per_night": 200.00,
    "capacity": 4,
    "total_rooms": 5
  }
}
```

#### Update Room Type
**PUT** `/room-types/{id}`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Executive Suite",
  "description": "A premium suite with a living area and private balcony.",
  "price_per_night": 400.000,
  "capacity": 3,
  "total_rooms": 8
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "name": "Executive Suite",
    "description": "A premium suite with a living area and private balcony.",
    "price_per_night": 400.000,
    "capacity": 3,
    "total_rooms": 8
  }
}
```

#### Delete Room Type
**DELETE** `/room-types/{id}`

**Response:**
```json
{
  "message": "success"
}
```

### Room Features

#### Add Room Feature
**POST** `/room-types/features`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "room_type_id": 1,
  "features": "Mini Bar"
}
```

**Response:**
```json
{
  "message": "success"
}
```

#### List Room Features
**GET** `/room-types/{room_type_id}/features`

**Response:**
```json
{
  "message": "success",
  "data": [
    {
      "id": 1,
      "room_type_id": 1,
      "features": "Free Wi-Fi"
    },
    {
      "id": 2,
      "room_type_id": 1,
      "features": "Air Conditioning"
    }
  ]
}
```

#### Update Room Feature
**PUT** `/room-types/features/{id}`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "room_type_id": 1,
  "features": "Free Breakfast"
}
```

**Response:**
```json
{
  "message": "success"
}
```

#### Delete Room Feature
**DELETE** `/room-types/{room_type_id}/features/{id}`

**Response:**
```json
{
  "message": "success"
}
```

### Room Photos

#### Upload Room Photo
**POST** `/room-types/{room_type_id}/photos`

**Headers:**
```
Content-Type: multipart/form-data
```

**Request Body:**
- Form Data:
    - `file`: Image file (room photo)
    - `is_primary`: Boolean (optional, default `false`)

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 3,
    "room_type_id": 1,
    "file_path": "https://example.com/photos/room3.jpg",
    "is_primary": false
  }
}
```

#### List Room Photos
**GET** `/room-types/{room_type_id}/photos`

**Response:**
```json
{
  "message": "success",
  "data": [
    {
      "id": 1,
      "room_type_id": 1,
      "file_path": "https://example.com/photos/room1.jpg",
      "is_primary": true
    },
    {
      "id": 2,
      "room_type_id": 1,
      "file_path": "https://example.com/photos/room2.jpg",
      "is_primary": false
    }
  ]
}
```

#### Delete Room Photo
**DELETE** `/room-types/photos/{id}`

**Response:**
```json
{
  "message": "success"
}
```

### Individual Rooms

#### Create Room
**POST** `/rooms`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "room_number": "103",
  "room_type_id": 1,
  "description": "A cozy room with a single bed."
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 3,
    "room_number": "103",
    "room_type_id": 1,
    "description": "A cozy room with a single bed."
  }
}
```

#### List All Rooms
**GET** `/rooms`

**Response:**
```json
{
  "message": "success",
  "data": [
    {
      "id": 1,
      "room_number": "101",
      "room_type_id": 1,
      "description": "A comfortable room with a queen-sized bed.",
      "room_type": {
        "id": 1,
        "name": "Deluxe Room",
        "description": "A luxurious room with a king-sized bed and ocean view.",
        "price_per_night": 150.00,
        "capacity": 2,
        "total_rooms": 10
      },
      "features": [
        {
          "id": 1,
          "room_type_id": 1,
          "features": "Free Wi-Fi"
        },
        {
          "id": 2,
          "room_type_id": 1,
          "features": "Air Conditioning"
        }
      ],
      "photos": [
        {
          "id": 1,
          "room_type_id": 1,
          "file_path": "https://example.com/photos/room1.jpg",
          "is_primary": true
        }
      ]
    }
  ]
}
```

#### Get Room Details
**GET** `/rooms/{id}`

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "room_number": "101",
    "room_type_id": 1,
    "description": "A comfortable room with a queen-sized bed.",
    "room_type": {
      "id": 1,
      "name": "Deluxe Room",
      "description": "A luxurious room with a king-sized bed and ocean view.",
      "price_per_night": 150.00,
      "capacity": 2,
      "total_rooms": 10
    },
    "features": [
      {
        "id": 1,
        "room_type_id": 1,
        "features": "Free Wi-Fi"
      }
    ],
    "photos": [
      {
        "id": 1,
        "room_type_id": 1,
        "file_path": "https://example.com/photos/room1.jpg",
        "is_primary": true
      }
    ]
  }
}
```

#### Edit Room
**PUT** `/rooms/{id}`

**Request Body:**
```json
{
  "room_number": "101",
  "room_type_id": 2,
  "status": "occupied",
  "description": "A luxurious room with a king-sized bed."
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "room_number": "101",
    "room_type_id": 2,
    "description": "A luxurious room with a king-sized bed."
  }
}
```