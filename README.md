# pair

## commands

### `on`

Set ticket id.

> [!tip]
> Set `ticketPrefix` in your config file if your tickets always have the same prefix,
> and you can just type the part after the prefix.

You can pass the ticket directly to the command, or just run `pair on` for an interactive
prompt to enter the ticket id.

### `with`

Select co-authors from the list defined in your config file.

### `commit`

Commit your staged changes using the values you have already set.
If you have not set a ticket id or selected co-authors you will be prompted to
set them before running the commit.

> [!tip]
> Want to sign your commits?
> Set `-s` in the `commitArgs` field in your config file.
