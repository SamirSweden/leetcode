use AdventureWorks2019

DECLARE @Counter INT; 
DECLARE @Name NVARCHAR(50); 
DECLARE @CurrentDate DATETIME; 

DECLARE @EmployeeID INT;
DECLARE @EmployeeName NVARCHAR(50);

SET @EmployeeID = 1;


SELECT @EmployeeName = FirstName 
FROM Person.Person
WHERE BusinessEntityID = @EmployeeID;


IF @EmployeeName = 'Ken'
    BEGIN 
        SET @EmployeeID = 1;
        PRINT 'The boss';
    END
ELSE
    BEGIN 
        PRINT 'JUST ANOTHER WORKER' 
    END

print 'Employee Name is ' + @EmployeeName;

SET @Counter = 1; 
SET @Name =  'Pete'; 
SET @CurrentDate  = GETDATE();

PRINT 'Name is ' + @Name;



-- checking for db 
IF NOT EXISTS (
    SELECT * 
    FROM sys.databases
    WHERE name = 'NewDatabase'
)
    CREATE DATABASE NewDatabase
ELSE 
    CREATE DATABASE NewDatabase2



IF DATENAME (weekday , GETDATE()) IN ('SATURDAY','SUNDAY')
    SELECT 'Holiday';
ELSE 
    SELECT 'WORKDAY';

GO
