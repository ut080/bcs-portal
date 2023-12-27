CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE member_type AS ENUM (
    'SENIOR',
    'CADET',
    'CADET SPONSOR',
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

CREATE TABLE members (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    capid INTEGER NOT NULL CHECK (capid = 0 OR capid >= 100000),
    last_name VARCHAR NOT NULL,
    first_name VARCHAR NOT NULL,
    member_type member_type NOT NULL,
    grade grade NOT NULL,
    join_date DATE,
    rank_date DATE,
    expiration_date DATE
);

CREATE TABLE duty_titles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR NOT NULL UNIQUE,
    title VARCHAR NOT NULL,
    member_type member_type NOT NULL,
    min_grade grade,
    max_grade grade,
    INDEX member_type
);

CREATE TABLE duty_assignments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    office_symbol VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    assistant BOOLEAN DEFAULT false,
    duty_title_id UUID NOT NULL REFERENCES duty_titles (id) ON DELETE CASCADE,
    assignee_id UUID REFERENCES members (id) ON DELETE SET NULL
);

CREATE TABLE flights (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL UNIQUE,
    abbreviation VARCHAR NOT NULL,
    flight_commander_id UUID REFERENCES duty_assignments (id) ON DELETE RESTRICT,
    flight_sergeant_id UUID REFERENCES duty_assignments (id) ON DELETE RESTRICT
);

CREATE TABLE elements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL UNIQUE,
    flight_id UUID NOT NULL REFERENCES flights (id) ON DELETE CASCADE,
    element_leader_id UUID NOT NULL REFERENCES duty_assignments (id) ON DELETE RESTRICT,
    asst_element_leader_id UUID NOT NULL REFERENCES duty_assignments (id) ON DELETE RESTRICT
);

CREATE TABLE element_members (
    element_id UUID NOT NULL REFERENCES elements (id) ON DELETE CASCADE,
    member_id UUID NOT NULL REFERENCES members (id) ON DELETE CASCADE,
    PRIMARY KEY (element_id, member_id)
);

CREATE TABLE staff_groups (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE staff_subgroups (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL UNIQUE,
    staff_group_id UUID REFERENCES staff_groups (id) ON DELETE CASCADE,
    leader_id UUID NOT NULL REFERENCES duty_assignments (id) ON DELETE RESTRICT
);

CREATE TABLE staff_subgroup_direct_reports (
    staff_subgroup_id UUID NOT NULL REFERENCES staff_groups (id) ON DELETE CASCADE,
    duty_assignment_id UUID NOT NULL REFERENCES duty_assignments (id) ON DELETE CASCADE,
    PRIMARY KEY (staff_subgroup_id, duty_assignment_id)
);

INSERT INTO duty_titles (code, title, member_type, min_grade, max_grade) VALUES
    ('SM-ACTIVITY', 'Activities Officer', 'SENIOR', null, null),
    ('SM-ADMIN', 'Administrative Officer', 'SENIOR', null, null),
    ('SM-CC-ADVISOR', 'Advisor to the Commander', 'SENIOR', null, null),
    ('SM-AE', 'Aerospace Education Officer', 'SENIOR', null, null),
    ('SM-ALERT', 'Alerting Officer', 'SENIOR', null, null),
    ('SM-CC', 'Commander', 'SENIOR', null, null),
    ('SM-CHAPLAIN', 'Chaplain', 'SENIOR', null, null),
    ('SM-CDI', 'Character Development Instructor', 'SENIOR', null, null),
    ('SM-COMMUNICATIONS', 'Communications Officer', 'SENIOR', null, null),
    ('SM-CYBER', 'Cyber Education Officer', 'SENIOR', null, null),
    ('SM-DEPUTY', 'Deputy Commander', 'SENIOR', null, null),
    ('SM-DCC', 'Deputy Commander for Cadets', 'SENIOR', null, null),
    ('SM-DCS', 'Deputy Commander for Seniors', 'SENIOR', null, null),
    ('SM-PREP', 'Disaster Preparedness Officer', 'SENIOR', null, null),
    ('SM-DDR', 'Drug Demand Reduction Officer', 'SENIOR', null, null),
    ('SM-ETO', 'Education and Training Officer', 'SENIOR', null, null),
    ('SM-ES', 'Emergency Services Officer', 'SENIOR', null, null),
    ('SM-EST', 'Emergency Services Training Officer', 'SENIOR', null, null),
    ('SM-FINANCE', 'Finance Officer', 'SENIOR', null, null),
    ('SM-FITNESS', 'Fitness Officer', 'SENIOR', null, null),
    ('SM-HSO', 'Health Services Officer', 'SENIOR', null, null),
    ('SM-HISTORY', 'Historian', 'SENIOR', null, null),
    ('SM-HOMELAND', 'Homeland Security Officer', 'SENIOR', null, null),
    ('SM-IT', 'Information Technologies Officer', 'SENIOR', null, null),
    ('SM-LEADERSHIP', 'Squadron Leadership Officer', 'SENIOR', null, null),
    ('SM-LOGISTICS', 'Logistics Officer', 'SENIOR', null, null),
    ('SM-MX', 'Maintenance Officer', 'SENIOR', null, null),
    ('SM-OPS', 'Operations Officer', 'SENIOR', null, null),
    ('SM-PERSONNEL', 'Personnel Officer', 'SENIOR', null, null),
    ('SM-PAO', 'Public Affairs Officer', 'SENIOR', null, null),
    ('SM-RECRUITING', 'Recruiting Officer', 'SENIOR', null, null),
    ('SM-SAFETY', 'Safety Officer', 'SENIOR', null, null),
    ('SM-SAR', 'Search and Rescue Officer', 'SENIOR', null, null),
    ('SM-NCO', 'Squadron NCO', 'SENIOR', null, null),
    ('SM-STANDEVAL', 'Standards/Evaluation Officer', 'SENIOR', null, null),
    ('SM-SUPPLY', 'Supply Officer', 'SENIOR', null, null),
    ('SM-TESTING', 'Testing Officer', 'SENIOR', null, null),
    ('SM-TRANSPORTATION', 'Transportation Officer', 'SENIOR', null, null),
    ('SM-WSA', 'Web Security Administrator', 'SENIOR', null, null),
    ('CDT-NCO-ACTIVITY', 'Cadet Activities NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-ACTIVITY', 'Cadet Activities Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-ADMIN', 'Cadet Administrative NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-ADMIN', 'Cadet Administrative Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-AE', 'Cadet Aerospace Education NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-AE', 'Cadet Aerospace Education Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-CC', 'Cadet Commander', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-COMMUNICATIONS', 'Cadet Communications NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-COMMUNICATIONS', 'Cadet Communications Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-CYBER', 'Cadet Cyber Education NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-CYBER', 'Cadet Cyber Education Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-CDO', 'Cadet Deputy Commander for Operations', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-CDS', 'Cadet Deputy Commander for Support', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-EL', 'Cadet Element Leader', 'CADET', 'C/Amn', 'C/CMSgt'),
    ('CDT-NCO-ES', 'Cadet Emergency Services NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-ES', 'Cadet Emergency Services Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-AMN-FINANCE', 'Cadet Finance Airman', 'CADET', 'C/A1C', 'C/SrA'),
    ('CDT-CCF', 'Cadet First Sergeant', 'CADET', 'C/MSgt', 'C/CMSgt'),
    ('CDT-NCO-FITNESS', 'Cadet Fitness Education NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-FITNESS', 'Cadet Fitness Education Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-FLT-CC', 'Cadet Flight Commander', 'CADET', 'C/MSgt', 'C/Capt'),
    ('CDT-FLT-CCF', 'Cadet Flight Sergeant', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-CAC-GPA', 'Cadet GCAC Assistant', 'CADET', null, null),
    ('CDT-CAC-GP', 'Cadet GCAC Representative', 'CADET', null, null),
    ('CDT-NCO-HISTORY', 'Cadet Historian NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-HISTORY', 'Cadet Historian Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-IT', 'Cadet IT NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-IT', 'Cadet IT Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-LEADERSHIP', 'Cadet Leadership Education NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-LEADERSHIP', 'Cadet Leadership Education Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-LOGISTICS', 'Cadet Logistics NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-LOGISTICS', 'Cadet Logistics Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-OPERATIONS', 'Cadet Operations NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-OPS', 'Cadet Operations Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-AMN-PERSONNEL', 'Cadet Personnel Airman', 'CADET', 'C/A1C', 'C/SrA'),
    ('CDT-NCO-PA', 'Cadet Public Affairs NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-PA', 'Cadet Public Affairs Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-RECRUITING', 'Cadet Recruiting NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-RECRUITING', 'Cadet Recruiting Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-NCO-SAFETY', 'Cadet Safety NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-SAFETY', 'Cadet Safety Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-AMN-SUPPLY', 'Cadet Supply Airman', 'CADET', 'C/A1C', 'C/SrA'),
    ('CDT-NCO-SUPPLY', 'Cadet Supply NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('CDT-OFF-SUPPLY', 'Cadet Supply Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('CDT-CAC-WGA', 'Cadet WCAC Assistant', 'CADET', null, null),
    ('CDT-CAC-WG', 'Cadet WCAC Representative', 'CADET', null, null),
    ('CDT-AMN-WEB', 'Cadet Web Maintenance Airman', 'CADET', 'C/A1C', 'C/SrA');
