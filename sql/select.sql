-- The task
-- List down the customers with the highest order value for each month of the a particular date range 
-- given tables - orders and orderdetails


SELECT EXTRACT(YEAR FROM orderdate) AS year, 
MONTHNAME(orderdate) AS month, 
customerID, 
SUM(orderCost) as sum_of_order_cost, 
MAX(sum_of_order_cost) FROM order as o
INNER JOIN orderdetails od ON o.orderid=od.orderid 
GROUP BY MONTH(orderdate), customerID 
order by orderCost DESC;