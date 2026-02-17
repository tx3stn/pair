#!/usr/bin/env bats

# e2e tests for the `pair done` command

main_branch='main'
test_branch='bats-tests'
test_dir='/tmp/project-pair-done'

setup_file() {
	echo "### suite setup ###"
	load ./setup-git.sh
	configure-git "$main_branch"

	load ./setup-git-repo.sh
	setup-git-repo-for-pair "$test_dir"
}

teardown_file() {
	echo "### suite teardown ###"
	rm -rf "$test_dir"
}

setup() {
	echo "### test setup ###"
	bats_load_library bats-support
	bats_load_library bats-assert
	cd "$test_dir" || exit 1
}

teardown() {
	echo "### test teardown ###"
	load ./teardown-git.sh
	tidy-git-changes "$main_branch" "$test_branch"
	cleanup-pair-state
}

@test "pair done: clear ticket and co-authors" {
	git checkout -b "$test_branch"

	mkdir -p "/tmp/pair"
	echo "TICKET-123" >"/tmp/pair/on"
	echo "Alice Smith <alice@example.com>" >"/tmp/pair/with"

	run pair done
	assert_success

	assert [ ! -f "/tmp/pair/on" ]
	assert [ ! -f "/tmp/pair/with" ]
}

@test "pair done: no existing state files" {
	git checkout -b "$test_branch"

	run pair done
	assert_success

	assert [ ! -f "/tmp/pair/on" ]
	assert [ ! -f "/tmp/pair/with" ]
}
