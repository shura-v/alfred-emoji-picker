# Alfred emoji picker

Input emojis from Alfred — at a blazing-fast speed!

<p align="center">
    <img src="https://user-images.githubusercontent.com/2824100/174484132-c76cf892-27e8-4d8c-bec7-76745016fe1a.png" data-canonical-src="https://user-images.githubusercontent.com/2824100/174484132-c76cf892-27e8-4d8c-bec7-76745016fe1a.png" width="400"/>
</p>

# Install

- Download the workflow from the [latest release](https://github.com/devnoname120/alfred-emoji-picker/releases/latest).
- Open the file and import it into Alfred.
- **Click on the workflow in Alfred and define a hotkey**.

👉 I recommend using <kbd>Command ⌘</kbd> <kbd>Control ⌃</kbd> <kbd>Space</kbd>

# Build

```shell
go install

./build.sh
```

Copy the executable in the Alfred workflow directory and export the new workflow from Alfred.

# Frequently used emoji

The binary now supports a small local usage database in Alfred's workflow data directory:

- Empty query shows only emoji the user has actually selected before.
- The amount shown on empty query is controlled by the Alfred variable `frequent_emoji_limit`.
- Matching results use usage frequency as a tiebreaker, so your go-to emoji gradually bubble up.

To persist usage counts, call the workflow binary once when the user selects an emoji:

```shell
./alfred-emoji-picker --record "{query}"
```

In Alfred, wire this as a separate `Run Script` / `External Trigger` step before the existing paste/copy action, and pass the selected emoji character as `{query}`.

# Update emojis

1) Emoji metadata (names, slugs, keywords) lives in the [`turtle`](https://github.com/devnoname120/turtle) module. Check the README on how to regenerate `emojis.go` to make it up-to-date, and push a new tag.

2) Bump the dependency of `alfred-emoji-picker` via `go get github.com/devnoname120/turtle@latest`.

3) Re-render the PNGs from the current Apple Color Emoji font and re-optimize them (macOS only; requires [uv](https://docs.astral.sh/uv/), `pngquant`, and `oxipng`):

    ```shell
    uv run --project scripts scripts/emojis-generator.py
    ./gen-emojis.sh
    ```

# TODO

- [x] Restore clipboard after pasting emoji
- [x] Support frequently used emoji
- [ ] Support for multiple words fuzzy search
- [x] Add scoring on results (exact match > partial match at beginning > partial match > keywords, categories, etc…)
- [x] Add scripts to update the emoji database
- [x] Support for auto-updates
- [ ] Add content for the "About this Workflow" tab in the config builder
- [ ] Support for skin tones (note: `Turtle` doesn't support them)
