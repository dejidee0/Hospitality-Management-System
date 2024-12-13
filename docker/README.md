To run this docker image

unpack the tar file with the command below
```docker load -i find_peace_api:1.0.0.tar```

run a container of the image with the command below
```docker run --env-file .env -p 8080:8080 find_peace_api:1.0.0```

NOTE: make sure you have .env file setup in the directory where you running the docker dockerrun command
