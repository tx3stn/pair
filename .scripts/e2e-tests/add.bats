#!/usr/bin/env bats

# e2e tests for the `pair add` command

main_branch='main'
test_branch='bats-tests'
test_dir='/tmp/project-pair-add'

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

# write_config writes the given config (read from stdin) to the pair config file.
write_config() {
	mkdir -p ~/.config
	cat >~/.config/pair.json
}

@test "pair add: accepts the suggested email" {
	git checkout -b "$test_branch"

	write_config <<'EOF'
{
  "accessible": true,
  "coAuthors": {"Alice Smith": "alice@example.com"},
  "commitArgs": "",
  "prefixes": ["fix"],
  "suggestedCoAuthorEmail": "{{.FirstInitial}}{{.LastName}}@example.com",
  "ticketPrefix": ""
}
EOF

	expect_script="$BATS_TEST_DIRNAME/add-suggested.exp"
	run "$expect_script"
	assert_success

	run cat ~/.config/pair.json
	assert_success
	assert_output --partial '"Jane Doe": "jdoe@example.com"'
	# the existing co-author is preserved.
	assert_output --partial '"Alice Smith": "alice@example.com"'
}

@test "pair add: lets you override the suggested email" {
	git checkout -b "$test_branch"

	write_config <<'EOF'
{
  "accessible": true,
  "coAuthors": {"Alice Smith": "alice@example.com"},
  "commitArgs": "",
  "prefixes": ["fix"],
  "suggestedCoAuthorEmail": "{{.FirstInitial}}{{.LastName}}@example.com",
  "ticketPrefix": ""
}
EOF

	expect_script="$BATS_TEST_DIRNAME/add-edit.exp"
	run "$expect_script"
	assert_success

	run cat ~/.config/pair.json
	assert_success
	assert_output --partial '"Jane Doe": "custom@example.com"'
	refute_output --partial "jdoe@example.com"
}

@test "pair add: works without a suggested email format" {
	git checkout -b "$test_branch"

	write_config <<'EOF'
{
  "accessible": true,
  "coAuthors": {"Alice Smith": "alice@example.com"},
  "commitArgs": "",
  "prefixes": ["fix"],
  "ticketPrefix": ""
}
EOF

	expect_script="$BATS_TEST_DIRNAME/add-no-suggestion.exp"
	run "$expect_script"
	assert_success

	run cat ~/.config/pair.json
	assert_success
	assert_output --partial '"Jane Doe": "jane.doe@example.com"'
}

@test "pair add: errors when no config file exists" {
	git checkout -b "$test_branch"

	rm -f ~/.config/pair.json

	run bash -c "unset XDG_CONFIG_DIR; pair add 2>&1"
	assert_failure
	assert_output --partial "no config file found"
}
