# GO Ecommerce Server

This is a pretty simple backend for an Ecommerce site. This is made with Go lang and built with **Std** library and **Mux** router.

Note: This repository is a good starting point for anyone who wants to transition from a library based language like Node.js or Python to something barebon like Go or Rust, where they dont need anything more that the **std** library.

---

## Features
- Full Fledged token based authentication (JWT)
- Full CRUD for Products and User accounts
- Mysql database handling with raw SQL queries
- Migration Sql scripts
- Followed by the best architecture from the trending projects.

---

## Endpoints

#### Users

#### Products
GET - /api/v1/products (Get all products)

POST - /api/v1/products (create product, NOTE - Requires the JWT token to be passed as well in the headers)
```bash
curl -X POST http://localhost:8080/api/v1/products \
     -H "Content-Type: application/json" \
     -H "Authorization: JWT_TOKEN" \
     -d '{"name": "Demo", "description": "Demo description for the demo product", "image": "", "price": 69.00, "quantity": 5 }
```
---

~ Piush - Copyright 2025
