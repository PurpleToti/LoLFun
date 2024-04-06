package ciad

type ExitCode int

const (
	EC_ok ExitCode = iota

	EC_unexpected

	EC_room_created_and_joined
	EC_room_joined
	EC_already_in_room
	EC_room_full
	EC_not_in_room
	EC_room_empty

	EC_room_created
	EC_no_room_created

	EC_user_not_in_room

	EC_room_not_in_map

	EC_bad_username
)
