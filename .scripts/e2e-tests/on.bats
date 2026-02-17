#!/usr/bin/env bats

# e2e tests for the `pair on` command

main_branch='main'
test_branch='bats-tests'
test_dir='/tmp/project-pair-on'

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

@test "pair on: set ticket id" {
	git checkout -b "$test_branch"
	run pair on TICKET-123
	assert_success

	assert [ -f "/tmp/pair/on" ]
	ticket=$(cat "/tmp/pair/on")
	assert_equal "TICKET-123" "$ticket"
}

@test "pair on: overwrite existing ticket" {
	git checkout -b "$test_branch"
	mkdir -p "/tmp/pair/"
	echo "OLD-TICKET" >"/tmp/pair/on"

	run pair on NEW-TICKET
	assert_success

	ticket=$(cat "/tmp/pair/on")
	assert_equal "NEW-TICKET" "$ticket"
}

@test "pair on: no ticket argument" {
	git checkout -b "$test_branch"

	expect_script="$BATS_TEST_DIRNAME/on-interactive.exp"
	run "$expect_script"
	assert_success

	assert [ -f "/tmp/pair/on" ]
	ticket=$(cat "/tmp/pair/on")
	assert_equal "TEST-999" "$ticket"
}
