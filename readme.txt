Instruction for newcomers
There are 5 main packages in the project:
- Common: Package common provides utility functions and provides initialization logic for the application.
- Controllers: Package controllers provides HTTP handler functions for the application. In addition, this package will be responsible for creating transactions with using Kubernetes Job controller
- Store: Package store provides persistence logic with PostgreSQL database.
- Model: Package model describes the data model of the application.
- Routers: Package routers implements HTTP request routers for the REST API.

To add new functionality to Inception service you need to do next:
1. Define your model in model package if needed. Note: your model should not contradict to existing database schema
    (please see confluence for more details)
2. Add


Useful tools which you need to install:
1. Gorilla mux - github.com/gorilla/mux
2. Glide - package manger for golang
3.