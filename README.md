# Golang Buster App

## Running the Golang skeleton
The documentation for the Golang Buster App assumes you've read and followed the
steps on the root directory's [README](../README.md) file.

### Running the skeleton with no live re-builds

This method relies solely on `docker-compose` and doesn't require any additional
software.  The downside is that you won't get "auto refresh" capabilities and a
manual container image rebuild and service restart will be required to see your
changes reflected on the running container.

```
docker-compose up --build
```

After doing source modifications:
```
docker-compose up --build --force-recreate -d buster-golang-app
```

### Running the skeleton WITH live automatic re-builds

In order for this to be possible, you'll need to download [Tilt](https://tilt.dev/).

If you're on MacOS and have `brew` available, you can easily install `tilt` like this:
```
brew install tilt-dev/tap/tilt
```

If not on a MacOS or don't have `brew` available you have plenty of
[installation options](https://docs.tilt.dev/install.html).  You can also download
the latest pre-compiled release binary by OS [here](https://github.com/tilt-dev/tilt/releases/).
The point is that all need is to have the `tilt` executable  available in your `PATH`.

Once `tilt` is installed the application can be started like this:

```
tilt up
```

The above command won't return and will keep running.  After changing the Go source
files you will see how the application container gets rebuilt and deployed.

## Examples
You can find examples on how to use the `database` and `busterapi` libraries in the
[examples](examples/) directory.
