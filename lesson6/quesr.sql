create table if not exists students (
    id uuid primary key default gen_random_uuid() not null,
    name varchar(200) not null,
    age int not null,
    created_at timestamp default current_timestamp, 
    deleted_at timestamp
);
create table if not exists course (
    id uuid primary key default gen_random_uuid() not null,
    name varchar(200) not null,
    created_at timestamp default current_timestamp, 
    deleted_at timestamp
);
create table if not exists grade (
    id uuid primary key default gen_random_uuid() not null,
    student_id uuid references students(id),
    course_id uuid references course(id),
    created_at timestamp default current_timestamp, 
    deleted_at timestamp,
    grade int 
);
create table if not exists student_course (
    id uuid primary key default gen_random_uuid() not null,
    student_id uuid references students(id),
    course_id uuid references course(id),
    created_at timestamp default current_timestamp, 
    deleted_at timestamp
);
insert into students(name,age) values('shamsiddin',20), 
('ali',21), ('vali',22), ('soli',23);
--course ids : b45ec620-db7f-4302-a80d-424038f340c7,   dfcdcba9-d6b6-4128-85c3-a7bef028382e
-- fd298646-2e2b-4d47-b5b0-fab5595fc8f1,   2b6e2cad-ad9a-4250-bf2f-b285f2a4a738


-- student ids : 1eb24db8-d9e5-4e01-9dc3-1e36812b461e,  21fc0267-3a71-4024-ac9c-9884b363f65f
-- a6e93632-9601-446b-9069-6db30a9d04ce,  c02aeb98-66f3-4ce9-824a-1902e4cdf1f3
insert into course(name) values('math'), ('physics'), ('chemistry'),('biology');
insert into student_course(student_id, course_id) 
values ('1eb24db8-d9e5-4e01-9dc3-1e36812b461e','b45ec620-db7f-4302-a80d-424038f340c7'),
('1eb24db8-d9e5-4e01-9dc3-1e36812b461e','dfcdcba9-d6b6-4128-85c3-a7bef028382e'),
('1eb24db8-d9e5-4e01-9dc3-1e36812b461e','fd298646-2e2b-4d47-b5b0-fab5595fc8f1'),

('21fc0267-3a71-4024-ac9c-9884b363f65f','b45ec620-db7f-4302-a80d-424038f340c7'),
('21fc0267-3a71-4024-ac9c-9884b363f65f','dfcdcba9-d6b6-4128-85c3-a7bef028382e'),
('21fc0267-3a71-4024-ac9c-9884b363f65f','fd298646-2e2b-4d47-b5b0-fab5595fc8f1'),

('a6e93632-9601-446b-9069-6db30a9d04ce','b45ec620-db7f-4302-a80d-424038f340c7'),
('a6e93632-9601-446b-9069-6db30a9d04ce','dfcdcba9-d6b6-4128-85c3-a7bef028382e'),
('a6e93632-9601-446b-9069-6db30a9d04ce','fd298646-2e2b-4d47-b5b0-fab5595fc8f1'),
('a6e93632-9601-446b-9069-6db30a9d04ce','2b6e2cad-ad9a-4250-bf2f-b285f2a4a738');






insert into grade(student_id,course_id,grade) 
values('1eb24db8-d9e5-4e01-9dc3-1e36812b461e','b45ec620-db7f-4302-a80d-424038f340c7',5),
('1eb24db8-d9e5-4e01-9dc3-1e36812b461e','dfcdcba9-d6b6-4128-85c3-a7bef028382e',4),
('1eb24db8-d9e5-4e01-9dc3-1e36812b461e','fd298646-2e2b-4d47-b5b0-fab5595fc8f1',5),

('21fc0267-3a71-4024-ac9c-9884b363f65f','b45ec620-db7f-4302-a80d-424038f340c7',2),
('21fc0267-3a71-4024-ac9c-9884b363f65f','dfcdcba9-d6b6-4128-85c3-a7bef028382e',3),
('21fc0267-3a71-4024-ac9c-9884b363f65f','fd298646-2e2b-4d47-b5b0-fab5595fc8f1',4),

('a6e93632-9601-446b-9069-6db30a9d04ce','b45ec620-db7f-4302-a80d-424038f340c7',4),
('a6e93632-9601-446b-9069-6db30a9d04ce','dfcdcba9-d6b6-4128-85c3-a7bef028382e',5),
('a6e93632-9601-446b-9069-6db30a9d04ce','fd298646-2e2b-4d47-b5b0-fab5595fc8f1',3),
('a6e93632-9601-446b-9069-6db30a9d04ce','2b6e2cad-ad9a-4250-bf2f-b285f2a4a738',2);



select s.name,g.grade,c.name 
from students s 
join grade g on g.student_id = s.id 
join course c on c.id = g.course_id
where g.grade in (select max(g.grade) from g where g.course_id = c.id) group by c.name;