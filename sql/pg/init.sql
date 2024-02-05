
-- CREATE DATABASE test;

CREATE TABLE IF NOT EXISTS public.fi_store (
    id SERIAL PRIMARY KEY,
    previous INTEGER not null,
    current INTEGER not null
);

