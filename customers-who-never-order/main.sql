select c.name as Customers 
from Customers as c 
left outer join Orders as o on c.id = o.customerId
where o.id is null;
