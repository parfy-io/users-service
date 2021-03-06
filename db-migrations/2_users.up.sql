CREATE TABLE users
(
    id        serial NOT NULL,
    full_name text   NOT NULL,
    names     text[] NOT NULL,
    client_id text   NOT NULL,
    CONSTRAINT users_id_unique PRIMARY KEY (id),
    CONSTRAINT users_clients_fkey FOREIGN KEY (client_id) REFERENCES clients (id)
);
