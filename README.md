# Gin Template
> Template for Gin & React projects.

## Features
+ [x] Built-in user management.
+ [x] [GitHub OAuth login](https://github.com/settings/applications/new).
+ [x] Email verification & password reset.
+ [x] Request rate limitation.
+ [x] Use GitHub Actions to build releases & Docker images. 
+ [x] Mobile friendly UI.

## Usage
1. Download built binaries from [GitHub Releases](https://github.com/songquanpeng/gin-template/releases/latest).
2. Run it: 
   1. `chmod u+x gin-template`
   2. `./gin-template --port 3000`
3. The username for the initial account is `root` and the password is `123456`.

## Basic Configurations
The system works out of the box.

There are several environment variables used to configure the system:
1. `REDIS_CONN_STRING`: if set, will use Redis as the store of rate limitation instead of memory.
   + Example: `REDIS_CONN_STRING=redis://default:redispw@localhost:49153`
2. `SESSION_SECRET`: if set, will fix session secret.
   + Example: `SESSION_SECRET=random_string`
3. `SQL_DSN`: if set, will use target SQL database instead of SQLite.
   + Example`SQL_DSN=root:123456@tcp(localhost:3306)/gofile`