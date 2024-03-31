package ciad

type ExitCode int

const (
	EC_ok ExitCode = iota

	EC_unexpected

	EC_room_created_and_joined
	EC_room_joined
	EC_already_in_room
	EC_room_full

	EC_room_created
	EC_no_room_created

	EC_room_not_in_map
)
