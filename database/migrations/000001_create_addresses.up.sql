CREATE TABLE employee.addresses (
	id bigserial NOT NULL,
	street text NULL,
	city text NULL,
	state text NULL,
	CONSTRAINT addresses_pkey PRIMARY KEY (id)
);