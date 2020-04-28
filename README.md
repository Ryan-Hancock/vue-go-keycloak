# Vue Auth using Go and Keycloak

Implementation of an auth flow using Vue, Go and Keycloak. 

## Set up 

Configuration for keycloak can be found in the `docker-compose.yml`, here you can
change user/password and set the realm json.

```bash
docker run -d -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=password -e KEYCLOAK_IMPORT=./keycloack.json -p 8080:8090 jboss/keycloak
```

or

```bash
docker-compose up -d
```  
to run the keycloak server `port: 8080`.  

```bash
go build && ./VueJWT
```  
to run the go server `port: 8090`.  

```bash
npm install && npm run dev
```   
to install and run the vue dev server `port: 8070`.  

## TODO

1. Add all the services to docker with easy configuration.  
2. Improve data storage for token on frontend.
3. Add nginx container to handle connections.
