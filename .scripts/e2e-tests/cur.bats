#!/usr/bin/env bats

# e2e tests for the `pair cur` command

main_branch='main'
test_branch='bats-tests'
test_dir='/tmp/project-pair-cur'

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

@test "pair cur: no active session" {
	git checkout -b "$test_branch"

	run bash -c "pair cur 2>&1"
	assert_success
	assert_output --partial "no active pairing session"
}

@test "pair cur: shows ticket and co-authors" {
	git checkout -b "$test_branch"

	mkdir -p "/tmp/pair"
	echo "TICKET-123" >"/tmp/pair/on"
	echo '{"name":"Alice Smith","email":"alice@example.com"}' >"/tmp/pair/with"

	run bash -c "pair cur 2>&1"
	assert_success
	assert_output --partial "pairing on"
	assert_output --partial "ticketID=TICKET-123"
	assert_output --partial "with 1 coauthors"
	assert_output --partial "Alice Smith"
}

@test "pair cur: shows ticket only when no co-authors set" {
	git checkout -b "$test_branch"

	mkdir -p "/tmp/pair"
	echo "TICKET-456" >"/tmp/pair/on"

	run bash -c "pair cur 2>&1"
	assert_success
	assert_output --partial "pairing on"
	assert_output --partial "ticketID=TICKET-456"
	assert_output --partial "with 0 coauthors"
}

@test "pair cur: shows co-authors only when no ticket set" {
	git checkout -b "$test_branch"

	mkdir -p "/tmp/pair"
	echo -e '{"name":"Alice Smith","email":"alice@example.com"}\n{"name":"Bob Jones","email":"bob@example.com"}' >"/tmp/pair/with"

	run bash -c "pair cur 2>&1"
	assert_success
	assert_output --partial "pairing on"
	assert_output --partial "with 2 coauthors"
	assert_output --partial "Alice Smith"
	assert_output --partial "Bob Jones"
}
