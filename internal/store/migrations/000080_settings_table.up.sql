BEGIN;

CREATE TABLE settings (
    setting text primary key not null,
    value text not null default ''
);

CREATE UNIQUE INDEX settings_lc_idx ON settings ((lower(settings)));

COMMIT;
