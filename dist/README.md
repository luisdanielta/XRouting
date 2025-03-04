# AWS ECS Fargate + ECR Deployment

This is a step-by-step instructions to deploy this application to AWS ECS Fargate, using Amazon ECR as the container image repository. It assumes you are using a `dev-user`, IAM user for previusly configured

---
## Configure AWS IAM Permissions

Verify permissions: Go to IAM > Users > `dev-user` > Permissions, and add these policies if missing.

#### 1. Assign Required Policies to dev-user
  - AmazonEC2ContainerRegistryFullAccess → Full access to ECR.
  - AmazonECS_FullAccess → Full access to ECS services.
  - AWSCloudFormationFullAccess → Required if using CloudFormation for setup.
  - IAMFullAccess → Required for creating execution roles.
---
## Push Docker Image to Amazon ECR

#### 1. Authenticate Docker with ECR
ECR requires authentication before pushing images, replace <AWS_ACCOUNT_ID> with your actual AWS account ID.

```sh
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <AWS_ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com
```

#### 2. Create an ECR Repository
If you haven't already, create an ECR repository:

```sh
aws ecr create-repository --repository-name xrouting-app
```

#### 3. Build & Push the Docker Image

```sh
cd /XRouting/dist/Dockerfile
docker build -t xrouting-app .
docker tag xrouting-app:latest <AWS_ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com/xrouting-app:latest
docker push <AWS_ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com/xrouting-app:latest
```
---
## Configure ECS Fargate Cluster

#### 1. Create an ECS Cluster
  - Open AWS Console and go to ECS.
  - Click Create Cluster.
  - Select Networking Only (AWS Fargate) and click Next.
  - Name your cluster `xrouting-cluster` and click Create.

#### 2. Creating `ecsTaskExecutionRole` for ECS Fargate
  - Go to AWS IAM → Roles → Create role.
  - Select "AWS Service" → Choose Elastic Container Service.
  - Use Case → Select Elastic Container Service Task.
  - Attach Policies
    -  `AmazonECSTaskExecutionRolePolicy`
    -  `CloudWatchLogsFullAccess`
    -  `AmazonDynamoDBFullAccess`
    -  `AmazonEC2ContainerRegistryReadOnly`
  - Set Role Name as `ecsTaskExecutionRole`
  - Click Create Role.
    
#### 3. Create an ECS Task Definition
  - Go to ECS > Task Definitions > Create new task definition.
  - Choose Fargate as the launch type.
  - Set the task name: `xrouting-task.`
  - In Task Role, select `ecsTaskExecutionRole`.
  - In Task Size, set:
    - CPU: 1 vCPU
    - Memory: 2 GB
  - Open Container Card:
    - Container name: `xrouting-container`
    - Image: `<AWS_ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com/xrouting-app:latest`
    - Port mappings, `8000 (backend)`, `4173 (frontend)`
    - Entrypoint:	`sh,-c`
    - Command: `/usr/bin/tini,--,sh,-c,./xrouting\ \&\ npm\ run\ preview`
    - Workdir: `/app`
    - Click Create
   
---
## Create a VPC with Public & Private Subnets

A VPC (Virtual Private Cloud) is needed, so we have to create one before deploying the ECS service. ECS Fargate requires a VPC with subnets and a security group to function properly.

#### 1. Create VPC an more
  - Go to AWS Console → VPC → Your VPCs.
  - Click Create VPC and more.
  - Set:
    - Name tag: `xrouting`
    - IPv4 CIDR block: 10.0.0.0/16
    - Leave IPv6 CIDR block disabled (optional).
   
  With this step AWS will automatically create and attach:
  - Subnet: `xrouting-public-subnet`
  - Subnet: `xrouting-private-subnet`
  - Internet gateway: `xrouting-igw`
  - Route table: `xrouting-route-table`

#### 2. Create a Security Group for ECS
  - Go to AWS Console → EC2 → Security Groups.
  - Click Create security group.
  - Set:
    - Name: `xrouting-sg`
    - Description: `Allow users to use the web app xRouting`
    - VPC: `xrouting-vpc`
- Allow Traffic to ECS Services
- Add rules:
  - Rule 1: `Type: Custom TCP`, `Port Range: 8000`, `Source: 0.0.0.0/0`
  - Rule 2: `Type: Custom TCP`, `Port Range: 4173`, `Source: 0.0.0.0/0`
- Click Create security groups.

---
##  Deploy ECS Service

#### 1. Create an ECS Service
  - Go to ECS > Clusters > xrouting-cluster.
  - Click Create Service.
  - Select Launch Type: Fargate.
  - Choose:
    - Task Definition: `xrouting-task`
    - Cluster: `xrouting-cluster`
    - Service Name: `xrouting-service`
    - Number of tasks: 1
  - In Networking:
    - Select `xrouting-vpc`.
    - Enable Auto-Assign Public IP.
    - Click Next and deploy the service.
