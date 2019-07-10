# Steps to Release

We're following a [semantic versioning](https://semver.org/) approach.


## Create a Release PR

Include all the meaningful changes since the last release into the [CHANGELOG.md](CHANGELOG.md) file.

Please visit [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) for information on why and how changelogs should be kept.

## Merge Release PR

Verify that the acceptance test suite has passed for the release PR, then merge the PR.

## Tag the release

```
git tag -am "terraform-provider-gorillastack_vX.Y.Z_x5 release" vX.Y.Z
git push --tags
```

## Goreleaser

We run [goreleaser](https://goreleaser.com/) via travis, this should run against master on commit. Here is how to run it manually:

```
GITHUB_TOKEN=xxx goreleaser release --rm-dist
```

## See the release in Github

You can find the releases in Github, e.g. [v0.1.0](https://github.com/gorillastack/terraform-provider-gorillastack/releases/tag/v0.1.0).

## Configuring travisci

To add/update build environment vars to the `.travis.yml` file, please utilize the travis cli tool.
```
gem install travis
```

Next, sign in to travis using your github token. (Note, the token only needs privileges to retrieve your email address).
```
travis login --github-token <token> --org
```

Next, encrypt some env vars and add automatically to the `.travis.yml` file
```
travis encrypt --org --add env.global EXAMPLE_VAR=jeff_goldblum
```

Finally, commit and push the changes.
