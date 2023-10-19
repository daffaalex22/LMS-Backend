![Logo InEdu](https://user-images.githubusercontent.com/74149436/151357103-e850a849-e677-4a84-b43d-15eb994ef6ed.jpeg)

# Contents
- [🎡 Overview](https://github.com/daffaalex22/LMS-Backend/new/main?readme=1#-overview)
  - [🎯 High-Level Architecture](https://github.com/daffaalex22/LMS-Backend/new/main?readme=1#-high-level-architecture)
  - [🌐 Live Preview](https://github.com/daffaalex22/LMS-Backend/new/main?readme=1#-live-preview)
- [🪛 Technical](https://github.com/daffaalex22/LMS-Backend/new/main?readme=1#-technical)
  - [⚙️ Technology and Frameworks
](https://github.com/daffaalex22/LMS-Backend/new/main?readme=1#%EF%B8%8F-technology-and-frameworks)

# 🎡 Overview

InEdu is a web-based platform for enhancing employee skills and knowledge. InEdu aims to improve the abilities of employees who need training where there are videos and learning readings so that employees can learn at their own pace. Each course has interactive slides featuring videos, readings & quizzes.


## 🎯 High-Level Architecture
The following figure demonstrates the tech stacks used on this project on a high-level architecture diagram.
![HLA](https://github.com/daffaalex22/LMS-Backend/assets/67172568/6c756968-a01f-4747-aa6a-4754d04a2e92)

## 🌐 Live Preview
> Notes: Kindly be patient with the latency. The whole full-stack app is deployed using a free service so it might take some time to load.
> If you still have trouble accessing the Live Preview, kindly email me at daffaalex22@gmail.com/daffaalexander.work@gmail.com

🔗This project is accessible at [https://inedu-backend.onrender.com](https://inedu-backend.onrender.com)

### Credential
> 
> **_Teacher Account_**
> 
>   email: alex@teacher.com
> 
>   password: 12345678
> 
> **_Student Account_**
> 
>   email: alex@google.com
> 
>   password: 12345678

# 🪛 Technical

## 📑 API Documentation
#### [ERD](https://www.figma.com/file/LiNhF4XjN2T8ZB9bToe9Qf/Learning-Managament-System?node-id=0%3A1): Entity Relationship Diagram
#### [Postman](https://documenter.getpostman.com/view/17525041/UVeAuoG7): The API Documentation
#### [Swagger](https://app.swaggerhub.com/apis/krisnadwipayana07/InEdu/1.0.0#/students/getLoginStudent): Also API Documentation

## ⚙️ Technology and Frameworks

- [GO](https://go.dev/doc/) as Programming Language
- [echo](https://labstack.com/echo) as high performance framework
- [GORM](https://gorm.io/docs/) for the initial migration and creation of the database schema
- [MySQL](https://dev.mysql.com/doc/)
- Implement custom Error Handling
- Using JWT as a Token via [jwt-go package](https://github.com/dgrijalva/jwt-go)
- Implement Role based authorization
- Write unit test for API endpoint and middlewares
- Using [Docker](https://docs.docker.com/) for packaging applications into containers
- Using [AWS](https://aws.amazon.com/) for Deployment

### 🧹 Clean Architecture

As previously mentioned, the project implements Clean Architecture. The four layers of the project are:

- Domain Layer
- Repository Layer
- Usecase Layer
- Controller Layer

#### The Diagram

![golang clean architecture](https://github.com/daffaalex22/jobdir/raw/main/CleanArch.png)

From the picture above, the four rounded rectangular correspond to each Clean Architecture layer. The slightly bolder arrow pointing from a certain layer to another pictures the dependency of the layer. For example, the domain layer (red) is pointing to repository layer (purple). This means that the repository layer imports package from domain and thus the repository layer depends on the domain layer.

