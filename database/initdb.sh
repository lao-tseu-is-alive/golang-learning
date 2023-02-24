#!/bin/bash
su -c 'createdb golangdb' postgres
su -c 'psql -f initdb.sql golangdb' postgres
