# Go-Rest-Server
## Introduction
Simple HTTP server written in standard library that provides endpoints for getting data for MongoDB and in-memory data storage. 
## Endpoints
API includes: 
- GET /records
  - Fetches data from MongoDB based on filters. Requires JSON with ``startDate``, ``endDate``, ``minCount`` and ``maxCount``
- POST /memory
  - Inserts key-value pair inside in-memory storage. Requires JSON with ``key`` and ``value``
- GET /memory
  - Gets record stored inside in-memory storage. Requires a query parameter ``key``.
