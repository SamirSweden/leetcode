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



--
--WHILE boolean_expression
--    { sql_statement | statement_block | BREAK | CONTINUE }
-- ЦИКЛЫ В БД 

DECLARE @Counter INT  = 0
WHILE @Counter <= 10 
BEGIN 
    SET @Counter += 1
    IF @Counter =  2 BEGIN CONTINUE END;
    PRINT 'COUNTER = ' + CAST(@Counter AS VARCHAR(10))
END;


--



-- BEGIN TRY
BEGIN TRY
    CREATE DATABASE NewDatabase 
END TRY 
BEGIN CATCH
    PRINT 'OPS!'
    PRINT ERROR_MESSAGE()
END CATCH
-- EXAMPLE BELOW
BEGIN TRY 
    CREATE DATABASE N
END TRY
BEGIN CATCH 
    -- DO CODE
END CATCH


--ERROR HANDLING ERROR_LINE, _PROCEDURE , _STATE , _SEVERITY , _NUMBER(50.000 numbers)
-- BEGIN CATCH 

--TROW METHOD 
    BEGIN TRY 
        THROW 50001 , 'NOT ALLOWED'
    END TRY
    BEGIN CATCH 
        
    END CATCH 

--





BEGIN TRY 
    BEGIN TRANSACTION 
    INSERT INTO Person.Person( FirstName )
    VALUES ('PETE');

    PRINT 'DATA INSERTED SUCCESSFULLY .'
END TRY 
BEGIN CATCH 
    PRINT '!!ERROR HERE ' + CONVERT(VARCHAR , ERROR_NUMBER());
END CATCH





SELECT GETDATE() AS CurrentDate;

WAITFOR DELAY '00:00:2';
SELECT GETDATE() AS SetTimeout
GO
-- programmability / stored procedures / 



CREATE PROCEDURE sp_GetNameCount
@FirstName VARCHAR(100)
AS 
BEGIN 
    DECLARE @EmployeeCount NUMERIC

    SELECT @EmployeeCount = COUNT(*)
    FROM Person.Person
    WHERE FirstName = @FirstName
    PRINT 'Number of workers with first name ' + @FirstName + ' Is ' + CAST(@EmployeeCount AS VARCHAR(10))
END;

EXECUTE sp_GetNameCount 'Joe'

-- CREATE PROCEDURE EmployeeNames
  --  AS 
       -- BEGIN TRY 




