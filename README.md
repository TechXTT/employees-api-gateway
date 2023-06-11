## Description

This repository is part of a combined project for Databases and Cloud Technologies classes @ TUES, Sofia<br>
It serves as a HTTP API Gateway for the Employees project<br>
Uses gRPC to communicate with the Authentication and Data services

## Repositories

- [API Gateway](https://github.com/TechXTT/employees-api-gateway) - HTTP API Gateway
- [Authentication Service](https://github.com/TechXTT/employees-auth-svc) - gRPC Authentication Service
- [Data Service](https://github.com/TechXTT/employees-data-svc) - gRPC Data Service

## Prerquisites

Before you can use this project, you need to have the following software installed:

- Docker
- Kubernetes
- Nginx Ingress

## Usage

```bash
$ kubectl apply -f combined-deployment-1.0.yml
```
