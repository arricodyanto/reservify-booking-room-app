CREATE DATABASE booking_room_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE employees (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50),
    division VARCHAR(50),
    position VARCHAR(50),
    contact VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE facilities (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name     VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
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
    ID INT PRIMARY KEY,
    employe_id uuid,
    room_id uuid,
    decription TEXT,
    status transaction_status DEFAULT 'pending', -- 'pending', 'accepted', 'declined'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (employe_id) REFERENCES employees(id),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);