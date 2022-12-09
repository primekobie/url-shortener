-- create url shortener database
CREATE
DATABASE url_shortener;

-- create links table
CREATE TABLE links
(
    url_code  VARCHAR(30) PRIMARY KEY,
    long_url  TEXT UNIQUE,
    short_url VARCHAR(60)
);