# work-example

run single docker file: docker run -dp 8080:8080 work-example  
# docker-compose 
make sure to have a volume b4 (not sure if it's made in the root)  
```docker volume create -d local --name pg_data```  
```docker-compose up``` can add detached mode  

article example:  https://www.mitrais.com/news-updates/how-to-dockerize-a-restful-api-with-golang-and-postgres/