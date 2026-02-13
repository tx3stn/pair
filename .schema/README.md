# Schema

The documents in this directory define the [JSON schema](https://json-schema.org/)
for the `pair` configuration file.

You can annotate your config file with the `$schema` keyword for a better
editing experience (provided your editor integrates with JSON schema tooling).

e.g.:

```json
{
  "$schema": "https://raw.githubusercontent.com/tx3stn/pair/main/.schema/schema.json",
  "accessible": true,
  "coAuthors": {
    "Alice Smith": "alice@example.com"
  },
  "commitArgs": "",
  "prefixes": ["fix", "feat", "docs", "test"],
  "ticketPrefix": "TICKET-"
}
```

## Examples

An [example file](./pair.json) has been added to display what the
configuration should look like.
