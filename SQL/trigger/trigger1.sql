USE crashcourse;

-- create trigger
CREATE TRIGGER newproduct AFTER INSERT ON products
FOR EACH ROW SELECT 'product added';

-- delete trigger 
-- DROP TRIGGER newproduct;

