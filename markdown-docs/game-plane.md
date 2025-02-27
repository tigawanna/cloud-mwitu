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
└── Makefile
```

### Core Application Files
- `/cmd/cloud-mwitu/main.go` - Main application entry point, initializes and starts the server

### Internal Components
- `/internal/api/`
  - `handlers/` - HTTP request handlers for each endpoint (create project, list projects, etc.)
  - `middleware/` - Authentication and other HTTP middleware
  - `routes.go` - Echo router configuration and endpoint definitions

- `/internal/auth/auth.go` - Authentication logic and user management

- `/internal/config/config.go` - Application configuration (ports, paths, environment variables)

- `/internal/db/`
  - `migrations/` - SQLite database schema migrations
  - `sqlite.go` - Database connection and operations

- `/internal/models/models.go` - Data structures for projects, users, and services

- `/internal/service/`
  - `caddy.go` - Caddy reverse proxy configuration management
  - `port.go` - Port availability checking and allocation
  - `systemd.go` - SystemD service management

### Support Files
- `/scripts/`
  - `install.sh` - Installation script for setting up the application
  - `migrations/` - Database migration scripts

- `/configs/`
  - `caddy.json` - Caddy server template configuration
  - `systemd.service` - SystemD service template file

- `/docs/api.swagger.json` - API documentation

- `Makefile` - Build and development commands
