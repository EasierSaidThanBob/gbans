BEGIN;
CREATE TABLE members
(
    members_id bigserial primary key,
    parent_id  bigint,
    members    jsonb,
    created_on timestamptz,
    updated_on timestamptz
);

CREATE UNIQUE INDEX members_parent_id_uindex ON members (parent_id);

ALTER TABLE ban
    ADD COLUMN include_friends bool default false;

COMMIT;
