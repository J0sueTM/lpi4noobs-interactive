CREATE TABLE _session (
       id          INTEGER,
       username    VARCHAR,
       article_id  INTEGER,
       exercise_id INTEGER,

       CONSTRAINT pk_session PRIMARY KEY (id),
       CONSTRAINT fk_session_article FOREIGN KEY (article_id)
                                     REFERENCES article(id),
       CONSTRAINT fk_session_exercise FOREIGN KEY (exercise_id)
                                      REFERENCES exercise(id)
);
