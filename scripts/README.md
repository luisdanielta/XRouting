# Lamdba Function: SpaceX Data Extraction

This function extracts, transforms, and stores data from the SpaceX API into DynamoDB using a **hexagonal architecture**.

## 1. Development Environment Setup

### 1️.1 Recommended: `.devcontainer`
For a fully pre-configured development environment, use **VS Code** and select **"Reopen in Container"**. The project includes a `.devcontainer` and a `Dockerfile` to streamline setup.

### 1.2️ Manual Setup
If you prefer to run the code locally, ensure you have **Python 3.11+** installed and install dependencies:

```sh
pip install -r requirements.txt
```
## 2. Project Architecture
This project follows a hexagonal architecture, separating concerns into distinct modules:

- `adapters/` → Interfaces with external services (SpaceX API, DynamoDB)
- `core/` → Business logic and domain entities
- `ports/` → Services connecting business logic to the infrastructure
- `deploy/` → Deployment scripts
- `migration/` → Data migration scripts
