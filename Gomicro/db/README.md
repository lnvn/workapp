## Running following command
### Change directory
```
$ cd Gomicro/db/
```

### Grant execute permission and create docker image
```
$ chmod +x start.sh
$ ./start.sh
```

### Access into container and database
```
$ docker exec -it <container id> bash
$ psql -U postgres
```

### Create schema
Looking for shcema.sql, copy and paste schema onto psql cmdline and run

### Insert data
Before create chema, looking for data.sql, copy and paste query onto psql cmdline and run
