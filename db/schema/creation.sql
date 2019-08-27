CREATE SCHEMA paack;
CREATE TABLE paack.customers(
   cid  SERIAL PRIMARY KEY,
   id CHAR(50),
   firstName CHAR(50),
   lastName CHAR(50),
   email CHAR(50),
   phone CHAR(50)
);