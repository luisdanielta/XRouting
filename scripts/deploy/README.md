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




