CREATE DATABASE booking_room_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE employees (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50),
    username VARCHAR(50),
    password VARCHAR(50),
    division VARCHAR(50),
    position VARCHAR(50),
    role VARCHAR(50),
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
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    room_id         uuid NOT NULL,
    facility_id     uuid NOT NULL,
    quantity        INT,
    -- status VARCHAR(10) DEFAULT 'used', -- 'used', 'returned'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
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