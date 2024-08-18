# Clean Architecture Go Backend

This repository contains the initialization of a Go backend structured using Clean Architecture principles and Domain-Driven Design (DDD). The project serves as a boilerplate for building scalable, maintainable, and testable backend systems in Go.

## Table of Contents

- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)

## Introduction

This project is my personal example of how to set up a Go backend using Clean Architecture and Domain-Driven Design (DDD). The goal is to separate concerns, enhance testability, and provide a foundation for long-term project growth.

## Project Structure

The project is organized into several layers according to Clean Architecture principles:

- `domain/` - Contains domain entities and core business logic.
- `infrastructure/`
  - `db/` - Database-related code, such as MongoDB and PostgreSQL connections.
    - `mongodb.go` - MongoDB connection setup.
    - `postgres.go` - PostgreSQL connection setup.
  - `router/` - HTTP router setup for handling routes.
    - `mux-router.go` - Configuration for mux router.
    - `router.go` - General router setup and configuration.
- `interface/`
  - `controllers/` - Contains controller logic to handle requests and responses.
  - `repository/` - Interfaces for repository patterns to abstract data access.
- `usecases/` - Application-specific business logic and use case implementations.
- `tests/`
  - `e2e/` - End-to-end tests for testing the application as a whole.
  - `integration/` - Integration tests for testing how different parts of the application work together.
- `.env.example` - Example environment variables file.
- `.gitignore` - Specifies files and directories to be ignored by Git.
- `go.mod` - Go module file that defines the project dependencies.
- `go.sum` - Go checksum file that ensures dependency integrity.
- `main.go` - Entry point of the application.

## Getting Started

To get started with this project, follow the instructions below.

### Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/) 1.18 or higher

### Installation

1. **Clone the repository**:

   ```
   git clone https://github.com/MikaSRahwono/initialize-go-project.git
   cd initialize-go-project
   ```

2. **Install the dependencies**:
    ```
   go mod tidy
   ```

3. **Rename the project modules and directory**:
    ```
   ./rename_project.sh github.com/MikaSRahwono/initialize-go-project github.com/your-username/your-project-name
   ```

4. **Copy .env.example**:
    ```
   cp .env.example .env
   ```

### Shout out

The original creators of this architecture is vidu171 in this repo
https://github.com/vidu171/clean-architecture-go.git