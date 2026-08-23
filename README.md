# Slotly

> **A production-style, multi-tenant appointment and booking SaaS backend built for service-based businesses.**

Slotly is a REST API platform that enables businesses such as **clinics, therapists, consultants, coaches, salons, and service providers** to manage their organization, staff, services, availability, and customer appointments.

The project is designed to demonstrate how a real-world SaaS backend can be structured with **multi-tenancy, authentication, role-based access control, scheduling conflict prevention, database transactions, rate limiting, and production-ready API design**.

## 🚀 Live Demo

| Resource              | URL           |
| --------------------- | ------------- |
| Live API              | `Coming Soon` |
| Swagger Documentation | `Coming Soon` |
| API Health Check      | `Coming Soon` |

---

# ✨ Features

## Authentication & Security

* JWT-based authentication
* Access and refresh tokens
* Secure password hashing
* Role-Based Access Control (RBAC)
* Rate limiting on public endpoints
* Protected organization resources
* Input validation
* Structured error responses
* Environment-based configuration

## Multi-Tenant Architecture

Each business operates as an isolated organization.

For example:

* One clinic cannot access another clinic's appointments
* One salon cannot modify another salon's services
* Staff members can only access resources belonging to their organization

Tenant isolation is enforced at the application and database query layers.

## Appointment Management

* Create appointments
* Update appointments
* Cancel appointments
* Reschedule appointments
* Appointment status management
* Prevent double booking
* Validate staff availability
* Validate service duration
* Customer booking history

## Staff Management

* Invite or create staff members
* Assign roles
* Configure staff availability
* Assign services to staff
* Manage schedules

## Service Management

Businesses can:

* Create services
* Update services
* Define service duration
* Define pricing
* Assign services to staff
* Enable or disable services

## Customer Management

* Create customers
* View appointment history
* Search customers
* Manage customer profiles

---

# 🏗 Architecture

```mermaid
flowchart TB

    Client[Web / Mobile Client]

    Client --> API[REST API]

    API --> Middleware[Middleware Layer]

    Middleware --> Auth[Authentication]
    Middleware --> RateLimit[Rate Limiting]
    Middleware --> Validation[Request Validation]
    Middleware --> RBAC[Role Authorization]

    Auth --> Controllers
    RateLimit --> Controllers
    Validation --> Controllers
    RBAC --> Controllers

    Controllers --> Services[Service Layer]

    Services --> UserService[User Service]
    Services --> OrgService[Organization Service]
    Services --> BookingService[Booking Service]
    Services --> AvailabilityService[Availability Service]

    UserService --> Repository[Repository Layer]
    OrgService --> Repository
    BookingService --> Repository
    AvailabilityService --> Repository

    Repository --> DB[(PostgreSQL)]

    Services --> Cache[(Redis)]
    Services --> Jobs[Background Jobs]

    Jobs --> Notifications[Email Notifications]
```

## Architectural Principles

The application follows a layered architecture:

```text
Routes
  ↓
Middleware
  ↓
Controllers
  ↓
Services
  ↓
Repositories
  ↓
Database
```

### Responsibilities

**Routes**

* Define API endpoints
* Attach middleware
* Map requests to controllers

**Middleware**

* Authentication
* Authorization
* Rate limiting
* Validation
* Error handling

**Controllers**

* Handle HTTP requests and responses
* Keep business logic minimal

**Services**

* Contain core business logic
* Enforce scheduling rules
* Coordinate database transactions

**Repositories**

* Database access
* Query abstraction

---

# 👥 Roles & Permissions

Slotly uses Role-Based Access Control.

| Role          | Description                                    |
| ------------- | ---------------------------------------------- |
| `SUPER_ADMIN` | Platform-level administrator                   |
| `ORG_ADMIN`   | Manages an organization                        |
| `STAFF`       | Manages assigned appointments and availability |
| `CUSTOMER`    | Books and manages appointments                 |

## Example Permission Matrix

| Action                     | Super Admin | Org Admin | Staff | Customer |
| -------------------------- | :---------: | :-------: | :---: | :------: |
| Manage organizations       |      ✅      |     ❌     |   ❌   |     ❌    |
| Manage organization staff  |      ✅      |     ✅     |   ❌   |     ❌    |
| Create services            |      ❌      |     ✅     |   ❌   |     ❌    |
| Manage own availability    |      ❌      |     ❌     |   ✅   |     ❌    |
| View assigned appointments |      ❌      |     ✅     |   ✅   |     ❌    |
| Book appointment           |      ❌      |     ❌     |   ❌   |     ✅    |
| Cancel own appointment     |      ❌      |     ❌     |   ❌   |     ✅    |

---

# 🗂 Project Structure

```text
src/
├── config/
│   ├── database
│   ├── redis
│   └── environment
│
├── modules/
│   ├── auth/
│   ├── users/
│   ├── organizations/
│   ├── staff/
│   ├── services/
│   ├── availability/
│   ├── appointments/
│   └── customers/
│
├── middleware/
│   ├── authentication
│   ├── authorization
│   ├── validation
│   ├── rate-limit
│   └── error-handler
│
├── shared/
│   ├── errors
│   ├── utils
│   └── constants
│
├── routes/
├── jobs/
└── app
```

---

# 🗄 Database Design

The core entities are:

```mermaid
erDiagram

    USER ||--o{ ORGANIZATION_MEMBER : belongs_to
    ORGANIZATION ||--o{ ORGANIZATION_MEMBER : has

    ORGANIZATION ||--o{ SERVICE : offers
    ORGANIZATION ||--o{ STAFF : employs
    ORGANIZATION ||--o{ CUSTOMER : manages

    STAFF ||--o{ STAFF_AVAILABILITY : has
    STAFF ||--o{ STAFF_SERVICE : provides

    SERVICE ||--o{ STAFF_SERVICE : assigned_to

    CUSTOMER ||--o{ APPOINTMENT : books
    STAFF ||--o{ APPOINTMENT : handles
    SERVICE ||--o{ APPOINTMENT : scheduled_for

    USER {
        uuid id PK
        string email
        string password_hash
        string role
    }

    ORGANIZATION {
        uuid id PK
        string name
        string timezone
    }

    SERVICE {
        uuid id PK
        uuid organization_id FK
        string name
        integer duration_minutes
        decimal price
    }

    STAFF {
        uuid id PK
        uuid organization_id FK
        uuid user_id FK
    }

    CUSTOMER {
        uuid id PK
        uuid organization_id FK
        string name
        string email
    }

    APPOINTMENT {
        uuid id PK
        uuid organization_id FK
        uuid customer_id FK
        uuid staff_id FK
        uuid service_id FK
        timestamp start_time
        timestamp end_time
        string status
    }
```

---

# 🔒 Multi-Tenant Isolation

Every tenant-specific resource is associated with an `organization_id`.

For example:

```text
Appointment
    ↓
organization_id
```

Every protected query validates the organization context.

Example:

```text
GET /appointments/:id

WHERE
    appointment.id = :appointmentId
    AND appointment.organization_id = :currentOrganizationId
```

This prevents users from accessing resources belonging to another organization.

---

# 📅 Booking Flow

```mermaid
sequenceDiagram

    Customer->>API: Create Appointment Request
    API->>Auth: Validate JWT
    Auth-->>API: Authenticated

    API->>Validation: Validate Request
    Validation-->>API: Valid

    API->>BookingService: Check Availability

    BookingService->>Database: Check Staff Schedule
    Database-->>BookingService: Available

    BookingService->>Database: Check Existing Appointments
    Database-->>BookingService: No Conflict

    BookingService->>Database: Create Appointment

    Database-->>BookingService: Appointment Created

    BookingService-->>API: Success
    API-->>Customer: 201 Created
```

---

# 🛡 Preventing Double Bookings

One of the important business problems handled by Slotly is preventing conflicting appointments.

Before creating an appointment, the system checks:

1. Whether the staff member exists
2. Whether the staff member belongs to the organization
3. Whether the requested service exists
4. Whether the staff member provides the requested service
5. Whether the requested time falls within staff availability
6. Whether another appointment already overlaps with the requested time

An appointment is rejected when an overlap exists.

Example:

```text
Existing Appointment
10:00 ───────── 11:00

Requested Appointment
       10:30 ───────── 11:30

Result: Conflict ❌
```

---

# 📡 API Endpoints

## Authentication

```text
POST   /api/v1/auth/register
POST   /api/v1/auth/login
POST   /api/v1/auth/refresh
POST   /api/v1/auth/logout
```

## Organizations

```text
POST   /api/v1/organizations
GET    /api/v1/organizations/:id
PATCH  /api/v1/organizations/:id
```

## Staff

```text
POST   /api/v1/staff
GET    /api/v1/staff
GET    /api/v1/staff/:id
PATCH  /api/v1/staff/:id
DELETE /api/v1/staff/:id
```

## Services

```text
POST   /api/v1/services
GET    /api/v1/services
GET    /api/v1/services/:id
PATCH  /api/v1/services/:id
DELETE /api/v1/services/:id
```

## Availability

```text
POST   /api/v1/staff/:staffId/availability
GET    /api/v1/staff/:staffId/availability
PATCH  /api/v1/availability/:id
DELETE /api/v1/availability/:id
```

## Appointments

```text
POST   /api/v1/appointments
GET    /api/v1/appointments
GET    /api/v1/appointments/:id
PATCH  /api/v1/appointments/:id
DELETE /api/v1/appointments/:id

POST   /api/v1/appointments/:id/cancel
POST   /api/v1/appointments/:id/reschedule
```

---

# 🧪 Example API Request

### Create Appointment

```http
POST /api/v1/appointments
Authorization: Bearer <access_token>
Content-Type: application/json
```

```json
{
  "customerId": "customer_uuid",
  "staffId": "staff_uuid",
  "serviceId": "service_uuid",
  "startTime": "2026-08-25T10:00:00Z"
}
```

### Successful Response

```json
{
  "success": true,
  "message": "Appointment created successfully",
  "data": {
    "id": "appointment_uuid",
    "status": "CONFIRMED",
    "startTime": "2026-08-25T10:00:00Z",
    "endTime": "2026-08-25T11:00:00Z"
  }
}
```

---

# ❌ Error Handling

All errors follow a consistent structure.

```json
{
  "success": false,
  "error": {
    "code": "APPOINTMENT_CONFLICT",
    "message": "The selected time slot is no longer available"
  }
}
```

Common error codes:

```text
VALIDATION_ERROR
UNAUTHORIZED
FORBIDDEN
RESOURCE_NOT_FOUND
APPOINTMENT_CONFLICT
RATE_LIMIT_EXCEEDED
INTERNAL_SERVER_ERROR
```

---

# ⚡ Rate Limiting

Public endpoints are protected against abuse.

Examples include:

```text
POST /auth/login
POST /auth/register
POST /appointments
```

Rate limits can be configured using environment variables.

Example:

```text
RATE_LIMIT_WINDOW_MS=60000
RATE_LIMIT_MAX_REQUESTS=100
```

---

# 🧰 Tech Stack

The project can be implemented using:

| Layer             | Technology             |
| ----------------- | ---------------------- |
| Runtime           | Node.js                |
| Framework         | Express.js / Fastify   |
| Language          | TypeScript             |
| Database          | PostgreSQL             |
| ORM               | Prisma / Drizzle       |
| Authentication    | JWT                    |
| Validation        | Zod                    |
| API Documentation | Swagger / OpenAPI      |
| Cache             | Redis                  |
| Containerization  | Docker                 |
| Deployment        | Render / Railway / AWS |

---

# 🚀 Getting Started

## Prerequisites

Make sure you have installed:

* Node.js 20+
* PostgreSQL
* Redis
* Docker (optional)

## Clone the Repository

```bash
git clone https://github.com/YOUR_USERNAME/slotly-backend.git
cd slotly-backend
```

## Install Dependencies

```bash
npm install
```

## Configure Environment Variables

Create a `.env` file:

```env
PORT=3000
NODE_ENV=development

DATABASE_URL=postgresql://user:password@localhost:5432/slotly

JWT_ACCESS_SECRET=your_access_secret
JWT_REFRESH_SECRET=your_refresh_secret

ACCESS_TOKEN_EXPIRY=15m
REFRESH_TOKEN_EXPIRY=7d

REDIS_URL=redis://localhost:6379

RATE_LIMIT_WINDOW_MS=60000
RATE_LIMIT_MAX_REQUESTS=100
```

## Run Database Migrations

```bash
npm run migration:run
```

## Start Development Server

```bash
npm run dev
```

The API should now be available at:

```text
http://localhost:3000
```

Swagger documentation:

```text
http://localhost:3000/api-docs
```

---

# 🐳 Running with Docker

```bash
docker compose up --build
```

This starts:

* API
* PostgreSQL
* Redis

---

# 🧪 Testing

Run the test suite:

```bash
npm test
```

Run tests with coverage:

```bash
npm run test:coverage
```

The test suite includes:

* Unit tests
* Service-layer tests
* Integration tests
* Authentication tests
* Authorization tests
* Appointment conflict tests

---

# 🔮 Roadmap

## Phase 1 — Core Platform

* [x] Project architecture
* [ ] Authentication
* [ ] JWT access and refresh tokens
* [ ] Organization management
* [ ] Role-based access control

## Phase 2 — Booking System

* [ ] Staff management
* [ ] Service management
* [ ] Customer management
* [ ] Availability management
* [ ] Appointment creation
* [ ] Double-booking prevention
* [ ] Rescheduling
* [ ] Cancellation

## Phase 3 — Production Features

* [ ] Redis caching
* [ ] Background jobs
* [ ] Email notifications
* [ ] Audit logs
* [ ] Docker support
* [ ] CI/CD pipeline

## Phase 4 — Advanced Features

* [ ] Calendar integrations
* [ ] Google Calendar sync
* [ ] Payment integration
* [ ] Webhooks
* [ ] Usage analytics
* [ ] Subscription plans

---

# 💡 What This Project Demonstrates

Slotly is intentionally designed to demonstrate more than basic CRUD operations.

It showcases experience with:

* Designing RESTful APIs
* Multi-tenant SaaS architecture
* JWT authentication
* Role-based authorization
* Relational database design
* Database migrations
* Transaction management
* Complex scheduling logic
* Preventing race conditions and double bookings
* Input validation
* Structured error handling
* API rate limiting
* Swagger/OpenAPI documentation
* Docker-based development
* Production deployment

---

# 📄 License

This project is licensed under the MIT License.

---

## Built to demonstrate production-oriented backend engineering.

If you are looking for a backend developer to build a SaaS platform, booking system, internal tool, marketplace, or API-driven application, this project demonstrates the type of backend architecture and engineering practices I can deliver.