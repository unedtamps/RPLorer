 #!/bin/sh

 set -e 
 
 echo "start migrate"
 migrate -path /app/migration -database "$DB_URI" --verbose up

 echo "start up"
 exec "$@"
