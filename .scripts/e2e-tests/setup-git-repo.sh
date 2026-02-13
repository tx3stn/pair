#!/usr/bin/env bash

# script to provision e2e test git repo for pair testing
setup-git-repo-for-pair() {
	mkdir "$1"
	git config --global --add safe.directory "$1"
	cd "$1" || exit

	git init
	echo "# Test Project" > README.md
	git add README.md
	git commit -m "initial commit"
	
	# Create a basic pair config
	mkdir -p ~/.config
	cat > ~/.config/pair.json << EOF
{
  "accessible": false,
  "coAuthors": {
    "Alice Smith": "alice@example.com",
    "Bob Jones": "bob@example.com", 
    "Charlie Brown": "charlie@example.com"
  },
  "commitArgs": "",
  "prefixes": ["fix", "feat", "docs", "test"],
  "ticketPrefix": ""
}
EOF
}
