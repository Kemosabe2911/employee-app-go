CREATE TABLE employee.projects (
	id bigserial NOT NULL,
	"name" text NULL,
	description text NULL,
	is_active bool NULL,
	CONSTRAINT projects_pkey PRIMARY KEY (id)
);