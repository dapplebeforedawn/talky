#! /usr/bin/env bash

cowsay "`printf 'the cow says, "%s"' "$(cat /dev/stdin)"`"
