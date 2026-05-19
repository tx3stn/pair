#!/usr/bin/env bats

# e2e tests for the `pair new` command

main_branch='main'
test_branch='bats-tests'
test_dir='/tmp/project-pair-new'

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

@test "pair new: sets co-authors and ticket id" {
	git checkout -b "$test_branch"

	expect_script="$BATS_TEST_DIRNAME/new-basic.exp"
	run "$expect_script"
	assert_success

	assert [ -f "/tmp/pair/with" ]
	run cat "/tmp/pair/with"
	assert_success
	assert_output --partial '{"name":"Alice Smith","email":"alice@example.com"}'

	assert [ -f "/tmp/pair/on" ]
	ticket=$(cat "/tmp/pair/on")
	assert_equal "NEW-123" "$ticket"
}

@test "pair new: overwrites existing session state" {
	git checkout -b "$test_branch"

	mkdir -p "/tmp/pair"
	echo "OLD-TICKET" >"/tmp/pair/on"
	echo '{"name":"old value","email":"old@example.com"}' >"/tmp/pair/with"

	expect_script="$BATS_TEST_DIRNAME/new-basic.exp"
	run "$expect_script"
	assert_success

	run cat "/tmp/pair/with"
	assert_success
	assert_output --partial '{"name":"Alice Smith","email":"alice@example.com"}'
	refute_output --partial "old value"

	ticket=$(cat "/tmp/pair/on")
	assert_equal "NEW-123" "$ticket"
}

@test "pair new: works with no existing state" {
	git checkout -b "$test_branch"

	assert [ ! -f "/tmp/pair/on" ]
	assert [ ! -f "/tmp/pair/with" ]

	expect_script="$BATS_TEST_DIRNAME/new-basic.exp"
	run "$expect_script"
	assert_success

	assert [ -f "/tmp/pair/with" ]
	assert [ -f "/tmp/pair/on" ]
}
