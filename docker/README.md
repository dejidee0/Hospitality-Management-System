## STEPS TO RUN THE IMAGE
It depends
### Tarball file
-- step 1: Unpack the tar file with the command below
```docker load -i find_peace_api:1.0.0.tar```

-- step 2: run a container of the image with the command below
```docker run --env-file .env -p 8080:8080 find_peace_api:1.0.0```
NOTE: make sure you have .env file setup in the directory where you running the docker dockerrun command

## pulling direcly from dockerhub
-- step 1: pull image
```docker pull imolebytes/find_peace:1.0.0```
-- step 2: run a container of the downloaded image
```docker run --env-file .env imolebytes/find_peace:1.0.0```

## Or simply pull and run directly with this one command
```docker run --env-file .env imolebytes/find_peace:1.0.0```
