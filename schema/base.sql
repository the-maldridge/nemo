CREATE TABLE streams (
       id          INTEGER        PRIMARY KEY AUTOINCREMENT,
       title       VARCHAR (512)  NOT NULL,
       description VARCHAR (4096),
       archived    BOOLEAN        DEFAULT FALSE
);

CREATE TABLE events (
       id          INTEGER        PRIMARY KEY AUTOINCREMENT,
       title       VARCHAR (512)  NOT NULL,
       description VARCHAR (4096) NOT NULL,
       archived    BOOLEAN        NOT NULL
                                  DEFAULT FALSE,
       stream_id   INTEGER        REFERENCES streams (id)
                                  NOT NULL
);
