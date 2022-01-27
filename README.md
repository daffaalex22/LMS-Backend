## BACKEND RESTFULL API INEDU MANAGAMENT LEARNING SYSTEM WITH CLEAN ARCHITECTURE

![Logo InEdu](https://user-images.githubusercontent.com/74149436/151357103-e850a849-e677-4a84-b43d-15eb994ef6ed.jpeg)

### Description

InEdu is a web-based platform for enhancing employee skills and knowledge. InEdu aims to improve the abilities of employees who need training where there are videos and learning readings so that employees can learn at their own pace. Each course has interactive slides featuring videos, readings & quizzes.

## API Documentation

##### ERD

<a href="https://www.figma.com/file/LiNhF4XjN2T8ZB9bToe9Qf/Learning-Managament-System?node-id=0%3A1">Entity Relationship Diagram - InEdu Learning Management System</a>

#### Swagger

<a href="https://app.swaggerhub.com/apis/krisnadwipayana07/InEdu/1.0.0#/students/getLoginStudent">InEdu Learning Management System API</a>

#### Postman

<a href="https://documenter.getpostman.com/view/17525041/UVeAuoG7">InEdu Learning Management System API</a>

### Technology and Frameworks yang digunakan

- [GO](https://go.dev/doc/) as Programming Language
- [echo](https://labstack.com/echo) as high performance framework
- [GORM](https://gorm.io/docs/) for the initial migration and creation of the database schema
- [MySQL](https://dev.mysql.com/doc/)
- Implement custom Error Handling
- Using JWT as Token via [jwt-go package](https://github.com/dgrijalva/jwt-go)
- Implement Role base authorization
- Write unit test for API endpoint and middlewares
- Using [Docker](https://docs.docker.com/) for packaging applications into containers
- Using [AWS](https://aws.amazon.com/) for Deployment

## Clean Architecture

As previously mentioned, the project implements Clean Architecture. The four layers on the project are:

- Domain Layer
- Repository Layer
- Usecase Layer
- Controller Layer

### The Diagram

![golang clean architecture](https://github.com/daffaalex22/jobdir/raw/main/CleanArch.png)

From the picture above, the four rounded rectangular corresponds to each Clean Architecture layer. The slightly bolder arrow pointing from certain layer to another, pictures the dependency of the layer. For example, the domain layer (red) is pointing to repository layer (purple). This means that the repository layer imports package from domain and thus the repository layer depends on the domain layer.

## How to use from this API

##### Clone the repository

```
git clone https://github.com/Kelompok-8-alterra/LMS-Backend.git
cd backend-InEdu
```
