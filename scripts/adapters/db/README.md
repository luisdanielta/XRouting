# DynamoDB Setup & AWS User Configuration

This document explains how to **set up an AWS IAM user**, grant the necessary permissions, and create the required DynamoDB tables. If you already have a user with the required permisisions, just go to Step 3.

---

### Step 1: Create the AWS IAM User (`dev-user`)

- Go to the **AWS Management Console** and open the **IAM (Identity and Access Management)** service.
- Navigate to **Users** > **Add users**.
- Set the **User name** as: `dev-user`
- **Select "Access key - Programmatic access"** under **Credential type** (this is required for `boto3`).
- Click **Next** to configure permissions.

---

###  Step 2: Assign Permissions to `dev-user`

- Select **"Attach existing policies directly"**.
- Search for and attach the following policies:
  - `AmazonDynamoDBFullAccess` → Full access to DynamoDB.
  - `AWSLambdaBasicExecutionRole` → Required for integration with AWS Lambda.
- Click **Next**, then **Create User**.

Once created, **download the access credentials** (`Access Key ID` and `Secret Access Key`), as they will be used to authenticate with `boto3`.

---

###  Step 3: Configure AWS Credentials Locally

To use `dev-user` for creating DynamoDB tables, configure your AWS credentials:

```sh
aws configure
```
Enter the following when prompted:

- AWS Access Key ID: (from dev-user)
- AWS Secret Access Key: (from dev-user)
- Default region name: us-east-1
  
---
## Step 4: Create DynamoDB Tables

DynamoDB tables are managed using the DynamoDBManager class, which leverages the official AWS SDK for Python (boto3) to interact with DynamoDB.

This class is responsible for:

- Checking if the table exists before attempting to create it.
- Creating the table if it does not exist, using the specified primary key schema.
- Waiting for the table to become fully available before proceeding.

#### **How It Works**

The `DynamoDBManager` class initializes an AWS DynamoDB client using boto3, which requires valid AWS credentials. The credentials are provided through aws configure or environment variables.
- `boto3.client("dynamodb")` creates a low-level connection to DynamoDB.
- The AWS Access Key ID and Secret Access Key configured in aws configure are required to authenticate requests.
- The region (us-east-1) ensures the database is created in the correct AWS data center.

--- 
#### **Primary Key Setup**

DynamoDB tables need a primary key for uniquely identifying each item. This includes:

- Partition Key (Primary Key) – Uniquely identifies each record.
- (Optional) Sort Key – Used for composite primary keys.
- The table creation method in DynamoDBManager defines these keys dynamically using the schemas defined on the [Entities](../../core/entities/). 

Each Entity represents a table on DynamoDB

---
## Step 5: Migration

#### **How to Create the Tables**

To execute the table creation process, simply run:

```python
python scripts/migration/create_tables.py
```
---
#### **How to Migrate Data**

```python
python scripts/migration/generateUsers.py
python scripts/migration/insert.py
python scripts/migration/insertAnalitics.py
```
Running the Lambda will also populate data on DynamoDB
```python
python scripts/main.py
```
