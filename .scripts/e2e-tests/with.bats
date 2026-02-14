#!/usr/bin/env bats

# e2e tests for the `pair with` command

main_branch='main'
test_branch='bats-tests'
test_dir='/tmp/project-pair-with'

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

@test "pair with: select single co-author" {
	git checkout -b "$test_branch"

	expect_script="$BATS_TEST_DIRNAME/with-single.exp"
	run "$expect_script"
	assert_success

	# Check that coauthors file was created
	today=$(date +%Y-%m-%d)
	assert [ -f "/tmp/pair/$today/with" ]
	run cat "/tmp/pair/$today/with"
	assert_success
	assert_output --partial '{"name":"Alice Smith","email":"alice@example.com"}'
}

@test "pair with: select multiple co-authors" {
	git checkout -b "$test_branch"

	expect_script="$BATS_TEST_DIRNAME/with-multiple.exp"
	run "$expect_script"
	assert_success

	today=$(date +%Y-%m-%d)
	run cat "/tmp/pair/$today/with"
	assert_success
	# Should contain both authors (order may vary)
	assert_output --partial "Alice Smith"
	assert_output --partial "Bob Jones"
}

@test "pair with: overwrite existing co-authors" {
	git checkout -b "$test_branch"
	today=$(date +%Y-%m-%d)
	mkdir -p "/tmp/pair/$today"
	echo "old-author" >"/tmp/pair/$today/with"

	expect_script="$BATS_TEST_DIRNAME/with-overwrite.exp"
	run "$expect_script"
	assert_success

	run cat "/tmp/pair/$today/with"
	assert_success
	assert_output --partial '{"name":"Charlie Brown","email":"charlie@example.com"}'
}

@test "pair with: no co-author argument" {
	git checkout -b "$test_branch"

	expect_script="$BATS_TEST_DIRNAME/with-none.exp"
	run "$expect_script"
	assert_failure
}
