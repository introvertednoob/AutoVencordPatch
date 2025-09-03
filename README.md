# Vencord Auto Patcher (macOS)
An efficient program which patches Vencord whenever Discord updates on macOS.</br>
On macOS, Vencord doesn't automatically patch itself when Discord updates, so this is a fix for that.

## Features
- Patches Vencord automatically, even through Discord updates
- VencordInstaller.app can patch Vencord without any user interaction, unlike the official installer
    - This is due to modifications made to the installer source.
- Notifications are used to communicate success, failure, and errors
- Very efficient, with <=0.1% of CPU and <5MB of RAM being used at idle

## Requirements
All original requirements for building the official installer apply here.</br>
Go 1.25 is also recommended, but you can probably use a lower version instead.

## Credits
Auto-patcher created by [Aaron Wijesinghe](https://github.com/introvertednoob)

This software uses a modified version of the [Vencord Installer](https://github.com/Vencord/Installer)</br>
Copyright (c) 2023 Vendicated and Vencord contributors</br>
Licensed under the GNU General Public License v3.0</br>
