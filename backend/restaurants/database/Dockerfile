FROM postgres

ENV POSTGRES_PASSWORD=admin1234
ENV POSTGRES_USER=fabricio
ENV POSTGRES_DB=restaurantdb

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]