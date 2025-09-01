cd installer/
go mod tidy
go build --tags cli
mkdir -p VencordInstaller.app/Contents/MacOS
mkdir -p VencordInstaller.app/Contents/Resources
cp macos/Info.plist VencordInstaller.app/Contents/Info.plist
mv VencordInstaller VencordInstaller.app/Contents/MacOS/VencordInstaller
cp macos/icon.icns VencordInstaller.app/Contents/Resources/icon.icns
mv VencordInstaller.app ../VencordInstaller.app

cd ..
cp vencordchecker/autovencordupdate.sh VencordInstaller.app/Contents/Resources/autovencordupdate.sh
cp vencordchecker/org.aaron.autovencordupdate.plist ~/Library/LaunchAgents/org.aaron.autovencordupdate.plist
rm -rf /Applications/VencordInstaller.app
mv VencordInstaller.app /Applications/VencordInstaller.app
launchctl load ~/Library/LaunchAgents/org.aaron.autovencordupdate.plist

echo "Installed AutoVencordUpdate"