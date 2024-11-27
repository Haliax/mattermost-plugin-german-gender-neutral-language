# German gender-neutral Language Plugin

This plugin is designed for Mattermost to process and modify messages before they are posted or updated. Specifically, it looks for words containing the segments `*innen` or `*in` (like `Freund*innen`) and escapes the asterisk (`*`) to avoid unintended markdown formatting issues.

## Features

- **Automatic Escaping**: The plugin detects words containing `*innen` or `*in` and replaces `*` with its escaped version `\*`.
- **Prevents Double Escaping**: Avoids modifying words where the asterisk is already escaped.
- **Message Hooks**: Works for both newly posted and updated messages.


## Example

A user posts the message:
```aiignore
Hallo Freund*innen und Kolleg*innen, wie geht es euch?
```

It automatically gets modified to:
```aiignore
Hallo Freund\*innen und Kolleg\*innen, wie geht es euch?
```

This prevents unwanted markdown formatting when such words appear multiple times in a sentence.


## Usage

Once installed and enabled, the plugin will automatically process messages according to the following rules:

1. **Message Posting**: When a message is posted containing words like `Freund*innen`, it will be modified to `Freund\*innen`.
2. **Message Updating**: Similarly, when a message is edited, the same processing will apply.


## License

This project is licensed under the terms of the Apache License 2.0.
