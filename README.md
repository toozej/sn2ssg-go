# sn2ssg-go
Convert Simplenote notes to SSG-formatted Markdown files with Golang
This can provide a portion of a fully-automated Simplenote note -> SSG-generated website or blogging pipeline

## Usage
### Convert specifically-tagged Simplenote notes to SSG-formatted Markdown
This is the primary use-case of sn2ssg-go, and can either be done via `make` or as part of a docker-compose project (see [docker-compose example](./docker-compose.yml))
1. build sn2ssg-go Docker image
    - `make build`
2. set sn2ssg-go environment variables
    - `cp .env.sample .env`
    - edit `.env` with correct values for your setup
3. run sn2ssg-go Docker image
    - `make run`

### Output all notes with a specified tag to terminal
Alternatively if you just want to see which notes are tagged with a specific tag, or otherwise download them for another purpose
you can use Docker + sncli package to do so:
1. build sncli Docker image
    - `docker build -t sncli https://github.com/insanum/sncli.git`
2. run sncli
    - `docker run --rm -i -e SN_USERNAME=SIMPLENOTE_USERNAME_GOES_HERE -e SN_PASSWORD=SIMPLENOTE_PASSWORD_GOES_HERE sncli --config=/dev/null -r dump TAG_TO_DOWNLOAD_GOES_HERE`

## changes required to use this as a starter template
- set built packages visibility in GitHub packages to public
    - navigate to https://github.com/users/$USERNAME/packages/container/$REPO/settings
    - scroll down to "Danger Zone"
    - change visibility to public
