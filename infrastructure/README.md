# This folder contains the code to create the infrastructure for the 'otp' service




# How to install MariaDB on server:


OPEN mariadb PORTS IN:
/etc/mysql/mariadb.conf.d 50-server.cnf
bind-address = 0.0.0.0 OR the backend server address when this is up

ADD user with permissions:
CREATE USER 'martin'@'%'