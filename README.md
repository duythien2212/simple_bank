# simple_bank

## Project Description
`simple_bank` is a backend service for managing bank accounts. It provides APIs for creating accounts, performing transactions, and tracking balances.

## Features
- Create and manage bank accounts.
- Perform transfers.
- Track account balances and transaction history.
- Secure authentication and authorization for users.

## Technologies Used
- **Programming Language**: Go (Golang)
- **Database**: PostgreSQL
- **API Framework**: Gin
- **Authentication**: JSON Web Tokens (JWT)

## Getting Started
1. Clone the repository:
   ```bash
   git clone https://github.com/duythien2212/simple_bank.git
   ```
2. Start the PostgreSQL database using Docker:
   ```bash
   make postgres
   ```
3. Create the database:
   ```bash
   make createdb
   ```
4. Run database migrations:
   ```bash
   make migrateup
   ```
5. (Optional) Generate SQL code using `sqlc`:
   ```bash
   make sqlc
   ```
6. Start the server:
   ```bash
   make server
   ```