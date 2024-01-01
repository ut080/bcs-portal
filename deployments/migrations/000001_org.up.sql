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

CREATE TYPE unit_type AS ENUM (
    'Composite Squadron',
    'Cadet Squadron',
    'Senior Squadron',
    'Activity'
);

CREATE TABLE members (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    capid INTEGER NOT NULL UNIQUE CHECK (capid = 0 OR capid >= 100000),
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
    max_grade grade
);

CREATE TABLE duty_assignments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    office_symbol VARCHAR NOT NULL UNIQUE,
    title VARCHAR NOT NULL,
    assistant BOOLEAN NOT NULL DEFAULT false,
    duty_title_id UUID NOT NULL REFERENCES duty_titles (id) ON DELETE CASCADE,
    reports_to UUID REFERENCES duty_assignments (id) ON DELETE SET NULL,
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
    name VARCHAR NOT NULL UNIQUE,
    staff_supergroup_id UUID REFERENCES staff_groups (id) ON DELETE CASCADE,
    leader_id UUID NOT NULL REFERENCES duty_assignments (id) ON DELETE RESTRICT
);

CREATE TABLE staff_group_direct_reports (
    staff_group_id UUID NOT NULL REFERENCES staff_groups (id) ON DELETE CASCADE,
    duty_assignment_id UUID NOT NULL REFERENCES duty_assignments (id) ON DELETE CASCADE,
    PRIMARY KEY (staff_group_id, duty_assignment_id)
);

CREATE TABLE units (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    charter_number VARCHAR NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    commander_id UUID NOT NULL REFERENCES duty_assignments (id) ON DELETE CASCADE
);

INSERT INTO duty_titles (id, code, title, member_type, min_grade, max_grade) VALUES
    ('31303b55-88be-42bf-af8f-3f42b74fa58c', 'SM-ACTIVITY', 'Activities Officer', 'SENIOR', null, null),
    ('004b8a7d-bc9b-4db9-b998-4be3e8059fa6', 'SM-ADMIN', 'Administrative Officer', 'SENIOR', null, null),
    ('db33c16a-687e-4852-a905-97bacee1924f', 'SM-CC-ADVISOR', 'Advisor to the Commander', 'SENIOR', null, null),
    ('291b7ce6-0c17-4b66-99cb-ed7ba73c5288', 'SM-AE', 'Aerospace Education Officer', 'SENIOR', null, null),
    ('6e0393d4-0cb9-4d53-bd7c-65401af9c825', 'SM-ALERT', 'Alerting Officer', 'SENIOR', null, null),
    ('defb6475-a2d2-4961-a063-a360c2a59aec', 'SM-CC', 'Commander', 'SENIOR', null, null),
    ('1e6e89dd-4fa3-451f-8710-981b5427c3a6', 'SM-CHAPLAIN', 'Chaplain', 'SENIOR', null, null),
    ('89028a19-6067-4e2c-ac1b-927afa84c34f', 'SM-CDI', 'Character Development Instructor', 'SENIOR', null, null),
    ('68040829-08b7-4a39-9086-8e18cdf5095f', 'SM-COMMUNICATIONS', 'Communications Officer', 'SENIOR', null, null),
    ('b316d0dd-bb58-4c4a-9e1d-ace950e75649', 'SM-CYBER', 'Cyber Education Officer', 'SENIOR', null, null),
    ('47ef770c-30a6-4fe5-9f9f-2171f6eb68a8', 'SM-DEPUTY', 'Deputy Commander', 'SENIOR', null, null),
    ('67010ad2-4c2b-4d02-886a-04b23fad5309', 'SM-DCC', 'Deputy Commander for Cadets', 'SENIOR', null, null),
    ('57b8a806-ae39-4a20-9046-11713daf55ab', 'SM-DCS', 'Deputy Commander for Seniors', 'SENIOR', null, null),
    ('ebf43530-cec6-418a-93df-2b6554a702c5', 'SM-PREP', 'Disaster Preparedness Officer', 'SENIOR', null, null),
    ('a66bfbb6-27be-4fb5-8206-b45fc551a811', 'SM-DDR', 'Drug Demand Reduction Officer', 'SENIOR', null, null),
    ('a82511eb-5000-4d0c-9ff3-13e2173e3085', 'SM-ETO', 'Education and Training Officer', 'SENIOR', null, null),
    ('d4ba46ac-f218-43d4-9c46-0ab4328874d7', 'SM-ES', 'Emergency Services Officer', 'SENIOR', null, null),
    ('f70a7c49-56cd-41d1-b3e3-58a09d2b344c', 'SM-EST', 'Emergency Services Training Officer', 'SENIOR', null, null),
    ('baf01c15-ec42-4401-9e7a-dc9264d104e6', 'SM-FINANCE', 'Finance Officer', 'SENIOR', null, null),
    ('31d347a7-166c-4af8-8381-6eaacf988423', 'SM-FITNESS', 'Fitness Officer', 'SENIOR', null, null),
    ('13633878-72d4-45bc-8e45-2aed3188cc97', 'SM-HSO', 'Health Services Officer', 'SENIOR', null, null),
    ('5b765281-b529-4ef7-99a6-59ba77a9df65', 'SM-HISTORY', 'Historian', 'SENIOR', null, null),
    ('a8fcb5fe-49d9-4040-8467-49a6cb47f1a1', 'SM-HOMELAND', 'Homeland Security Officer', 'SENIOR', null, null),
    ('1c07babe-07ce-4a82-a9c1-ff11cb5813e0', 'SM-IT', 'Information Technologies Officer', 'SENIOR', null, null),
    ('c3e26655-d8ee-44b2-8f46-8489c3330281', 'SM-LEADERSHIP', 'Squadron Leadership Officer', 'SENIOR', null, null),
    ('eba06ece-a956-439d-829c-793b7d360d10', 'SM-LOGISTICS', 'Logistics Officer', 'SENIOR', null, null),
    ('ed2353a0-4848-4fd1-9d00-24ba1e28b2ba', 'SM-MX', 'Maintenance Officer', 'SENIOR', null, null),
    ('3c3f5016-ed18-488e-bc53-17fe8337f80c', 'SM-OPS', 'Operations Officer', 'SENIOR', null, null),
    ('96f5fc73-79be-491e-ab9c-a558a20c10b4', 'SM-PERSONNEL', 'Personnel Officer', 'SENIOR', null, null),
    ('a01f8d53-5905-4662-8284-b8e339c35434', 'SM-PAO', 'Public Affairs Officer', 'SENIOR', null, null),
    ('53f9d491-76bd-4cc1-beed-109f32f3f450', 'SM-RECRUITING', 'Recruiting Officer', 'SENIOR', null, null),
    ('5aadc559-a9d0-4c55-8604-9dc40296056f', 'SM-SAFETY', 'Safety Officer', 'SENIOR', null, null),
    ('09a59f13-e8a2-4519-9001-7be0a60a0440', 'SM-SAR', 'Search and Rescue Officer', 'SENIOR', null, null),
    ('9b627217-465d-4da7-9a3e-8eb9e8c3939f', 'SM-NCO', 'Squadron NCO', 'SENIOR', null, null),
    ('d85daa4b-72c0-42e7-8692-90bc4adcf657', 'SM-STANDEVAL', 'Standards/Evaluation Officer', 'SENIOR', null, null),
    ('370110ba-12f6-441d-9a51-95f986836cac', 'SM-SUPPLY', 'Supply Officer', 'SENIOR', null, null),
    ('6d4a064b-c2eb-4c8d-86c8-dd200cdf9048', 'SM-TESTING', 'Testing Officer', 'SENIOR', null, null),
    ('815e1efd-879b-4e62-98c0-3b355676f1c0', 'SM-TRANSPORTATION', 'Transportation Officer', 'SENIOR', null, null),
    ('4ffea32f-0a34-4df6-b843-7381d78aa5cb', 'SM-WSA', 'Web Security Administrator', 'SENIOR', null, null),
    ('4f8cd548-f91b-4e14-8d84-1109fae06384', 'CDT-NCO-ACTIVITY', 'Cadet Activities NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('ce984a1a-e481-4a3e-8c21-9bd499bd5508', 'CDT-OFF-ACTIVITY', 'Cadet Activities Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('fc2733e1-6079-4c99-899a-dfb730542f21', 'CDT-NCO-ADMIN', 'Cadet Administrative NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('9ac6305d-0b7c-4143-9ebd-62420e432858', 'CDT-OFF-ADMIN', 'Cadet Administrative Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('c74b7f61-bd67-4eb0-b54d-106bc325fcce', 'CDT-NCO-AE', 'Cadet Aerospace Education NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('e2dffdba-f626-4d94-aec0-dd214abced54', 'CDT-OFF-AE', 'Cadet Aerospace Education Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('1bfe0c38-3c26-43b2-a06e-2ecffa7cf2de', 'CDT-CC', 'Cadet Commander', 'CADET', 'C/2d Lt', 'C/Col'),
    ('863c6496-5329-4ea0-be6e-a52ec80ffa75', 'CDT-NCO-COMMUNICATIONS', 'Cadet Communications NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('b180486f-70c5-4a4a-825e-1f97888b8515', 'CDT-OFF-COMMUNICATIONS', 'Cadet Communications Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('4cb2c9c9-a06b-400c-927d-310546eb8127', 'CDT-NCO-CYBER', 'Cadet Cyber Education NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('bea1e013-f14c-4ff0-8726-bb1d06fbf3e3', 'CDT-OFF-CYBER', 'Cadet Cyber Education Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('e8d59f30-0b45-4962-a132-f9395ca4d706', 'CDT-CDO', 'Cadet Deputy Commander for Operations', 'CADET', 'C/2d Lt', 'C/Col'),
    ('717fa3b7-49d1-4ee0-a664-de342c103025', 'CDT-CDS', 'Cadet Deputy Commander for Support', 'CADET', 'C/2d Lt', 'C/Col'),
    ('ec6ea052-6ac1-437f-a59d-c748862c8393', 'CDT-EL', 'Cadet Element Leader', 'CADET', 'C/Amn', 'C/CMSgt'),
    ('c2c11c2a-0911-49b1-9bc4-9c234f22d20d', 'CDT-NCO-ES', 'Cadet Emergency Services NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('2549d647-79d6-43ae-944b-b434879eba89', 'CDT-OFF-ES', 'Cadet Emergency Services Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('2c6b41fe-6959-4a70-9ba0-f574479cd612', 'CDT-AMN-FINANCE', 'Cadet Finance Airman', 'CADET', 'C/A1C', 'C/SrA'),
    ('fe177c68-78ba-4fa5-aec9-fbefa3c657c7', 'CDT-CCF', 'Cadet First Sergeant', 'CADET', 'C/MSgt', 'C/CMSgt'),
    ('66f99a55-b311-4a30-a619-0c65ba55a082', 'CDT-NCO-FITNESS', 'Cadet Fitness Education NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('f1d8db17-b786-4265-928e-9e525842f05c', 'CDT-OFF-FITNESS', 'Cadet Fitness Education Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('f1095ef3-6851-4497-bb74-b704c467e089', 'CDT-FLT-CC', 'Cadet Flight Commander', 'CADET', 'C/MSgt', 'C/Capt'),
    ('74d88591-e845-4031-8818-991cd1902f61', 'CDT-FLT-CCF', 'Cadet Flight Sergeant', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('dffea29f-f93a-4a58-be5b-806701a6b3aa', 'CDT-CAC-GPA', 'Cadet GCAC Assistant', 'CADET', null, null),
    ('55952d0c-0be9-4a06-a9e9-5abca77485c4', 'CDT-CAC-GP', 'Cadet GCAC Representative', 'CADET', null, null),
    ('79b4cd6e-d3c5-4720-90a7-b25a3d5aa71b', 'CDT-NCO-HISTORY', 'Cadet Historian NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('676cc682-6743-4c18-a291-cd8cd7d6c372', 'CDT-OFF-HISTORY', 'Cadet Historian Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('568550cc-823e-408c-a916-2dcbd0e5a7eb', 'CDT-NCO-IT', 'Cadet IT NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('b96d3574-1aa6-419b-8f1b-16c3014089bc', 'CDT-OFF-IT', 'Cadet IT Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('afc0fce4-1e0d-4077-b911-829f5ce22468', 'CDT-NCO-LEADERSHIP', 'Cadet Leadership Education NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('c30397ac-d1f1-4d20-b8ab-96d24cead115', 'CDT-OFF-LEADERSHIP', 'Cadet Leadership Education Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('a05561c0-9bc7-48dc-80d4-9d1ac7b5129f', 'CDT-NCO-LOGISTICS', 'Cadet Logistics NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('fcf71444-9cbd-448d-b9a7-0807234c7d13', 'CDT-OFF-LOGISTICS', 'Cadet Logistics Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('34ecba57-84c1-4a3e-8ca7-5eaa52489f99', 'CDT-NCO-OPERATIONS', 'Cadet Operations NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('b9a83916-cb33-4bc1-942d-2cb1e14f2692', 'CDT-OFF-OPS', 'Cadet Operations Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('e3b31a0b-8220-4d85-936e-07b15d2f905f', 'CDT-AMN-PERSONNEL', 'Cadet Personnel Airman', 'CADET', 'C/A1C', 'C/SrA'),
    ('a62b55ff-df55-49e6-85ad-8f491ff49ee8', 'CDT-NCO-PA', 'Cadet Public Affairs NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('b9814756-8753-41a3-8741-1aee549ecf7e', 'CDT-OFF-PA', 'Cadet Public Affairs Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('a0861bbd-d850-466a-87b1-39e054444dfb', 'CDT-NCO-RECRUITING', 'Cadet Recruiting NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('a70cc6e7-f69d-46ab-9f28-69cd3de2f2b2', 'CDT-OFF-RECRUITING', 'Cadet Recruiting Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('9256b72d-d013-4fce-a154-849ae1715cda', 'CDT-NCO-SAFETY', 'Cadet Safety NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('c2c2d953-87db-41ec-bd4a-df773498e101', 'CDT-OFF-SAFETY', 'Cadet Safety Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('8cd803ab-73a3-4278-a928-b71865ff1d09', 'CDT-AMN-SUPPLY', 'Cadet Supply Airman', 'CADET', 'C/A1C', 'C/SrA'),
    ('0896211f-ee86-460e-ab4b-ff153482f8cd', 'CDT-NCO-SUPPLY', 'Cadet Supply NCO', 'CADET', 'C/SSgt', 'C/CMSgt'),
    ('71e5cbca-d75d-4c6b-8966-8d32fa3e0d49', 'CDT-OFF-SUPPLY', 'Cadet Supply Officer', 'CADET', 'C/2d Lt', 'C/Col'),
    ('e634c06c-c5ed-4c3f-a046-68bc91a46941', 'CDT-CAC-WGA', 'Cadet WCAC Assistant', 'CADET', null, null),
    ('10459b80-ccb6-48ce-8f00-b4c75cd83d5b', 'CDT-CAC-WG', 'Cadet WCAC Representative', 'CADET', null, null),
    ('c5914b12-bef6-462d-b2ff-c0cdfef56481', 'CDT-AMN-WEB', 'Cadet Web Maintenance Airman', 'CADET', 'C/A1C', 'C/SrA');
