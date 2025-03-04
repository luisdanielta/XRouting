# Lamdba Function: SpaceX Data Extraction

This function extracts, transforms, and stores data from the SpaceX API into DynamoDB using a **hexagonal architecture**.

## 1. Development Environment Setup

### Recommended: `.devcontainer`
For a fully pre-configured development environment, use **VS Code** and select **"Reopen in Container"**. The project includes a `.devcontainer` and a `Dockerfile` to streamline setup.

### Manual Setup
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

## 3. Lambda Function Overview
The Lambda function runs an ETL (Extract, Transform, Load) process every 6 hours, fetching `SpaceX` data and storing it in DynamoDB. The main workflow is:

- Extract Data from the SpaceX API.
- Transform Data into structured entities:
- Convert raw API data into `Components`, `SubComponents`, `Comments`, and `Analytical Metrics`.
- Generate custom analytics based on launches and hardware performance.
- Load Data into DynamoDB using dedicated repositories.

## 4. Code Structure
The Lambda function is implemented in a single script and follows this core flow:

```python
def lambda_handler(event: Dict[str, str], context: object) -> Dict[str, str | int]:
    spaceXData = fetchSpaceXData()
    components, comments, subComponents, analytics = transformData(spaceXData)
    return saveToDynamoDB(components, comments, subComponents, analytics)
```

## 5. Running the Code Locally 🏃

**Prerequisites**
    - Python enviroment configured (with previous steps)
    - Your AWS user is authenticated and has the necessary permissions.
    - You have configured the required AWS IAM policies for your user:
        - AmazonDynamoDBFullAccess
        - AWSLambdaBasicExecutionRole

**1. Create DynamoDB Tables**
Before running the Lambda function, you need to create the necessary DynamoDB tables. This can be done by running the migration script:
```python
python migration/create_tables.py
```

**2. Run the Lambda Function Locally**
Once the tables are set up, you can run the Lambda function directly, `python main.py`. This will:

    - Fetch data from the SpaceX API.
    - Process and transform the data.
    - Insert it into DynamoDB.
    - Return a JSON response with the status code.

After execution, you should see a response similar to:

```python
{
    "statusCode": 200,
    "body": {
        "message": "Data successfully saved to DynamoDB.",
        "componentsSaved": 150,
        "subComponentsSaved": 75,
        "commentsSaved": 200,
        "analyticsUpdated": 10
    }
}
```

If any issue occurs, check your AWS authentication and DynamoDB setup.

This version ensures that you have clear steps to:
1. Authenticate with AWS.
2. Create the necessary DynamoDB tables.
3. Execute the Lambda function locally.

### To read about the deploymeny docs, please go to [Lambda Deployment Docs](deploy/README.md)

