CREATE DATABASE mytest;
GO

USE mytest;
GO

CREATE TYPE tvptable AS TABLE   
( 
    p_binary            BINARY(3),
    p_varchar           VARCHAR(500), 
    p_nvarchar          NVARCHAR(100),
    p_id                UNIQUEIDENTIFIER, 
    p_varbinary         VARBINARY(MAX),
    p_tinyint           TINYINT,
    p_smallint          SMALLINT,
    p_int               INT,
    p_bigint            BIGINT,
    p_bit               BIT,
    p_float             FLOAT
);  
GO 

CREATE PROCEDURE spwithtvp
    @param1 tvptable READONLY,
    @param2 tvptable READONLY,
    @param3 NVARCHAR(10)
AS   
BEGIN
    SET NOCOUNT ON; 

    SELECT * FROM @param1;
    SELECT * FROM @param2;
    SELECT @param3;
END;
GO