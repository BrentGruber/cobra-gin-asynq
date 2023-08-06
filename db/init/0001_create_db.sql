/* This file is only used when starting docker-compose stack on local
        it will create all databases even if they are not needed 

    All sql files in this directory will be run on postgres startup
    by mounting the directory in /docker-entrypoint-initdb.d/
*/
CREATE DATABASE release;
CREATE DATABASE deploy;

/* This DB is used for monolithic mode */
CREATE DATABASE releasr;