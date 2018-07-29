#!/bin/bash
go build && ./${PWD##*/} # executable name is the same as the directory name
