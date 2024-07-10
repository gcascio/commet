<p align="center">
  <picture>
    <img src="https://github.com/gcascio/commet/assets/35739042/289afcf0-4e83-47b4-bb89-c3d0753dd881"  height="120" alt="Banner for Commet">
  </picture>
</p>

## Intro

Commet creates commit messages using Ollama. There are plenty of tools out there that do this but I was looking for something that does not run a pr-commit hook but is just an alternative CLI command to `git commit` which I can use across my projects and integrates into my current workflow.

## How it works

Commet does three things using your local git installation and Ollama:

1. Get the current diff of the staged files (`git diff --cached`)
2. Send the diff with a prompt asking for a commit message to Ollama
3. Commit the changes and opening the editor to potentially edit the message (`git commit -m <the message> -e`)

## Installation

### Homebrew Tap

`brew tap gcascio/commet`

`brew install commet`

### Manually

Download the pre-compiled binaries from the [releases page](https://github.com/gcascio/commet/releases) and copy them to the desired location.

### Requirements

1. Local git installation
2. Access to ollama API (by default at http://localhost:11434) with a model of choice (default: mistral)

## Usage

`commet [flags]`

### Options

```
-a, --all            Commit all changed files i.e. 'git commit -a'
-h, --help           help for commet
    --llm string     URL to LLM API i.e. Ollama, defaults to: http://localhost:11434/api/chat
-m, --model string   LLM model to be used, defaults to: mistral
```

### Configuration

For persistent configurations commet can be configured with a .commet.yaml file in the home directory. E.g.

```yaml
llm: https://mydomain.com
model: mistral
```

## Limitations

Commet works best with small to medium sized changes. Trying to generate a commit message for a large change over multiple files usually results in unsatisfying results. In these cases it is recommended to split the changes into multiple smaller commits.
