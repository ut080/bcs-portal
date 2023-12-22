CREATE TYPE coordination_action AS ENUM (
    'COORD',
    'APPROVE',
    'SIGN',
    'ACTION'
);

CREATE TYPE plan_type AS ENUM (
    'OPLAN',
    'CONPLAN',
    'Meeting Plan'
);

CREATE TABLE plans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    plan_type plan_type NOT NULL,
    plan_number VARCHAR NOT NULL UNIQUE,
    title VARCHAR NOT NULL,
    planning_start TIMESTAMP,
    planning_due TIMESTAMP,
    event_start TIMESTAMP,
    event_end TIMESTAMP,
    project_officer_id UUID REFERENCES members (id),
    cadet_project_officer_id UUID REFERENCES members (id)
);

CREATE TABLE coordination (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    plan_id UUID NOT NULL REFERENCES plans (id),
    coord_order SMALLINT NOT NULL CHECK (coord_order > 0),
    office_id UUID NOT NULL REFERENCES duty_assignments (id),
    action coordination_action NOT NULL DEFAULT 'COORD',
    completed TIMESTAMP,
    outcome VARCHAR NOT NULL
);

CREATE TABLE plan_sections (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    plan_id UUID NOT NULL REFERENCES plans (id),
    section_number SMALLINT NOT NULL CHECK (section_number > 0),
    title VARCHAR NOT NULL,
    body TEXT NOT NULL
);

CREATE TABLE lesson_plans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR NOT NULL,
    duration SMALLINT NOT NULL CHECK (duration >= 0),
    objectives VARCHAR NOT NULL,
    resources VARCHAR NOT NULL,
    outline VARCHAR NOT NULL
);

CREATE TABLE training_blocks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    plan_id UUID NOT NULL REFERENCES plans (id),
    block_number INT NOT NULL CHECK (block_number > 0),
    opr UUID NOT NULL REFERENCES duty_assignments (id),
    cdt_opr UUID NOT NULL REFERENCES duty_assignments (id),
    topic VARCHAR NOT NULL,
    instructor UUID NOT NULL REFERENCES members (id),
    lesson_plan_id UUID REFERENCES lesson_plans (id)
);
