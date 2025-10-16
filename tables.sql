use AdventureWorks2019
select @@VERSION

SELECT year( getdate())

SELECT * 
from Person.Person



select * 
from HumanResources.Employee as e
where 'Kim' in (
    select FirstName 
    from Person.Person
    where BusinessEntityID = e.BusinessEntityID
)





select * 
from HumanResources.Employee

select 
    ( select FirstName 
       from  Person.Person
       where BusinessEntityID = e.BusinessEntityID)
from HumanResources.Employee as e






SELECT *
from  (
 SELECT
    FirstName 
   ,LastName 
   from Person.Person
   ) as p
   where FirstName = 'Kim'

with PersonKim as(
    SELECT
    FirstName 
   ,LastName 
   from Person.Person
   
   where FirstName = 'Kim'
)

select LastName 
from PersonKim

 SELECT *
from  (
 SELECT
    FirstName 
   ,LastName 
   from Person.Person
   ) as p
   where FirstName = 'Kim'


WITH SampleCTE as (
    select 'Test'  as txt
)
-- SampleCTE -- virtual view

select *
from SampleCTE;




WITH 
SampleCTE as (
    select 'Test'  as txt
),

SeCTE as (
    select *  as SampleCTE
),

select *
from SeCTE;





select 
    Name , 
    case 
        when StandardCost < 50 then 'cheap' 
        when StandardCost < 200 then 'so-so' 
        else  'expensive' 
    end as priceCategory 
from Production.Product
where StandardCost <> 0
and  case 
        when StandardCost < 50 then 'cheap' 
        when StandardCost < 200 then 'so-so' 
        else  'expensive' 
    end = 'expensive' 






select *


