#!/bin/sh

# Parse the e2e test output and create a simple markdown table from the results
# to display as a summary in the GITHUB_STEP_SUMMARY output.

# reads from /dev/stdin so you can pipe tests to the script.
summary=$(cat -)

echo '## 🧪 end to end test results'
echo ''
echo '| status | command | test name |'
echo '| --- | --- | --- |'
echo "$summary" | while IFS= read -r line; do
	if [ "$line" != "${line#1..}" ]; then
		continue
	fi

	case "$line" in
		ok\ *|not\ ok\ *)
			;;
		*)
			continue
			;;
	esac

	status=$(echo "$line" | awk '{print $1}')

	if [ "$status" = "ok" ]; then
		result_icon='✓'
	else
		result_icon='✕'
	fi

	# Parse lines shaped like:
	# ok 10 pair with: select single co-author
	payload=$(echo "$line" | sed -n \
		-e 's/^ok [0-9][0-9]* pair //p' \
		-e 's/^not ok [0-9][0-9]* pair //p')

	cmd_name=${payload%%:*}
	test_name=${payload#*: }

	if [ -z "$cmd_name" ] || [ -z "$test_name" ] || [ "$cmd_name" = "$payload" ]; then
		continue
	fi

	echo "| $result_icon | pair $cmd_name | $test_name |"
done
