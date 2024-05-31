# GH SSH Signing

GH SSH Signing is a GitHub CLI extension that allows you to manage SSH Signing keys with your GitHub account.

## Functionality

This CLI extension allows you to:

* Initialize your system to be able to use SSH Signing on your local machine
* Export your local SSH Keys so that GitHub (and other users) can verify your signed commits
    * Note that the default GitHub CLI Token is incapable of doing this; you'll need to mint a Fine Grained Token (see [the GitHub documentation for more info](https://docs.github.com/en/rest/users/ssh-signing-keys?apiVersion=2022-11-28#create-a-ssh-signing-key-for-the-authenticated-user--fine-grained-access-tokens))
* Import SSH keys from any GitHub user so that you can verify their commits locally

## Installation

To install GH SSH Signing, you need to have the GitHub CLI (`gh`) installed. If you don't have it installed, you can follow the installation instructions [here](https://cli.github.com).

Once you have `gh` installed, you can install the GH SSH Signing extension by running the following command:

```bash
gh extensions install rzhade3/gh-ssh-signing
```

### Building from Source

You can also build this extension from source by cloning this repository locally, then running [these instructions](https://docs.github.com/en/github-cli/github-cli/creating-github-cli-extensions#creating-an-interpreted-extension-manually).


