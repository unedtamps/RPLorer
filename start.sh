 #!/bin/sh

 set -e 
 
 echo "start migrate"
 migrate -path /app/migration -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:5432/$POSTGRES_DB?sslmode=disable" --verbose up

 echo "start app"
 exec "$@"
