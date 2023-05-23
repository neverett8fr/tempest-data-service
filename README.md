# tempest-data-service
this service is responsible for handling data storage and retrieval operations.  
this service is designed to be used with `gcloud`s `bucket` storage  
  
# Features  
- data upload 
- data retrieval   

  
# How to run  
this application contains a `Dockerfile` - this allows you run build and run the service using Docker console commands   
## Build  
```bash
docker build -t .
 ```
   
 ## Run  
 ```bash
docker run -p 8080:8080 -v . -e ENV_VARIABLE=value .
 ```
   
 ## Stop the container  
 ```bash
 docker stop container-name
 ```
