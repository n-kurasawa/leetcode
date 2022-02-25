select email as Email count(email) as cnt from Person group by email having cnd > 1;
