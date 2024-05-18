# How to create Release
1. Update `version.md` with version number
2. Edit `gen-release.sh` with version number
3. Ensure web repo is on a tagged version
4. `bash gen-release.sh`
5. Take generated dir and zip it up
6. git commit, tag, upload
7. Create github release and upload the zip