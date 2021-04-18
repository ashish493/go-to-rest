# Go-To-REST API 
A REST API template with a load balancer.

## Installation & Run

### Without Docker

testing
```bash
# Download this project
go get github.com/ashish493/go-to-rest
cd go-to-rest

sudo /etc/init.d/nginx start 

# build and run the RESTful API server
make build
make run

#run the load balancer from it's conf file
cp /nginx/nginx.conf /etc/nginx/nginx.conf
sudo systemctl start nginx
```
Before running the server make sure to change the values in the [config.json](https://github.com/ashish493/go-to-rest/blob/main/config.go) file according to your uses or else leave it to it's default values. You can also load these values from a ENV file, jus change the file name in the [config.go](https://github.com/ashish493/go-to-rest/blob/main/config.go)

### With Docker
You can build the image manually from the [Dockerfile](https://github.com/ashish493/go-to-rest/blob/main/Dockerfile) or you can directly pull the image by following steps:-
```shell
# pull the latest build of the image
docker pull ashishmalik47/go-to-rest:latest
# Running the container
docker run -d -p 5000:5000 go-to-rest
```
## API

#### /projects
* `GET` : Get all projects
* `POST` : Create a new project

#### /projects/:title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /projects/:title/archive
* `PUT` : Archive a project
* `DELETE` : Restore a project 

#### /projects/:title/tasks
* `GET` : Get all tasks of a project
* `POST` : Create a new task in a project

#### /projects/:title/tasks/:id
* `GET` : Get a task of a project
* `PUT` : Update a task of a project
* `DELETE` : Delete a task of a project

#### /projects/:title/tasks/:id/complete
* `PUT` : Complete a task of a project
* `DELETE` : Undo a task of a project

## Usage
The API endpoint is http://localhost:5000.
Try the URL `http://localhost:5000/projects` in a browser to get a response of all listed projects. 

If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following :-

```shell
curl -v http://localhost:5000/projects

curl -d 'title=todo' http://localhost:5000/projects/{title}

curl -X DELETE http://localhost:5000/projects/{todo}
```
