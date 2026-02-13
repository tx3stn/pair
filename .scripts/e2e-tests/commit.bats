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

# FIXME: expect script is getting stuck
# @test "pair commit: basic commit with ticket and co-authors" {
# 	git checkout -b "$test_branch"
#
# 	# Set up pair state in /tmp/pair/<DATE>/
# 	today=$(date +%Y-%m-%d)
# 	mkdir -p "/tmp/pair/$today"
# 	echo "TICKET-456" > "/tmp/pair/$today/on"
# 	echo -e "Alice Smith <alice@example.com>\nBob Jones <bob@example.com>" > "/tmp/pair/$today/with"
#
# 	# Make a change to commit
# 	echo "test change" >> README.md
# 	git add README.md
#
# 	expect_script="$BATS_TEST_DIRNAME/commit-basic.exp"
# 	run "$expect_script"
# 	assert_success
#
# 	# Check the commit message includes ticket and co-authors
# 	run git log -1 --pretty=format:"%s%n%b"
# 	assert_success
# 	assert_output --partial "TICKET-456"
# 	assert_output --partial "test commit message"
# 	assert_output --partial "Co-authored-by:"
# }
#
# @test "pair commit: prompt for ticket and co-authors when not set" {
# 	git checkout -b "$test_branch"
#
# 	# Make a change to commit
# 	echo "test change" >>README.md
# 	git add README.md
#
# 	expect_script="$BATS_TEST_DIRNAME/commit-prompt-all.exp"
# 	run "$expect_script"
# 	assert_success
#
# 	# Check the commit message includes prompted values
# 	run git log -1 --pretty=format:"%s%n%b"
# 	assert_success
# 	assert_output --partial "PROMPT-123"
# 	assert_output --partial "test commit message"
# 	assert_output --partial "Co-authored-by:"
# }
#
# @test "pair commit: prompt for co-authors when ticket set" {
# 	git checkout -b "$test_branch"
#
# 	# Set up only ticket
# 	today=$(date +%Y-%m-%d)
# 	mkdir -p "/tmp/pair/$today"
# 	echo "TICKET-789" >"/tmp/pair/$today/on"
#
# 	# Make a change to commit
# 	echo "test change" >>README.md
# 	git add README.md
#
# 	expect_script="$BATS_TEST_DIRNAME/commit-prompt-coauthors.exp"
# 	run "$expect_script"
# 	assert_success
#
# 	# Check the commit message includes ticket and prompted co-authors
# 	run git log -1 --pretty=format:"%s%n%b"
# 	assert_success
# 	assert_output --partial "TICKET-789"
# 	assert_output --partial "test commit message"
# 	assert_output --partial "Co-authored-by:"
# }
#
# @test "pair commit: no changes to commit" {
# 	git checkout -b "$test_branch"
#
# 	run pair commit
# 	assert_failure
# 	assert_output --partial "nothing to commit"
# }
