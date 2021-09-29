
USE crashcourse;

DROP PROCEDURE IF EXISTS ordertotal;

DELIMITER //

CREATE PROCEDURE ordertotal(
    IN onumber INT,
    OUT ototal DECIMAL(8,2)
)
BEGIN
    SELECT Sum(item_price * quantity)
    FROM orderitems
    WHERE order_num = onumber
    INTO ototal;
END //

DELIMITER ;