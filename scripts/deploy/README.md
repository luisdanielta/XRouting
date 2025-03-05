# 🚀 AWS Lambda Deployment

This guide covers the prerequisites, Lambda function creation, packaging, deployment, and testing for this project.

## 1. Prerequisites

Before deploying, ensure you have the following:

### 1.1 AWS CLI Installed & Configured.

```bash
curl "https://awscli.amazonaws.com/AWSCLIV2.pkg" -o "AWSCLIV2.pkg"
sudo installer -pkg AWSCLIV2.pkg -target /
aws configure
```
Enter your AWS Access Key ID, Secret Access Key, Region. 

IIf you don’t have an AWS user for using credentials, go to the **AWS IAM Console**:

- **Create a new user**  
    - Click **Add user** → Set a username (e.g., `dev-user`).  
    - Choose **Access Key – Programmatic access**.  

- **Assign permissions**  
    - Select **Attach policies directly**.  
    - Add the following policy:  
        - **IAMFullAccess** *(if full access is needed but not recommended)*  
        - **Or create a custom policy** *(see next step).*  

- **Download credentials**  
    - Save the **Access Key ID** and **Secret Access Key**.  


### 1.2. IAM Role with Lambda Execution Permissions.

```bash
aws iam create-role --role-name LambdaExecutionRole --assume-role-policy-document file://trust-policy.json
```
Attach Necessary Policies to the Role

```bash
aws iam attach-role-policy --role-name LambdaExecutionRole --policy-arn arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess
aws iam attach-role-policy --role-name LambdaExecutionRole --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```


## 2. Package the Lambda Function
Run the `deploy_lambda.sh` to install dependencies and package the Lambda function, you can also follow the step by step commands.

## 3. Deploy the Lambda Function

```bash
aws lambda create-function \
  --function-name updateSpaceXData \
  --runtime python3.12 \
  --role arn:aws:iam::YOUR_ACCOUNT_ID:role/LambdaExecutionRole \
  --handler main.lambda_handler \
  --zip-file fileb://lambda_package.zip
```

This command creates an AWS Lambda function named `updateSpaceXData`. It assigns the `LambdaExecutionRole` with the necessary permissions, and sets `main.lambda_handler` as the entry point. Additionally, it uploads the packaged code from `lambda_package.zip`, making the function ready for execution in AWS.

**Get YOUR_ACCOUNT_ID**
```bash
aws sts get-caller-identity --query Account --output text
```

## 4. Test the Lambda Function

**Invoke Lambda with a Test Event**
```bash
aws lambda invoke \
  --function-name updateSpaceXData \
    response.json
cat response.json
```

**View Logs in Real-Time**
```bash
aws logs tail /aws/lambda/updateSpaceXData --follow
```

## 5. Schedule Execution Every 6 Hours

To automatically execute the Lambda function every **6 hours**, we use **Amazon EventBridge (CloudWatch Events)**.

### 5.1 Create an EventBridge Rule

Run the following command to create a rule that triggers every 6 hours:

```bash
aws events put-rule \
    --name lambda-scheduled-execution \
    --schedule-expression "rate(6 hours)"
```
This creates a rule named lambda-scheduled-execution that runs on a fixed interval of 6 hours.

### 5.2 Assign the Rule to Lambda
Finally, associate the rule with the Lambda function:

```bash
aws events put-targets \
    --rule lambda-scheduled-execution \
    --targets "Id"="1","Arn"="arn:aws:lambda:YOUR_ACCOUNT_ID:function:updateSpaceXData"
```

### 5.3 Grant Permissions to EventBridge
Ensure that EventBridge can invoke the Lambda function by adding the correct permissions:

```bash
aws lambda add-permission \
    --function-name updateSpaceXData \
    --statement-id eventbridge-invoke \
    --action "lambda:InvokeFunction" \
    --principal events.amazonaws.com \
    --source-arn arn:aws:events:YOUR_ACCOUNT_ID:rule/lambda-scheduled-execution
```

Replace `YOUR_ACCOUNT_ID` with your actual AWS Account ID when necessary.