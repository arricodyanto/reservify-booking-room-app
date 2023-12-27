# Resevify Aplication

## Prerequisites

Before running the Resevify application, make sure you have fulfilled the following prerequisites:

- Go (Golang) is installed on your system.
- PostgreSQL is installed, and you have created the tables as specified in the `ddl.sql` file. Then, insert the table contents from the `dml.sql` file as dummy data.
- An active internet connection is required to download Go dependencies.

## Running the Application

Once the application is running, you can access it through a web browser or use it through an API client such as Postman or cURL. Then, you can log in using an account created by the admin. This application provides APIs for managing Rooms, Facilities, Employees, and Transactions.

## Using the API

Below are instructions on how to use the API based on the features provided by the Resevify application:

### API Spec

#### Login API

Request :

- Method : `POST`
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "username": "string",
  "password": "string"
}
```

#### Employee API

##### Create Employee

Request :

- Method : POST
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
  "status": {
    "code": 201,
    "message": "Created"
  },
  "data": {
    "id": "string",
    "name": "string",
    "username": "string",
    "password": "string",
    "role": "string",
    "division": "string",
    "position": "string",
    "contact": "string",
    "createdAt": "2000-01-01T12:00:00Z", (curent time)
    "updatedAt": "2000-01-01T12:00:00Z"  (curent time)
  }
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "name": "string",
  "username": "string",
  "password": "string",
  "role": "string",
  "division": "string",
  "position": "string",
  "contact": "string"
}
```

##### Get Employees

Request :

- Method : GET
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Query Param :
  - page : int `optional`
  - size : int `optional`
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data": [
        {
            "id": "string",
            "name": "string",
            "username": "string",
            "password": "string",
            "role": "string",
            "division": "string",
            "position": "string",
            "contact": "string",
            "createdAt": "2000-01-01T12:00:00Z", (curent time)
            "updatedAt": "2000-01-01T12:00:00Z"  (curent time)
        }
    ],
    "paging": {
        "page": 1,          (default value)
        "rowsPerPage": 5,   (default value)
        "totalRows": int,
        "totalPages": int
    }
}

```

##### Get By Employee Id

Request :

- Method : GET
- Endpoint : `/employees/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
  "status": {
    "code": 200,
    "message": "Ok"
  },
  "data": {
    "id": "string",
    "name": "string",
    "username": "string",
    "password": "string",
    "role": "string",
    "division": "string",
    "position": "string",
    "contact": "string",
    "createdAt": "2000-01-01T00:00:00Z",
    "updatedAt": "2000-01-01T00:00:00Z"
  }
}
```

##### Get By Employee Usename

Request :

- Method : GET
- Endpoint : `/employees/username/:username`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
  "status": {
    "code": 200,
    "message": "Ok"
  },
  "data": {
    "id": "string",
    "name": "string",
    "username": "string",
    "password": "string",
    "role": "string",
    "division": "string",
    "position": "string",
    "contact": "string",
    "createdAt": "2000-01-01T00:00:00Z",
    "updatedAt": "2000-01-01T00:00:00Z"
  }
}
```

##### Update Employee

Request :

- Method : PUT
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

```json
{
  "id": "string",
  "name": "string",
  "username": "string",
  "password": "string",
  "role": "string",
  "division": "string",
  "position": "string",
  "contact": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
  "status": {
    "code": 200,
    "message": "Updated Successfully"
  },
  "data": {
    "id": "string",
    "name": "string",
    "username": "string",
    "password": "string",
    "role": "string",
    "division": "string",
    "position": "string",
    "contact": "string",
    "createdAt": "2000-01-01T00:00:00Z",
    "updatedAt": "2000-01-01T00:00:00Z"
  }
}
```

#### Facility API

##### Create Facility

Request :

- Method : POST
- Endpoint : `/facilities`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
    "name": "string",
    "quantity": int
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "status": {
        "code": 201,
        "message": "Created"
    },
    "data": {
        "id": "string",
        "name": "string",
        "quantity": int,
        "createdAt": "2000-01-01T00:00:00Z",
        "updatedAt": "2000-01-01T00:00:00Z"
    }
}
```

##### Get Facilities

Request :

- Method : GET
- Endpoint : `/facilities`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Query Param :
  - page : int `optional`
  - size : int `optional`

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "ok"
    },
    "data": [
        {
            "id": "string",
            "name": "string",
            "quantity": int,
            "createdAt": "2000-01-01T00:00:00Z",
            "updatedAt": "2000-01-01T00:00:00Z"
        }
    ],
    "paging": {
        "page": 1, (default value)
        "rowsPerPage": 5, (default value)
        "totalRows": int,
        "totalPages": int
    }
}
```

##### Get Facility By Id

Request :

- Method : GET
- Endpoint : `/facilities/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "ok"
    },
    "data": {
        "id": "string",
        "name": "string",
        "quantity": int,
        "createdAt": "2000-01-01T00:00:00Z",
        "updatedAt": "2000-01-01T00:00:00Z"
    }
}

```

##### Update Facility By Id

Request :

- Method : GET
- Endpoint : `/facilities/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

```json
{
    "id": "string",
    "name": "string",
    "quantity": int
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Updated"
    },
    "data": {
        "id": "string",
        "name": "string",
        "quantity": int,
        "createdAt": "2000-01-01T00:00:00Z",
        "updatedAt": "2000-01-01T00:00:00Z" (curent time)
    }
}

```

#### Room API

##### Create Room

Request :

- Method : POST
- Endpoint : `/rooms`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
    "name": "string",
    "room_type": "string",
    "capacity": int,
    "status": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "status": {
        "code": 201,
        "message": "Created"
    },
    "data": {
        "id": "string",
        "name": "string",
        "room_type": "string",
        "capacity": int,
        "status": "string",
        "createdAt": "2000-01-01T12:00:00Z", (curent time)
        "updatedAt": "2000-01-01T12:00:00Z"  (curent time)
    }
}
```
