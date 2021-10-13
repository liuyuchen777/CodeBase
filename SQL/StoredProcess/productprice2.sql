USE crashcourse;

DROP PROCEDURE IF EXISTS productpricing;

DELIMITER //

CREATE PROCEDURE productpricing(
    OUT pl DECIMAL(8,2),
    OUT ph DECIMAL(8,2),
    OUT pa DECIMAL(8,2)
)
BEGIN
    SELECT Min(prod_price)
    INTO pl
    FROM products;
    SELECT Max(prod_price)
    INTO ph
    FROM products;
    SELECT Avg(prod_price)
    INTO pa
    FROM products;
END //

DELIMITER ;