CREATE TABLE article (
    id int(11) unsigned NOT NULL AUTO_INCREMENT,
    title varchar(60),
    content text,
    published boolean NOT NULL,
    published_at TIMESTAMP NOT NULL,
    PRIMARY KEY(`id`)
);

