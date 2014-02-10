#! /usr/bin/env bash

cowsay "`printf 'the cow says, "%s"' "$(go run talky.go < /dev/fd/0)"`"
