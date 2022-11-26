# quillnote-server 
### currently in development things might break


Runs a lightweight Nextcloud-like Notes Server. Made for the Quillnote App - https://qosp.org/  

## Purpose

Use the Sync feature of Quillnote app without a full Nextcloud Instance. 
This server emulates all the required APIs for "Notes" from NextCloud.

Currently the notes data is stored in an sqlite db.


## Setup

#### Clone Repo

```
git clone https://github.com/arunk140/quillnote-server.git
cd quillnote-server
```

### With Docker 

#### Build Docker Image

```
docker build -f "Dockerfile" -t quillnoteserver:latest "."
```

#### Run Docker Container

```
docker run -d  -p 3000:3000/tcp quillnoteserver:latest
```

#### Add User using a Bash Shell in the Container (Optional)

New Users can be created automatically just by logging through the Quillnote/Quillpad App - just choose a username and password when setting up sync settings in the App and a new Account will be created. - This feature should be togglable in the future with an enviroment variable.

```
docker exec -it [container-name/id] sh
./server user add [username] [password]
```



### Without Docker

#### Build

```
go build -o server .
```

#### Init DB (required)
```
./server migrate
```


#### Create User (Optional)

New Users can be created automatically just by logging through the Quillnote/Quillpad App - just choose a username and password when setting up sync settings in the App and a new Account will be created. - This feature should be togglable in the future with an enviroment variable.

```
./server user add [username] [password]

```

#### Run Server

```
./server
```


### Quillnote App - https://qosp.org/

In the `Settings` -> `Go to sync settings` -> Set the `Syncing service` to 'Nextcloud' 

In the `Nextcloud Instance URL` add the IP:PORT or the URL for this quillnote-server and for the `Nextcloud account` use the username and password you used when adding an account to the server in the server setup.

---

## TODO 

* Testing 
* Push built Docker Image to a Image Repository
* Add docker-compose
* Add customizable Enviorment Variables 
    * Suport for Postgres/MySQL ... DBs
    * HTTP Server Port and Listen Addr
    * Enable/Disable Advanced Logging
    * Enable/Disable /metrics Endpoint