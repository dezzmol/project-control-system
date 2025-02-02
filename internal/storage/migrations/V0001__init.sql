CREATE SCHEMA IF NOT EXISTS itmo AUTHORIZATION postgres;

CREATE TABLE IF NOT EXISTS itmo.users
(
    id       VARCHAR(255) NOT NULL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email    VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS itmo.projects
(
    id           VARCHAR(255) NOT NULL PRIMARY KEY,
    project_name VARCHAR(255) NOT NULL,
    description  TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS itmo.users_projects
(
    id         VARCHAR(255) NOT NULL PRIMARY KEY,
    user_id    VARCHAR(255) NOT NULL REFERENCES itmo.users (id),
    project_id VARCHAR(255) NOT NULL REFERENCES itmo.projects (id),
    role       VARCHAR(40)  NOT NULL
);

CREATE TABLE IF NOT EXISTS itmo.milestones
(
    id         VARCHAR(255) NOT NULL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    start_date DATE         NOT NULL,
    end_date   DATE         NOT NULL,
    project_id VARCHAR(255) REFERENCES itmo.projects (id),
    status     VARCHAR(50)  NOT NULL DEFAULT 'open'
);

CREATE TABLE IF NOT EXISTS itmo.tickets
(
    id           VARCHAR(255) NOT NULL PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    description  TEXT         NOT NULL,
    project_id   VARCHAR(255) NOT NULL REFERENCES itmo.projects (id),
    milestone_id VARCHAR(255) NOT NULL REFERENCES itmo.milestones (id),
    user_id      VARCHAR(255)          DEFAULT NULL REFERENCES itmo.users (id),
    status       VARCHAR(50)  NOT NULL DEFAULT 'new'
);

CREATE TABLE IF NOT EXISTS itmo.bug_report
(
    id           VARCHAR(255) PRIMARY KEY,
    description  TEXT         NOT NULL,
    date_created DATE         NOT NULL,
    project_id   VARCHAR(255) NOT NULL REFERENCES itmo.projects (id),
    status       VARCHAR(50)  NOT NULL DEFAULT 'new'
);
