# Testing Golang and building an image from ECS, ECR, EC2 with nginx

## Steps

### Build Docker image
```shell
docker build -t golang-app:1.0 .
# If you are running on Mac OS and M1 micro, you need to build with `buildx`
docker buildx build --platform=linux/amd64 -t golang-app:1.0 .
```

### Create repository
```shell
aws ecr create-repository --repository-name golang-app:1.0 --region ap-southeast-2
```

### Link local image with remote image
```shell
docker tag golang-app:1.0 <YOUR ACCOUNT NUMBER>.dkr.ecr.ap-southeast-2.amazonaws.com/golang-app:1.0
```

### Login to Docker
```shell
aws ecr get-login-password --region ap-southeast-2 | docker login --username AWS --password-stdin <YOUR ACCOUNT NUMBER>.dkr.ecr.ap-southeast-2.amazonaws.com/golang-app
```

### Push to ECR
```shell
docker push <YOUR ACCOUNT NUMBER>.dkr.ecr.ap-southeast-2.amazonaws.com/golang-app:1.0
```

### ECS steps
1. Create cluster ECS
2. Create Task Definition
3. Run the ECS Task! Deploy --> Run Task

## Source

* (How to create a Docker Image with Nginx from an EC2 Instance and Push to ECR Repository)[https://dev.to/aws-builders/how-to-create-a-docker-image-with-nginx-from-an-ec2-instance-and-push-to-ecr-repository-nkb]