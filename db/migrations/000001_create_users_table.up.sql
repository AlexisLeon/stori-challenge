CREATE TABLE IF not EXISTS {{ index .Options "Namespace" }}.users (
    id UUID DEFAULT gen_random_uuid(),
    email varchar(255) NULL,
    PRIMARY KEY (id)
);

comment on table {{ index .Options "Namespace" }}.users is 'Users: the users of the application';
