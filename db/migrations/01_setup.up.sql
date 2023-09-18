CREATE TABLE article (
       id      INTEGER,
       remote  VARCHAR,
       content BLOB,

       CONSTRAINT pk_content PRIMARY KEY (id)
);

CREATE TABLE exercise (
       id         INTEGER,
       article_id INTEGER,
       remote     VARCHAR,
       content    BLOB,
       done       BOOLEAN,

       CONSTRAINT pk_exercise PRIMARY KEY (id),
       CONSTRAINT fk_article  FOREIGN KEY (article_id)
                              REFERENCES  article (id)
);
