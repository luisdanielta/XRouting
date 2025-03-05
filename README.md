# 🚀 XRouting Maintenance & Analytics App

This project is a **Go-based application** that processes SpaceX data, generates maintenance tasks, and provides analytics using a **microservices approach**. The **Python component** is the AWS Lambda function responsible for data ingestion and analytics generation.

## Business Logic Overview

- The system **automatically processes SpaceX API data** to track components like **rockets, ships, capsules, and cores**.
- **Every 6 hours**, an **AWS Lambda function** fetches and updates SpaceX data on our DB
- **Maintenance tasks** are automatically generated when a failure occurs in a launch.
- **Analytics are computed** from the data and stored in a dedicated `analytics` table for visualization.
- **Users can log maintenance actions** and track historical data.
- The web application provides an **interactive dashboard** to view components, maintenance records, and analytics.

---

## Project Setup

### 1️. Clone the Repository
```sh
git https://github.com/luisdanielta/XRouting.git
cd XRouting
```

### 2. Recommended: Use DevContainer (VS Code)
This project is fully containerized using DevContainer, ensuring a consistent development environment.

To get started:
- Open the project in VS Code.
- Select "Reopen in Container".
- Let DevContainer build and initialize the environment.

This will set up:

- The Go backend.
- The React frontend.
- The Python Lambda function.

⚠️ If you prefer a manual setup, refer to the specific README files inside each module.

---

## Documentation Structure
Each module has its own README with detailed instructions:

- Lambda Function → [Lambda Function Overview](scripts/README.md)
- DB setup  →  [DynamoDB config](scripts/adapters/db/README.md)
- Deployment Setup → [Lambda Deploy](scripts/deploy/README.md) / [Web Deploy](dist/README.md)  
- Web App (Go/React) → [Web App](web/README.md)

For more details on how each component works, refer to the respective module documentation.

## Run Locally with Docker
To run the entire system locally using Docker, follow these steps:

- Complete the setup in [DynamoDB config](scripts/adapters/db/README.md) to create a development user with the required permissions for Lambda and DynamoDB.
- Extract the necessary AWS credentials for execution.
- Run the migration scripts to create tables and insert initial data.
- Once these steps are done, you can start the application with:
```sh
docker run --rm -it -p 8000:8000 -p 4173:4173 \
  -e AWS_ACCESS_KEY_ID="YOUR_ACCESS_KEY" \
  -e AWS_SECRET_ACCESS_KEY="YOUR_SECRET_ACCESS_KEY" \
  -e AWS_REGION="YOUR_REGION" \
  taluisdaniel/xrouting

```
⚠️ **Note**: If the migration is not executed, the database will be empty, and no data will be available in the app.
