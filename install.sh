cd installer/
go mod tidy
go build --tags cli
mkdir -p VencordInstaller.app/Contents/MacOS
mkdir -p VencordInstaller.app/Contents/Resources
cp macos/Info.plist VencordInstaller.app/Contents/Info.plist
mv VencordInstaller VencordInstaller.app/Contents/MacOS/VencordInstaller
cp macos/icon.icns VencordInstaller.app/Contents/Resources/icon.icns
rm -rf ../VencordInstaller.app
mv VencordInstaller.app ../VencordInstaller.app

cd ../autovencordpatch
go get github.com/fsnotify/fsnotify
go build -o autovencordpatch autovencordpatch.go
chmod +x autovencordpatch
mv autovencordpatch ../VencordInstaller.app/Contents/Resources/autovencordpatch

cd ..
cp autovencordpatch/org.aaron.autovencordpatch.plist ~/Library/LaunchAgents/org.aaron.autovencordpatch.plist
rm -rf /Applications/VencordInstaller.app
mv VencordInstaller.app /Applications/VencordInstaller.app
launchctl unload ~/Library/LaunchAgents/org.aaron.autovencordpatch.plist > /dev/null 2>&1
launchctl load ~/Library/LaunchAgents/org.aaron.autovencordpatch.plist > /dev/null 2>&1
open /Applications/VencordInstaller.app

echo "Installed AutoVencordPatch"