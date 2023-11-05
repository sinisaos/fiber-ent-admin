Example application using [Ent](https://entgo.io/), [Fiber](https://gofiber.io/) and [React Admin](https://marmelab.com/react-admin/).

## How to use

Clone the repo.
```
$ git clone https://github.com/sinisaos/fiber-ent-admin.git
$ cd fiber-ent-admin
$ cp .env.example .env && rm .env.example
$ make server
```

## Generate the schema
To generate changes in the Ent schema.
```
$ make generate
```

## Run tests
To run tests.
```
$ make tests
```

After application is running you can visit ``localhost:3000/swagger/`` and use interactive Swagger documentation which was generated using [Entoas](https://github.com/ent/contrib/tree/master/entoas), with minor modifications to use authorization.

## Admin interface

To create superuser.
```
$ make superuser
```
When the superuser is created, restart the server and at ``localhost:3000`` we can login as superuser in the admin interface. After that we can perform CRUD operations with basic filtering and sorting.