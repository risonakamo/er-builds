# How to create Release
1. Update `doc/for-release/readme.md` with version number
2. Edit `gen-release.sh` with version number
3. Edit `er-builds-web/package.json` with version number
4. `bash gen-release.sh`
5. Take generated dir and zip it up
6. git commit, tag, upload
7. Create github release and upload the zip