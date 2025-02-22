# API_REST

# Before you run the migrations, make sure you create the database on MySQL. Run the following command:
CREATE DATABASE namedatabase;

# Apply migration to the user table
mysql -u root -p namedatabase < internal/db/migrations/create_users.up.sql

# Apply migration to the events table 
mysql -u root -p namedatabase < internal/db/migrations/create_events.up.sql

# After running the commands abouve, acess to MySQL and make sure the tables were created correctly
USE namedatabase;
SHOW TABLES;