#!/usr/bin/env bash
set -euo pipefail
command -v stty >/dev/null
command -v uname >/dev/null
normalize() {
	local raw="$1" iflag oflag cflag lflag rest
	if [[ "$(uname -s)" != Darwin ]]; then
		printf '%s' "$raw"
		return
	fi
	# Darwin may set PENDIN while retyping pending input after raw mode.
	# It is kernel bookkeeping, not a user-configured terminal mode.
	IFS=: read -r iflag oflag cflag lflag rest <<<"$raw"
	printf '%s:%s:%s:%x:%s' "$iflag" "$oflag" "$cflag" "$((16#$lflag & ~0x20000000))" "$rest"
}
before="$(stty -g)"
before_details="$(stty -a)"
status=0
"$WK_BINARY" || status=$?
after="$(stty -g)"
if [[ "$(normalize "$before")" == "$(normalize "$after")" ]]; then
	printf '\r\nTERMINAL_RESTORED\r\n'
else
	printf '\r\nTERMINAL_CHANGED\r\nBEFORE=%s\r\nAFTER=%s\r\n%s\r\n' "$before" "$after" "$before_details"
	stty -a
	exit 97
fi
exit "$status"
