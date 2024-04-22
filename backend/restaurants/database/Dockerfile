FROM postgres

ENV POSTGRES_PASSWORD=admin1234

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]