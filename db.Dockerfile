# Use the official MySQL image from Docker Hub
FROM mysql:latest

# Set environment variables for MySQL
ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=todoapp
ENV MYSQL_USER=your_user
ENV MYSQL_PASSWORD=password

# Customize MySQL configuration
RUN echo '[mysqld]' >> /etc/mysql/my.cnf
RUN echo 'bind-address = 0.0.0.0' >> /etc/mysql/my.cnf

# Expose the default MySQL port
EXPOSE 3308

# The CMD instruction sets the command to be executed when running the container
CMD ["mysqld"]

