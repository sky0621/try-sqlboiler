
-- +migrate Up
CREATE TABLE customer (
  id bigserial NOT NULL,
  name varchar(64) NOT NULL,
  age int NOT NULL,
  prefectures varchar(32) NOT NULL,
  created_by varchar(256),
  created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  updated_by varchar(256),
  updated_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE customer;
