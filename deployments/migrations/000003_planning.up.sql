CREATE TYPE coordination_action AS ENUM (
    'Coord',
    'Approve',
    'Sign'
);

CREATE TYPE plan_type AS ENUM (
    'OPLAN',
    'CONPLAN',
    'Meeting Plan'
);

CREATE TABLE coordination (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

)