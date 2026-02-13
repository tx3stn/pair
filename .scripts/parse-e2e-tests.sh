#!/bin/bash

# Parse e2e test output and create a summary

awk '
BEGIN {
    tests = 0
    passed = 0
    failed = 0
    skipped = 0
}

/^ok/ {
    tests++
    passed++
    print "✅ " $0
}

/^not ok/ {
    tests++
    failed++
    print "❌ " $0
}

/^ok.*# SKIP/ {
    skipped++
    print "⏭️  " $0
}

END {
    print ""
    print "## Summary"
    print "- **Total tests:** " tests
    print "- **Passed:** " passed
    print "- **Failed:** " failed
    if (skipped > 0) {
        print "- **Skipped:** " skipped
    }
}'
