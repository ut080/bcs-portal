#!/bin/bash

migrate -database 'postgres://postgres:postgres@localhost:5432/bcs_portal?sslmode=disable' -source 'file://./' "$@"