FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD: "root"
ENV MYSQL_DATABASE: tester
ENV MYSQL_USER: 'user'
ENV MYSQL_PASSWORD: 'pass'

COPY tester.sql tester.sql

RUN "mysql < tester.sql"