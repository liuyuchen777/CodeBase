USE crashcourse;

DELIMITER //

DROP PROCEDURE IF EXISTS processorders;

CREATE PROCEDURE  processorders()
BEGIN
    -- declare local varibles
    DECLARE done BOOLEAN DEFAULT 0;
    DECLARE o INT;
    DECLARE t DECIMAL(8, 2);

    -- declare the cursor
    DECLARE ordernumbers CURSOR 
    FOR
    SELECT order_num FROM orders;

    -- declare continue handler
    DECLARE CONTINUE HANDLER FOR SQLSTATE '02000' SET done=1;

    -- delete table if exist 
    DROP TABLE IF EXISTS ordertotals;

    -- create new table to store the results
    CREATE TABLE IF NOT EXISTS ordertotals
    (order_num INT, total DECIMAL(8, 2));

    -- open the cursor
    OPEN ordernumbers;

    -- loop through all rows
    REPEAT

        -- get order number
        FETCH ordernumbers INTO o;

        -- GET the total for this order
        CALL ordertotal(o, 1, t);

        -- insert order and total into ordertotals
        INSERT INTO ordertotals(order_num, total)
        VALUES(o, t);
    
    -- end loop
    UNTIL done END REPEAT;

    -- close the cursor
    CLOSE ordernumbers;
END //

DELIMITER ;

call processorders();

select * from ordertotals order by total asc;