CREATE TABLE employee.employees (
	id bigserial NOT NULL,
	"name" text NULL,
	username text NULL,
	"password" text NULL,
	age int8 NULL,
	is_active bool NULL,
	department_id int8 NULL,
	role_id int8 NULL,
	address_id int8 NULL,
	CONSTRAINT employees_pkey PRIMARY KEY (id)
);


-- employee.employees foreign keys

ALTER TABLE employee.employees ADD CONSTRAINT fk_employee_employees_address FOREIGN KEY (address_id) REFERENCES employee.addresses(id);
ALTER TABLE employee.employees ADD CONSTRAINT fk_employee_employees_department FOREIGN KEY (department_id) REFERENCES employee.departments(id);
ALTER TABLE employee.employees ADD CONSTRAINT fk_employee_employees_role FOREIGN KEY (role_id) REFERENCES employee.roles(id);