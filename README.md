# EFV-ICBC API

A high-performance Go HTTP API for Electric Vehicle (EV) integration with the Insurance Corporation of British Columbia (ICBC). This service enables vehicle verification, driver validation, and insurance status checking with production-ready OpenShift deployment support.

**Version**: 0.1.7 | **Go**: 1.20+ | **Status**: Active Development

---

## Quick Links

### 📚 API Documentation (Swagger)

| Environment | Link |
|-------------|------|
| **Local Development** | [http://localhost:8080/swagger/](http://localhost:8080/swagger/) |
| **Development** | [https://efv-icbc-api-17db4f-dev.apps.silver.devops.gov.bc.ca/swagger/](https://efv-icbc-api-17db4f-dev.apps.silver.devops.gov.bc.ca/swagger/) |

---

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Configuration](#configuration)
- [Deployment](#deployment)
- [Contributing](#contributing)

---

## Overview

The EFV-ICBC API provides RESTful endpoints for:

- 🔍 **Driver Validation** - Verify driver license status against BC registry
- 🚗 **Vehicle Verification** - Confirm vehicle registration using VIN and plate
- 📋 **Insurance Status** - Check real-time insurance policy information
- 📖 **API Documentation** - Interactive Swagger UI and OpenAPI 3.0 spec

**Current Implementation**: Mock endpoints for development and testing

---

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Docker (optional, for containerized deployment)

### Run Locally

**macOS / Linux:**
```bash
git clone https://github.com/bcgov/efv-icbc-api.git
cd efv-icbc-api
export PORT=8080
go run .