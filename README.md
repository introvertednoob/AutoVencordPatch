# Vencord Auto Updater (macOS)
An efficient program which patches Vencord whenever Discord updates on macOS.</br>
On macOS, Vencord doesn't automatically patch itself when Discord updates, so this is a fix for that.

# What this code does
- Builds a modified CLI installer (.app) which can install Vencord without any user interaction
- Injects a ZSH script which automatically patches your Discord app when it updates
    - This can be located inside the installer .app, in VencordInstaller.app/Contents/Resources/
- Adds a login item to make the ZSH script run on startup
**NOTE**: You can also run VencordInstaller.app directly to force-patch Vencord into Discord.

# Requirements
All original requirements for building the official installer apply here.</br>
Go 1.25 is also recommended, but you can probably use a lower version instead.

Created by [Aaron Wijesinghe](https://github.com/introvertednoob)</br>

## Credits
Auto-patcher created by [Aaron Wijesinghe](https://github.com/introvertednoob)

This software uses the [Vencord Installer](https://github.com/Vencord/Installer)</br>
Copyright (c) 2023 Vendicated and Vencord contributors</br>
Licensed under the GNU General Public License v3.0</br>
