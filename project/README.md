
Build the image using the Dockerfile 
# docker build -t PostsProject .
Then run the container to listen on 8080 port number
# docker run -d --name go_project -p 8080:8080 PostsProject
To test http://localhost:8080


