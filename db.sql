


select GETDATE() 
--    current date    --

select distinct day ( ModifiedDate ) -- this func will extract a year 
from Person.Person

select * from Person.Person where ModifiedDate < getdate()

select * from Person.Person
where year(ModifiedDate) = 2007


select distinct * from Person.Person
where ModifiedDate < ('2010-01-01' , '2009-01-01')
--   FirstName = 'Rob' 
--   or FirstName = 'Gail'

-- ) and EmailPromotion >  1; 
where FirstName not in ('Rob','Gail' , 'Peter');




SELECT * FROM users 
WhERE username = 'alice' 



select * 
from Person.Person
where FirstName like '[n - p]%y' -- % ищет именна Начиная с а и б 

select * 
from Person.Person
where FirstName like '[n - p]%' -- % ищет именна Начиная с N и P 


select * 
from Person.Person
where FirstName like '[a,b, c]%' -- % ищет именна Начиная с а и б 

select * 
from Person.Person
where FirstName like 'Ja_' -- % ищет по смиволам 


select * 
from Person.Person
where FirstName like '%b' -- % ищет по смиволам - именно последние символы 

select *
from Person.Person
where FirstName like 'Ja%'-- % representing zero or more chars following ищет по смиволам  





select *
from Person.Person
where Title is not null  -- null is exception


select * 
from HumanResources.Employee
where 
  VocationHours = 10 
  or VocationHours = 20
-- ищет по 10 и 20 


select * 
from HumanResources.Employee
where 
  VocationHours between 10 and 30
  and JobTitle like '%manager%'
-- ищет между 10 и 30 - тут %manager% 


select * 
from HumanResources.Employee
where HireDate > '2010-12-31'
  and JobTitle like '%sales%'
  AND Gender  LIKE 'F'




select top 10  * 
  from HumanResources.Employee
  where HireDate > '2010-12-31'
  and JobTitle like '%sales%'
  order by Gender asc ,  BirthDate asc


select top 10  * 
  from HumanResources.Employee
  where HireDate > '2010-12-31'
  and JobTitle like '%sales%'
  order by 9,  BirthDate asc
--   order by 9, выберает с 0-9 колонок 



select Gender , BirthDate as [Birth Day of Employee] 
  from HumanResources.Employee
  where HireDate > '2010-12-31'
  and JobTitle like '%sales%'
  order by 1,[Birth Day of Employee]








