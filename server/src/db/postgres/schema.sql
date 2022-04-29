-- noinspection SqlNoDataSourceInspectionForFile

CREATE TABLE IF NOT EXISTS USERS(
    id VARCHAR(100),
    username VARCHAR(100),
    email VARCHAR(255),
    password VARCHAR(255),
    create_at BIGINT,
    update_at BIGINT,
    delete_at BIGINT,
    PRIMARY KEY(ID)
);

CREATE INDEX user_email_index(email);

CREATE TABLE IF NOT EXISTS ROOMS(
    id VARCHAR(100),
    name VARCHAR(100),
    description VARCHAR(1000),
    create_at BIGINT,
    update_at BIGINT,
    delete_at BIGINT
);

CREATE INDEX room_id_index(id);


CREATE TABLE USER_ROOM(
    user_id VARCHAR(100),
    room_id VARCHAR(100),
    foreign key(user_id) references USERS(id),
    foreign key(room_id) references ROOMS(id)
)