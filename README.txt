The goal of this repo is to make a working demo of traefik reverse proxy for GRPC, in docker mode.

STATUS : for now, the problem is that traefik connects to the backend using the ip address. Since the ip can change any time, in many sceanrios,
       	 there is no way currently for traefik to validate the authenticity of the backend's certificate.
         Therefore, one solution is to generate the certificate matching the ip address at the container's startup, altough you should
         be aware of the security implications when implementing it. 

structure :

.
├── certs
│   ├── backend.cert
│   ├── backend.key
│   ├── frontend.cert
│   └── frontend.key
├── docker-compose.yaml
├── hello-world-grpc
│   ├── Dockerfile
│   ├── client.go
│   └── server.go
└── traefik
    └── traefik.conf


in 'certs', you can find the certificates used for ssl. they were generated with the following commands
  # openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./backend.key -out ./backend.crt
  # openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./frontend.key -out ./frontend.crt
using this demo : https://docs.traefik.io/user-guide/grpc/#a-grpc-example-in-go

in 'hello-world-grpc', there is the source for the image aguilbau/hello-world-grpc:ssl
it is a modified version of google's grpc example
the differences are :
  - server uses a certificate for ssl, so grpc.WithInsecure() won't work.
  - the client accepts 2 env variables : ADDRESS for the server endpoint
                                         AUTHORITY to set a custom Authority header

in 'traefik', you can find the config file used for traefik


The docker-compose file launches a traefik instance in docker mode.

To test: 
  # docker-compose up --build
  # docker exec -it $(docker ps | grep aguilbau/hello-world-grpc:ssl | grep 'tail -f'  | awk '{print $1}') bash
  inside the container, run
  # /client


For now, the result of the test is :

time="2017-10-02T13:38:59Z" level=warning msg="Error forwarding to https://172.21.0.4:50051, err: x509: cannot validate certificate for 172.21.0.4 because it doesn't contain any IP SANs"
