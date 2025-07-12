# Vyapaar
Vyapaar is a simple and scalable e-commerce platform inspired by Indian local trade, enabling users to sell products online with ease.




# 🛒 Product Sale Platform - API Endpoints

## 🔐 Authentication & User Management

| Method | Endpoint       | Description                              |
|--------|----------------|------------------------------------------|
| POST   | `/register`    | Register a new user                      |
| POST   | `/login`       | User login, returns JWT token            |
| POST   | `/logout`      | Logout user                              |
| GET    | `/profile`     | Get current user profile                 |
| PUT    | `/profile`     | Update user profile                      |
| GET    | `/users/:id`   | (Admin) Get specific user details        |

---

## 🛍️ Product Management

| Method | Endpoint                    | Description                            |
|--------|-----------------------------|----------------------------------------|
| POST   | `/products`                 | Create/upload new product (seller)     |
| GET    | `/products`                 | Get all products with filters/search   |
| GET    | `/products/:id`            | Get product details                    |
| PUT    | `/products/:id`            | Update product (seller only)           |
| DELETE | `/products/:id`            | Delete product (seller only)           |
| POST   | `/products/:id/images`     | Upload images for a product            |

---

## 🛒 Cart Management

| Method | Endpoint              | Description                              |
|--------|-----------------------|------------------------------------------|
| POST   | `/cart`               | Add item to cart                         |
| GET    | `/cart`               | View cart items                          |
| PUT    | `/cart/:itemId`       | Update cart item quantity                |
| DELETE | `/cart/:itemId`       | Remove item from cart                    |

---

## 📦 Orders

| Method | Endpoint                | Description                              |
|--------|-------------------------|------------------------------------------|
| POST   | `/orders`              | Place a new order                         |
| GET    | `/orders`              | Get current user's orders                 |
| GET    | `/orders/:id`         | Get order details                         |
| PUT    | `/orders/:id/status`  | (Admin/Seller) Update order status        |

---

## 💳 Payments

| Method | Endpoint                        | Description                            |
|--------|----------------------------------|----------------------------------------|
| POST   | `/payments/initiate`            | Initiate payment                        |
| POST   | `/payments/verify`              | Verify payment after redirect          |
| GET    | `/payments/:orderId/status`     | Check payment status for an order      |

---

## ⭐ Reviews & Ratings

| Method | Endpoint                        | Description                            |
|--------|----------------------------------|----------------------------------------|
| POST   | `/products/:id/reviews`         | Add a review for a product             |
| GET    | `/products/:id/reviews`         | Get reviews for a product              |
| DELETE | `/reviews/:id`                  | Delete a review (user or admin)        |

---

## 🛠️ Admin Management (Optional)

| Method | Endpoint                        | Description                            |
|--------|----------------------------------|----------------------------------------|
| GET    | `/admin/users`                 | View all users                          |
| GET    | `/admin/products`              | View all products                       |
| PUT    | `/admin/products/:id/approve` | Approve or reject a product            |
| GET    | `/admin/orders`               | View all orders                         |

---

## 🧩 Miscellaneous

| Method | Endpoint              | Description                              |
|--------|-----------------------|------------------------------------------|
| GET    | `/categories`        | Get product categories                    |
| GET    | `/tags`              | Get product tags (optional)               |
| GET    | `/search?q=...`      | Search products by keyword                |

---


