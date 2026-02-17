#!/usr/bin/env bats

# e2e tests for the `pair commit` command

main_branch='main'
test_branch='bats-tests'
test_dir='/tmp/project-pair-commit'

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

@test "pair commit: basic commit with ticket and co-authors" {
	git checkout -b "$test_branch"

	mkdir -p "/tmp/pair"
	echo "TICKET-456" >"/tmp/pair/on"
	echo -e '{"name":"Alice Smith","email":"alice@example.com"}\n{"name":"Bob Jones","email":"bob@example.com"}' >"/tmp/pair/with"

	# Make a change to commit
	echo "test change" >>README.md
	git add README.md

	expect_script="$BATS_TEST_DIRNAME/commit-basic.exp"
	run "$expect_script"
	assert_success

	# Check the commit message includes ticket and co-authors
	run git --no-pager log -1 --pretty=format:"%s%n%b"
	assert_success
	assert_output --partial "fix(TICKET-456): test commit message"
	assert_output --partial "Co-authored-by: Alice Smith"
	assert_output --partial "Co-authored-by: Bob Jones"
}

@test "pair commit: prompt for ticket and co-authors when not set" {
	git checkout -b "$test_branch"

	# Make a change to commit
	echo "test change" >>README.md
	git add README.md

	expect_script="$BATS_TEST_DIRNAME/commit-prompt-all.exp"
	run "$expect_script"
	assert_success

	# Check the commit message includes prompted values
	run git --no-pager log -1 --pretty=format:"%s%n%b"
	assert_success
	assert_output --partial "fix(PROMPT-123): test commit message"
	assert_output --partial "Co-authored-by: Alice Smith"
}

@test "pair commit: prompt for co-authors when ticket set" {
	git checkout -b "$test_branch"

	mkdir -p "/tmp/pair"
	echo "TICKET-789" >"/tmp/pair/on"

	# Make a change to commit
	echo "test change" >>README.md
	git add README.md

	expect_script="$BATS_TEST_DIRNAME/commit-prompt-coauthors.exp"
	run "$expect_script"
	assert_success

	# Check the commit message includes ticket and prompted co-authors
	run git --no-pager log -1 --pretty=format:"%s%n%b"
	assert_success
	assert_output --partial "feat(TICKET-789): test commit message"
	assert_output --partial "Co-authored-by: Bob Jones"
}

@test "pair commit: no changes to commit" {
	git checkout -b "$test_branch"

	run pair commit
	assert_failure
	assert_output --partial "error running git commands: exit status 1"
}
