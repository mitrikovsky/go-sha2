CREATE SCHEMA hash_data
  CREATE TABLE jobs
  (
    id              serial PRIMARY KEY NOT NULL,
    payload         varchar(255)       NOT NULL DEFAULT '',
    hash_rounds_cnt int                NOT NULL DEFAULT 0,
    status          smallint           NOT NULL DEFAULT 0,
    hash            varchar(255)       NOT NULL DEFAULT ''
  );
