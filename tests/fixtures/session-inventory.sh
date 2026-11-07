#!/usr/bin/env bash
set -euo pipefail

printf '%s\n' "$*" >>"${SESSION_INVENTORY_LOG:?}"
case "$*" in
*' hosts --json')
	warnings='[]'
	if [[ "${SESSION_INVENTORY_WARNING:-}" == "1" ]]; then
		warnings='[{"message":"cached home inventory","from":"saved-home"}]'
	fi
	printf '{"home":{},"hosts":[{"host":{"id":"host:00000000-0000-0000-0000-000000000001","label":"studio","endpoints":["studio.example"]}}],"warnings":%s}\n' "$warnings"
	;;
*' api call -- session.list')
	if [[ "$2" == "studio.example" ]]; then
		if [[ "${SESSION_REMOTE_UNAVAILABLE:-}" == "1" ]]; then
			exit 7
		fi
		if [[ "${SESSION_REMOTE_EMPTY:-}" == "1" ]]; then
			printf '%s\n' '{"sessions":[]}'
		else
			printf '%s\n' '{"sessions":[{"session_id":"session:00000000-0000-0000-0000-000000000002","label":"Editor"}]}'
		fi
	else
		printf '%s\n' '{"sessions":[{"session_id":"session:00000000-0000-0000-0000-000000000001","label":"Local"}]}'
	fi
	;;
*' session inspect --json -- session:00000000-0000-0000-0000-000000000002')
	[[ "$2" == "studio.example" ]] || exit 8
	[[ "${SESSION_REMOTE_GONE:-}" != "1" ]] || exit 9
	printf '%s\n' '{"session_id":"session:00000000-0000-0000-0000-000000000002","label":"Renamed after listing","windows":[]}'
	;;
*)
	printf 'unexpected fixture arguments: %s\n' "$*" >&2
	exit 10
	;;
esac
