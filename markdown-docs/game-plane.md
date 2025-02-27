feture i need :
given a  project name ona remote linux box , i want to 
- create a directory of that name
- create a systemd servive of that name , it should also check for a not used port
-  add that searvice to caddy witha suddomain pointo to that name
-  with all the above staeps add a special prefix to allow easy listing adn managing
-  have enpoints to list create and update using echo
-  add some sort of authentication backed by an sqlite db (iitila account creation should be done on the cli during the binary instalation process)
-  

- Posiible File structure 
```sh 
/cloud-mwitu
├── cmd/
│   └── cloud-mwitu/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── routes.go
│   ├── auth/
│   │   └── auth.go
│   ├── config/
│   │   └── config.go
│   ├── db/
│   │   ├── migrations/
│   │   └── sqlite.go
│   ├── models/
│   │   └── models.go
│   └── service/
│       ├── caddy.go
│       ├── port.go
│       └── systemd.go
├── scripts/
│   ├── install.sh
│   └── migrations/
├── configs/
│   ├── caddy.json
│   └── systemd.service
├── docs/
│   └── api.swagger.json
├── go.mod
├── go.sum
```


