#!/usr/bin/env bash

configure-git() {
	git config --global init.defaultBranch "$1"
	git config --global user.email "int-tests@pair.com"
	git config --global user.name "integration tests"
}
