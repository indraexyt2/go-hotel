# Booking Service Documentation

## Overview
This document outlines the API endpoints for the Booking Service. The service provides functionality for managing hotel room bookings, including creating, retrieving, and updating booking information.

## Base URL
All API endpoints are prefixed with `/api/booking/v1`

## Endpoints

### Create Booking
**POST** `/bookings`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "room_id": 1,
  "guest_id": 2,
  "checkin_date": "2025-02-01",
  "checkout_date": "2025-02-05",
  "total_price": 200.00,
  "status": "pending"
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "room_id": 1,
    "guest_id": 2,
    "checkin_date": "2025-02-01",
    "checkout_date": "2025-02-05",
    "total_price": 200.00,
    "status": "pending",
    "created_at": "2025-01-29T10:00:00Z",
    "updated_at": "2025-01-29T10:00:00Z"
  }
}
```

### Get Booking by ID
**GET** `/bookings/{bookingID}`

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "room_id": 1,
    "guest_id": 2,
    "checkin_date": "2025-02-01",
    "checkout_date": "2025-02-05",
    "total_price": 200.000,
    "status": "pending",
    "created_at": "2025-01-29T10:00:00Z",
    "updated_at": "2025-01-29T10:00:00Z"
  }
}
```

### Get User Bookings
**GET** `/bookings/user/{userID}`

**Response:**
```json
{
  "message": "success",
  "data": [
    {
      "id": 1,
      "room_id": 1,
      "guest_id": 1,
      "checkin_date": "2025-02-01",
      "checkout_date": "2025-02-05",
      "total_price": 200.00,
      "status": "pending",
      "created_at": "2025-01-29T10:00:00Z",
      "updated_at": "2025-01-29T10:00:00Z"
    },
    {
      "id": 2,
      "room_id": 3,
      "guest_id": 1,
      "checkin_date": "2025-02-10",
      "checkout_date": "2025-02-15",
      "total_price": 350.00,
      "status": "confirmed",
      "created_at": "2025-01-30T12:00:00Z",
      "updated_at": "2025-01-30T12:30:00Z"
    }
  ]
}
```

### List All Bookings
**GET** `/bookings`

**Response:**
```json
{
  "message": "success",
  "data": [
    {
      "id": 1,
      "room_id": 1,
      "guest_id": 2,
      "checkin_date": "2025-02-01",
      "checkout_date": "2025-02-05",
      "total_price": 200.00,
      "status": "pending",
      "created_at": "2025-01-29T10:00:00Z",
      "updated_at": "2025-01-29T10:00:00Z"
    },
    {
      "id": 2,
      "room_id": 3,
      "guest_id": 4,
      "checkin_date": "2025-02-10",
      "checkout_date": "2025-02-15",
      "total_price": 350.00,
      "status": "confirmed",
      "created_at": "2025-01-30T12:00:00Z",
      "updated_at": "2025-01-30T12:30:00Z"
    }
  ]
}
```

### Update Booking
**PUT** `/bookings/{bookingID}`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "checkin_date": "2025-02-02",
  "checkout_date": "2025-02-06",
  "total_price": 220.000,
  "status": "resechedule"
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "room_id": 1,
    "guest_id": 2,
    "checkin_date": "2025-02-01",
    "checkout_date": "2025-02-05",
    "total_price": 200.000,
    "status": "resechedule"
  }
}
```

### Update Booking Status
**PATCH** `/bookings/{id}/status`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "status": "completed"
}
```

**Response:**
```json
{
  "message": "success",
  "data": {
    "id": 1,
    "status": "completed"
  }
}
```

## Booking Status Types
The booking service supports the following status types:
- `pending`: Initial booking status
- `confirmed`: Booking has been confirmed
- `completed`: Stay has been completed
- `resechedule`: Booking dates have been modified