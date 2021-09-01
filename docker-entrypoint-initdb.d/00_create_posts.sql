DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
	id bigserial NOT NULL,
	content varchar(255) NOT NULL,
	CONSTRAINT posts_pkey PRIMARY KEY (id)
);