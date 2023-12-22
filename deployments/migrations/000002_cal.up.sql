CREATE TYPE uniform AS ENUM (
    'Semi-formal',
    'Service Dress',
    'Service',
    'Utility',
    'Field',
    'PT',
    'Civilian'
);


CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR NOT NULL,
    start_datetime TIMESTAMP NOT NULL,
    end_datetime TIMESTAMP NOT NULL,
    uod uniform,
    poc VARCHAR
);

CREATE TABLE meetings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_id UUID NOT NULL REFERENCES events (id),
    topic VARCHAR NOT NULL
);
