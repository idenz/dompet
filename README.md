## Requirement
- Golang
- MongoDB
- Docker (no required)

### Clone project
```sh
> git clone https://github.com/idenz/dompet.git
```

### Start Project

Don't change dbhost in folder [config.yml](https://github.com/idenz/dompet/blob/main/config.yml) config.yml if you use docker compose

```sh
# Docker running
> docker compose up

# Golang running
> go run main.go
```