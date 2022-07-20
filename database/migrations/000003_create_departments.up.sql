CREATE TABLE employee.departments (
	id bigserial NOT NULL,
	"name" text NULL,
	department_details_id int8 NULL,
	CONSTRAINT departments_pkey PRIMARY KEY (id)
);


-- employee.departments foreign keys

ALTER TABLE employee.departments ADD CONSTRAINT fk_employee_departments_department FOREIGN KEY (department_details_id) REFERENCES employee.department_details(id);