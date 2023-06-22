#!/bin/sh
# test.sh: Validate an answer by comparing the effective to the expected result.
# Copyright © 2023  Savio Sena <savio.sena@acm.org>.

if [ $# -eq 0 ]; then
    printf "Usage: %s <dir>\n" "$(basename $0)"
    exit 127
fi

prepare() {
    challenge="$1"

    if [ ! -d "$challenge" ]; then
        printf "Error: Challenge not found: %s\n" "$challenge" >&2
        exit 128
    fi

    if [ ! -d "$challenge/input" ]; then
        printf "Error: Input directory is missing for challenge: %s\n" "$challenge" >&2
        exit 129
    fi

    if [ ! -d "$challenge/output" ]; then
        printf "Warning: Output directory is missing for challenge: %s: Creating it...\n" "$challenge" >&2
        mkdir -p $challenge/output
    fi

    if [ ! -f "./$challenge/main" ]; then
        make $challenge/main
    fi
}

main() {
    for challenge in "$@"; do
        prepare $challenge
        program=$challenge/main

        for input in "$challenge"/input/*; do
            output=$(dirname $(dirname "$input"))/output/$(basename "$input" | sed -e 's/input/output/g')
            tmpout=$(mktemp)

            $program <$input >$tmpout 2>$tmpout.error

            if [ "$(wc -c $tmpout.error | cut -f1 -d ' ')" != "0" ]; then
                printf "%50s: DEBUG\n%s\n" "$challenge" "$(cat $tmpout.error)" >&2
            fi

            if ! diff "$tmpout" "$output" >$tmpout.diff; then
                printf "%50s: FAILED\n\noutput(%s) = {\n%s\n}\n\nexpected(%s) = {\n%s\n}\n\ndiff = {\n%s\n}\n\n\n" \
                    "$challenge" "$tmpout" "$(cat -e $tmpout)" "$output" "$(cat -e $output)" \
                    "$(cat $tmpout.diff)" >&2
                exit 130
            fi
        done
        printf "%50s: OK!\n" "$challenge" >&2
    done
}

main "$@"
