CREATE TABLE channels
(
    type        text NOT NULL,
    CONSTRAINT channels_type_unique PRIMARY KEY (type)
);

