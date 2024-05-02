-- homework 2

select array_agg(s.name), g.grade, c.course_name from student s 
join student_course sc on s.student_id = sc.student_id
join grade g on g.student_course_id = sc.student_course_id 
join course c on c.course_id = sc.course_id
where (sc.course_id,g.grade) in 
(select sc.course_id, max(g.grade) from grade g 
join   student_course sc on g.student_course_id = sc.student_course_id
group by  sc.course_id
 ) 
group by g.grade,c.course_name;

-- homework 3

select c.course_name, round(avg(g.grade)::numeric,2) from grade g 
join student_course sc on sc.student_course_id = g.student_course_id
join course c on sc.course_id = c.course_id
group by c.course_name;

-- homework 4

select array_agg(s.name),s.age ,c.course_name from student s 
join student_course sc on sc.student_id = s.student_id 
join course c on c.course_id = sc.course_id 
where s.age in 
(select min(age) from student s 
join student_course sc on s.student_id = sc.student_id 
group by sc.course_id)
group by s.age,c.course_name;

-- homework 5 

select c.course_name, avg(g.grade) as max_average_grade from course c 
join student_course sc on sc.course_id = c.course_id
join grade g on g.student_course_id = sc.student_course_id 
group by c.course_name  
having avg(g.grade) = (
    select max(avg_grade) 
    from (select avg(g.grade) as avg_grade 
    from grade g join student_course sc on sc.student_course_id = g.student_course_id 
    group by sc.course_id) AS subquery
    )
 ;