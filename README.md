# URL Shortener

A simple and efficient URL shortener application built with Go (Gin) and React (Vite).

## Tech Stack

- **Backend:** Go, Gin, PostgreSQL, MongoDB (Driver included in go.mod)
- **Frontend:** React, Vite
- **Database:** PostgreSQL
- **Deployment:** Vercel

## Project Structure

- `service/`: Contains the core application logic and services.
- `router/`: Handles HTTP routing and middleware.
- `api/`: Vercel Serverless Function entry point (`index.go`).
- `frontend/`: React frontend application.
- `config/`: Configuration management.
- `db/`: Database connection logic.

## Local Development Setup

### Prerequisites

- Go 1.24+
- Node.js & npm
- PostgreSQL

### Backend

1.  Clone the repository.
2.  Create a `.env` file in the root directory with your database credentials (see `.env.example`).
3.  Run the server:

    ```bash
    go run main.go
    ```

    The server will start on port `8003`.

### Frontend

1.  Navigate to the frontend directory:

    ```bash
    cd frontend
    ```

2.  Install dependencies:

    ```bash
    npm install
    ```

3.  Start the development server:

    ```bash
    npm run dev
    ```

    The frontend will be available at `http://localhost:5173` (or similar).

## Deployment

This project is configured for deployment on **Vercel**.

- The `api/index.go` file serves as the entry point for Vercel Serverless Functions.
- The `service` directory contains the business logic, separated from the Vercel handler to avoid conflicts.