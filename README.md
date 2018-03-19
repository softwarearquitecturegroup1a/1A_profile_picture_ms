# Profile Picture Service - Go with Docker and MongoDB

Overview
========

 * User profile Service: Provides the url picture of a user.
 


Requirements
===========

* Docker 1.12
* Docker Compose 1.8

We must **add virtual domains** in order to use each **api entry point**. By default we are using: **profilepicture.local**

**Virtual domains** has been defined in `docker-compose.yml` file and configured in `/etc/hosts` file. Add the following line in your `/etc/hosts` file:

```
127.0.0.1   profilepicture.local
```

Starting services
==============================

```
docker-compose up -d
```

Stoping services
==============================

```
docker-compose stop
```

Including new changes
==============================

If you need change some source code you can deploy it typing:

```
docker-compose build
```

Restore database information
======================

You can start using an empty database for all microservices, but if you want you can restore a preconfigured data following this steps:

**_Access to mongodb container typing:_**

```
docker exec -it profilepicture-db /bin/bash
```

**_Restore data typing:_**

```
/backup/restore.sh
```

**_Leave the container:_**

```
exit
```


Documentation
======================

## User Service

This service returns information about the profilepicture of Cinema.

**_Routes:_**

* GET - http://profilepicture.local/profilepicture : Get all profilepicture
* POST - http://profilepicture.local/profilepicture : Create user
* DELETE - http://profilepicture.local/profilepicture/{id} : Remove user by id