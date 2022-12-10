-- create url shortener database
CREATE
DATABASE url_shortener;

-- create links table
CREATE TABLE urls
(
    url_code  VARCHAR(10) PRIMARY KEY,
    long_url  TEXT UNIQUE,
    short_url VARCHAR(30)
);