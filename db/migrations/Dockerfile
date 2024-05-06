FROM postgres:15

RUN apt-get update && apt-get install -y postgresql-contrib postgres
COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432

CMD ["./init.sql"]