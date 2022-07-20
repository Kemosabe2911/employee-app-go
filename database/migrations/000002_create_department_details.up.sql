CREATE TABLE employee.department_details (
	id bigserial NOT NULL,
	department_room text NULL,
	department_code text NULL,
	website text NULL,
	CONSTRAINT department_details_pkey PRIMARY KEY (id)
);