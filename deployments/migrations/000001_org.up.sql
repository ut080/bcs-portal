CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE member_type AS ENUM (
    'SENIOR',
    'CADET SPONSOR',
    'CADET',
    'PATRON'
);

CREATE TYPE grade AS ENUM (
    'Maj Gen',
    'Brig Gen',
    'Col',
    'Lt Col',
    'Maj',
    'Capt',
    '1st Lt',
    '2d Lt',
    'SFO',
    'TFO',
    'FO',
    'CMSgt',
    'SMSgt',
    'MSgt',
    'TSgt',
    'SSgt',
    'SM',
    'C/Col',
    'C/Lt Col',
    'C/Maj',
    'C/Capt',
    'C/1st Lt',
    'C/2d Lt',
    'C/CMSgt',
    'C/SMSgt',
    'C/MSgt',
    'C/TSgt',
    'C/SSgt',
    'C/SrA',
    'C/A1C',
    'C/Amn',
    'CADET'
);

CREATE TABLE member (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    capid INTEGER NOT NULL,
    last_name VARCHAR NOT NULL,
    first_name VARCHAR NOT NULL,
    member_type member_type NOT NULL,
    grade grade NOT NULL
);

CREATE TABLE duty_title (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR NOT NULL,
    member_type member_type NOT NULL,
    min_grade grade,
    max_grade grade,
    INDEX member_type
);

CREATE TABLE duty_assignment (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    office_symbol VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    assistant BOOLEAN DEFAULT false,
    duty_title_id UUID NOT NULL REFERENCES duty_title (id) ON DELETE CASCADE,
    assignee_id UUID REFERENCES member (id) ON DELETE SET NULL
);

CREATE TABLE flight (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL UNIQUE,
    abbreviation VARCHAR NOT NULL,
    flight_commander_id UUID REFERENCES duty_assignment (id) ON DELETE RESTRICT,
    flight_sergeant_id UUID REFERENCES duty_assignment (id) ON DELETE RESTRICT
);

CREATE TABLE element (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL UNIQUE,
    flight_id UUID NOT NULL REFERENCES flight (id) ON DELETE CASCADE,
    element_leader_id UUID NOT NULL REFERENCES duty_assignment (id) ON DELETE RESTRICT,
    asst_element_leader_id UUID NOT NULL REFERENCES duty_assignment (id) ON DELETE RESTRICT
);

CREATE TABLE element_members (
    element_id UUID NOT NULL REFERENCES element (id) ON DELETE CASCADE,
    member_id UUID NOT NULL REFERENCES member (id) ON DELETE CASCADE,
    PRIMARY KEY (element_id, member_id)
);

CREATE TABLE staff_group (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE staff_subgroup (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL UNIQUE,
    staff_group_id UUID REFERENCES staff_group (id) ON DELETE CASCADE,
    leader_id UUID NOT NULL REFERENCES duty_assignment (id) ON DELETE RESTRICT
);

CREATE TABLE staff_subgroup_direct_reports (
    staff_subgroup_id UUID NOT NULL REFERENCES staff_group (id) ON DELETE CASCADE,
    duty_assignment_id UUID NOT NULL REFERENCES duty_assignment (id) ON DELETE CASCADE,
    PRIMARY KEY (staff_subgroup_id, duty_assignment_id)
)