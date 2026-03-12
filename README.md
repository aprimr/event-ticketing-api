# 🎟️ Event Ticketing API

A RESTful API built with **Go** and **PostgreSQL** for managing events and ticket bookings. Features capacity management using database transactions, pagination, and dynamic filtering.

---

## 🚀 Tech Stack

| Tool              | Purpose                      |
| ----------------- | ---------------------------- |
| Go (net/http)     | HTTP server and routing      |
| PostgreSQL (Neon) | Cloud database               |
| pgx/v5            | PostgreSQL driver            |
| godotenv          | Environment variable loading |

---

## Endpoints

### Events

| Method   | Endpoint                  | Description                |
| -------- | ------------------------- | -------------------------- |
| `POST`   | `/events`                 | Create a new event         |
| `GET`    | `/events`                 | Get all events (paginated) |
| `GET`    | `/events?category=music`  | Filter events by category  |
| `GET`    | `/events?date=2024-06-15` | Filter events by date      |
| `GET`    | `/events?page=1&limit=10` | Paginate results           |
| `GET`    | `/events/:id`             | Get single event by ID     |
| `PUT`    | `/events/:id`             | Update event by ID         |
| `DELETE` | `/events/:id`             | Delete event by ID         |

### Bookings

| Method   | Endpoint                    | Description                   |
| -------- | --------------------------- | ----------------------------- |
| `POST`   | `/events/:id/bookings`      | Book tickets for an event     |
| `GET`    | `/events/:id/bookings`      | Get all bookings for an event |
| `DELETE` | `/events/:id/bookings/:bid` | Cancel a booking              |

---

### Environment Variables

```env
DATABASE_URL=your_postgresql_connection_string
PORT=8000
```

---

## 📬 Example Requests

### Create Event

```bash
curl -X POST http://localhost:8000/events \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Go Conference 2024",
    "description": "Annual Go developer conference",
    "location": "New York",
    "category": "tech",
    "capacity": 100,
    "price": 49.99,
    "event_date": "2024-06-15"
  }'
```

### Book Tickets

```bash
curl -X POST http://localhost:8000/events/1/bookings \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "seats": 2
  }'
```

### Get Events with Filters

```bash
curl "http://localhost:8000/events?category=tech&page=1&limit=10"
```

---

## 📦 Response Format

Every response follows a consistent structure:

**Success**

```json
{
  "success": true,
  "message": "Booking created",
  "data": { ... }
}
```

**Error**

```json
{
  "success": false,
  "message": "Not enough seats"
}
```
