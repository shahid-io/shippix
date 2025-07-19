# SHIPPIX

SHIPPIX is a shipment tracking system that allows users to keep track of their shipments from different carriers. It provides a unified API for multiple carriers allowing users to get the latest status of their shipments.

## Features

- User Authentication:
  - OAuth 2.0 (Google, GitHub, etc.)
  - Phone OTP Login
  - Email OTP Login
- Create & Manage Shipments
- Shipment Status Tracking (Real-Time & Manual)
- Timeline of Shipment Events
- Admin & User Dashboards (Future)
- Notifications (Email/SMS – Optional)
- API-first Design

## User Roles

| Role   | Permissions                              |
|--------|------------------------------------------|
| User   | Can create/view shipments                |
| Admin  | Manage all users and shipments           |
| Carrier| Update shipment events/status (Optional) |

## Authentication Methods

1. **OAuth 2.0**
   - Google, GitHub, etc.
2. **Phone Number OTP**
   - Via Twilio/SMS gateway
3. **Email OTP**
   - Via SMTP or SendGrid

## Database Schema Overview

### `users`
| Column         | Type     | Description                     |
|----------------|----------|---------------------------------|
| id             | UUID     | Primary Key                     |
| email          | VARCHAR  | Nullable                        |
| phone_number   | VARCHAR  | Nullable                        |
| oauth_provider | VARCHAR  | e.g. `google`                   |
| oauth_id       | VARCHAR  | Unique external ID              |
| created_at     | TIMESTAMP|                                 |
| updated_at     | TIMESTAMP|                                 |

### `otp_tokens`
| Column     | Type     | Description              |
|------------|----------|--------------------------|
| id         | UUID     |                          |
| contact    | VARCHAR  | Email or Phone Number    |
| otp_code   | VARCHAR  | OTP sent to user         |
| type       | VARCHAR  | `email` or `phone`       |
| expires_at | TIMESTAMP| OTP Expiry               |
| verified   | BOOLEAN  | OTP status               |
| created_at | TIMESTAMP|                          |

### `shipments`
| Column              | Type     | Description              |
|---------------------|----------|--------------------------|
| id                  | UUID     |                          |
| user_id             | UUID     | FK to users              |
| origin              | TEXT     |                          |
| destination         | TEXT     |                          |
| carrier_name        | VARCHAR  | FedEx, UPS, etc.         |
| status              | VARCHAR  | `pending`, `delivered`   |
| estimated_delivery  | TIMESTAMP| Optional ETA             |
| created_at          | TIMESTAMP|                          |
| updated_at          | TIMESTAMP|                          |

### `shipment_events`
| Column      | Type     | Description              |
|-------------|----------|--------------------------|
| id          | UUID     |                          |
| shipment_id | UUID     | FK to shipments          |
| location    | TEXT     | Event location           |
| status      | VARCHAR  | `dispatched`, etc.       |
| event_time  | TIMESTAMP|                          |
| notes       | TEXT     | Optional                 |

## Project Structure
