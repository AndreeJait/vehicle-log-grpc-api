-- +migrate Up
CREATE TABLE public.types
(
    id   varchar(36) PRIMARY KEY,
    name varchar not null unique
);

CREATE TABLE public.vehicles
(
    id          varchar(36) NOT NULL PRIMARY KEY,
    plat_number varchar     NOT NULL UNIQUE,
    type_id     varchar(36) NOT NULL,
    CONSTRAINT fk_type_vehicle FOREIGN KEY (type_id) REFERENCES types (id)
);

CREATE TABLE public.towns
(
    id        varchar(36) NOT NULL PRIMARY KEY,
    town_code varchar     not null,
    town_name varchar     not null
);

CREATE TABLE public.town_slots
(
    id       varchar(36) PRIMARY KEY,
    town_id  varchar(36) not null,
    type_id  varchar(36) not null,
    quantity bigint      not null,
    price    bigint      not null,
    CONSTRAINT fk_town_types FOREIGN KEY (town_id) REFERENCES towns (id),
    CONSTRAINT fk_type_towns FOREIGN KEY (type_id) REFERENCES types (id)
);

CREATE TABLE public.logs
(
    id            varchar PRIMARY KEY,
    town_slots_id varchar(36) NOT NULL,
    vehicle_id    varchar(36) NOT NULL,
    time_in       varchar        NOT NULL,
    time_out      varchar,
    date_at       varchar        NOT NULL,
    date_out_at   varchar,
    status        varchar     NOT NULL,
    CONSTRAINT fk_log_town_slots FOREIGN KEY (town_slots_id) REFERENCES town_slots (id)
);
-- +migrate Down
DROP TABLE public.logs;
DROP TABLE public.town_slots;
DROP TABLE public.towns;
DROP TABLE public.vehicles;
DROP TABLE public.types;