-- Create the database
CREATE DATABASE IF NOT EXISTS `smart_mess`;

-- Create the user 'smart_mess_user' with the specified password
CREATE USER IF NOT EXISTS 'smart_mess_user'@'%' IDENTIFIED BY '12345678';

-- Grant all privileges to 'smart_mess_user' on the 'smart_mess' database
GRANT ALL PRIVILEGES ON smart_mess.* TO 'smart_mess_user'@'%';

-- Apply changes
FLUSH PRIVILEGES;
