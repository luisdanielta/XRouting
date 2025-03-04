# 🚀 AWS Lambda Deployment

This guide covers the prerequisites, Lambda function creation, packaging, deployment, and testing for this project.

## 1. Prerequisites

Before deploying, ensure you have the following:

#### 1. AWS CLI Installed & Configured.

```bash
curl "https://awscli.amazonaws.com/AWSCLIV2.pkg" -o "AWSCLIV2.pkg"
sudo installer -pkg AWSCLIV2.pkg -target /
aws configure
```
Enter your AWS Access Key ID, Secret Access Key, Region. 

If you don’t have an AWS user for using credentials, go to the AWS IAM Console:

Create a new user
- Click Add user → Set a username (e.g., dev-user).
- Choose Access Key – Programmatic access.

Assign permissions:
- Select Attach policies directly.
- Add the following policy: IAMFullAccess (if full access is needed but I personally do not recommend) or create a custom policy (see next step).

Download credentials:
- Save the Access Key ID and Secret Access Key.

#### 2. IAM Role with Lambda Execution Permissions.

```bash
aws iam create-role --role-name LambdaExecutionRole --assume-role-policy-document file://trust-policy.json
```
Attach Necessary Policies to the Role

```bash
aws iam attach-role-policy --role-name LambdaExecutionRole --policy-arn arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess
aws iam attach-role-policy --role-name LambdaExecutionRole --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

#### 3. Package the Lambda Function
Run the `deploy_lambda.sh` to install dependencies and package the Lambda function, you can also follow the step by step commands.

#### 4. Deploy the Lambda Function

```bash
aws lambda create-function \
  --function-name updateSpaceXData \
  --runtime python3.12 \
  --role arn:aws:iam::YOUR_ACCOUNT_ID:role/LambdaExecutionRole \
  --handler main.lambda_handler \
  --zip-file fileb://lambda_package.zip
```

This command creates an AWS Lambda function named updateSpaceXData. It assigns the LambdaExecutionRole with the necessary permissions, and sets main.lambda_handler as the entry point. Additionally, it uploads the packaged code from lambda_package.zip, making the function ready for execution in AWS.

- YOUR_ACCOUNT_ID
```bash
aws sts get-caller-identity --query Account --output text
```



