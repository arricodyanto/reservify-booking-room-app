CREATE DATABASE booking_room_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE role_type AS ENUM ('employee', 'admin', 'ga');

CREATE TABLE employees (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50),
    username VARCHAR(50) UNIQUE,
    password VARCHAR(200),
    division VARCHAR(50),
    position VARCHAR(50),
    role role_type DEFAULT 'employee',
    contact VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);


CREATE TABLE facilities (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name     VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    quantity INT
);

CREATE TYPE status_type AS ENUM ('available', 'booked', 'unavailable' );

CREATE TABLE rooms (
    id   uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name      VARCHAR(100),
    room_type VARCHAR(100),
    capacity  INT,
    status status_type DEFAULT 'available', -- 'available', 'booked', 'unavailable'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE trx_room_facility (
    id uuid DEFAULT uuid_generate_v4() UNIQUE,
    room_id         uuid NOT NULL,
    facility_id     uuid NOT NULL,
    quantity        INT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_composite PRIMARY KEY (column1, column2)
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TYPE transaction_status AS ENUM ('pending', 'accepted', 'declined');

CREATE TABLE transactions (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    employee_id uuid,
    room_id uuid,
    description TEXT,
    status transaction_status DEFAULT 'pending', -- 'pending', 'accepted', 'declined'
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);