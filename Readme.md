docker pull postgres docker images docker run --name postgresql -e POSTGRES_USER=myusername -e
POSTGRES_PASSWORD=mypassword -p 5432:5432 -v ~/Desktop/temp/data:/var/lib/postgresql/data -d postgres

docker pull dpage/pgadmin4:latest docker run --name my-pgadmin -p 82:80 -e 'PGADMIN_DEFAULT_EMAIL=user@domain.local'
-e 'PGADMIN_DEFAULT_PASSWORD=postgresmaster' -d dpage/pgadmin4
