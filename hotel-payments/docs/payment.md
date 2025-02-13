# Payment Service Documentation

## Overview
This document outlines the API endpoints for the Payment Service. The service handles payment processing, notifications, and refunds through Midtrans integration.

## Base URL
All API endpoints are prefixed with `/api/payment/v1`

## Endpoints

### Midtrans Notification Handler
Handles payment notifications from Midtrans.

**POST** `/midtrans/notification`

**Response:**
```json
{
  "message": "success"
}
```

### Process Refund
Processes refund requests for bookings.

**POST** `/midtrans/refund`

**Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "booking_id": 1,
  "reason": "i lost my money"
}
```

**Response:**
```json
{
  "message": "success"
}
```

### Get Payment Details
Retrieves payment details for a specific booking.

**GET** `/:bookingID`

**Response:**
Contains payment details for the specified booking ID.