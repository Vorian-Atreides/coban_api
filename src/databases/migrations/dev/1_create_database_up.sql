
CREATE DATABASE 'coban_dev';

CREATE USER 'coban'@'localhost' IDENTIFIED BY 'password';

GRANT ALL ON coban_dev.* TO 'coban'@'localhost';