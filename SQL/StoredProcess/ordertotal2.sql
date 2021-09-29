-- Name: ordertotal.sql
-- Paramters: onumber = order number
--            taxable = 0 if not taxable, 1 if taxable
--            ototal = order total varible

USE crashcourse;

DROP PROCEDURE IF EXISTS ordertotal;

DELIMITER //

CREATE PROCEDURE ordertotal(
    IN onumber INT,
    IN taxable BOOLEAN,
    OUT ototal DECIMAL(8,3)
) COMMENT 'Obtain order total, optionally adding tax'
BEGIN

    -- Declare varible for total
    DECLARE total DECIMAL(8,3);
    -- Declare tax percentage
    DECLARE taxrate INT DEFAULT 6;

    -- Get the order total
    SELECT Sum(item_price*quantity)
    FROM orderitems
    WHERE order_num = onumber
    INTO total;

    -- Is this taxable?
    IF taxable THEN
        -- Yes, so add taxrate to the total
        SELECT total +(total/100*taxrate) INTO total;
    END IF;

    -- finally, save to out varible
    SELECT total INTO ototal;
END //

DELIMITER ;

CALL ordertotal(20005, 0, @total);
SELECT @total;

CALL ordertotal(20005, 1, @total2);
SELECT @total2;