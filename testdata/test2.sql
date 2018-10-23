use i8_test
GO

CREATE PROCEDURE GetPersonByAddress @City nvarchar(30) = NULL, @AddressLine1 nvarchar(60) = NULL
AS   
BEGIN
    SET NOCOUNT ON; 

    SELECT  *
    FROM    Person
    WHERE   City = ISNULL(@City,City)
    AND     AddressLine1 LIKE '%' + ISNULL(@AddressLine1 ,AddressLine1) + '%'

END;
GO

CREATE PROCEDURE GetNamesAndPlaces @City nvarchar(30) = NULL
AS   
BEGIN
    SET NOCOUNT ON; 

    SELECT  DISTINCT FullName, City
    FROM    Person
    WHERE   City = ISNULL(@City,City);

    SELECT  DISTINCT AddressLine1, City
    FROM    Person
    WHERE   City = ISNULL(@City,City);

END;
GO 

CREATE TYPE PersonAddressTable AS TABLE   
( 
    full_name           VARCHAR(500), 
    city                VARCHAR(500),
    address_line_1      VARCHAR(500)
);  
GO 

CREATE PROCEDURE SavePeople
    @people PersonAddressTable READONLY
AS   
BEGIN
    SET NOCOUNT ON; 

    INSERT INTO Person(FullName, City, AddressLine1)
    SELECT  full_name, city, address_line_1
    FROM    @people;

    SELECT  AddressLine1, City, FullName
    FROM    Person;
END;
GO 

--EXEC GetPersonByAddress @City = N'Sydney' --, @AddressLine1 = N'hurst';  
--GO