- git clone stuff...
- make prepare
- run in 4 sperate console tabs:
```bash
cd auth-svc && make server
cd crawler-svc && make server
cd post-crud-svc && make server
cd api-gateway && make server
```