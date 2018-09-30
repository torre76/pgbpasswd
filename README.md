# PgbPasswd - A helper to build PgBouncer auth_file using CLI

[![Build Status](https://travis-ci.org/torre76/pgbpasswd.svg?branch=master)](https://travis-ci.org/torre76/pgbpasswd)

This *tool* build a [PostgreSQL](https://www.postgresql.org) [MD5](https://en.wikipedia.org/wiki/MD5) password hash starting from a valid login and password.

It is an helper tool useful when you are trying to configure [PgBouncer](https://pgbouncer.github.io) and its relative `auth_file` config property.  
PgBouncer companion script [mkauth.py](https://github.com/pgbouncer/pgbouncer/blob/044082be9d6bc6577fe5700db917b3f48d20b87a/etc/mkauth.py), that creates a valid `auth_file` starting from a PostgreSQL instance, assumes that you have access to `pg_shadow`, which is not always true (an example is [Amazon RDS for PostgreSQL](https://aws.amazon.com/it/rds/postgresql/), the master user cannot access such informations).

The formula to obtain a valid hash is quite simple:

```
    hash = 'md5' + md5(login + password)
```

but, when you are working with [CLI](https://en.wikipedia.org/wiki/Command-line_interface) you hope for a *quick 'n' dirty* script that will do the job for you.

`pgbpasswd` produces the hash requested by *PgBouncer* starting from a valid login and password and optionally it appends the data to an existing `auth_file` (or create it if is missing).

## Usage

```bash
pgbpasswd <login> <password>
```

will output the hash and the row that could be appended to `auth_file` to stanrdad output.

```bash
pgbpasswd <login> <password> -f <filename>
```

will output the hash to standard output and append a line to `auth_file` with login and hashed password.

## Build

`pgbpasswd` is written in [Go](https://golang.org) and it has a `Makefile` that handles the build procedure.

Assuming you have a *valid installation of Go in place* and *autotools installed*, the command to build is:

```bash
make
```

the compiled command will be into the `build` directory of the project (the script is statically linked and ready to be deployed in place).