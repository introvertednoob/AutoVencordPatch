#!/bin/zsh

# Paths
CACHE_FILE="$HOME/cached.txt"
DISCORD_JSON="/Applications/Discord.app/Contents/Resources/build_info.json"
INSTALLER="/Applications/VencordInstaller.app"

# Ensure cache file exists
if [[ ! -f "$CACHE_FILE" ]]; then
    echo "0.0.0" > "$CACHE_FILE"
fi

cached_version=$(cat "$CACHE_FILE")

# Function to check version and run installer if changed
check_version() {
    if [[ -f "$DISCORD_JSON" ]]; then
        current_version=$(jq -r '.version' "$DISCORD_JSON")
        if [[ "$current_version" != "$cached_version" ]]; then
            cached_version="$current_version"
            echo "$cached_version" > "$CACHE_FILE"
            open "$INSTALLER"
        fi
    else
        echo "Discord not found. Retrying..."
    fi
}

# Initial check
check_version

# Monitor changes with fswatch
fswatch -0 "$DISCORD_JSON" | while read -d "" event; do
    check_version
done