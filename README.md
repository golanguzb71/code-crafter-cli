# CodeCrafter

## Overview

The **CodeCrafter** CLI tool is designed to streamline the creation and management of project structures. Ideal for developers working on monolithic applications, CodeCrafter automates the setup and organization of your project's directory and file structure.

<p>Join our community on Telegram for updates, tips, and more: <a href="https://t.me/kuyov_taraf" target="_blank">https://t.me/kuyov_taraf</a></p>

## Features

<ul>
    <li><strong>Generate Linking Classes:</strong> Automatically create linking classes based on your project's requirements.</li>
    <li><strong>Project Structure Setup:</strong> Establish a comprehensive project structure, including directories, files, and configurations.</li>
    <li><strong>YAML Configuration:</strong> Use a <code>codeCrafter.yml</code> file to efficiently manage your project's structure.</li>
    <li><strong>Shell Autocompletion:</strong> Enhance your development workflow with autocompletion scripts for various shells.</li>
</ul>

## Installation

### Download the Binary

<pre><code>wget https://github.com/golanguzb71/codeCrafter/releases/download/v1.0.0/codeCrafter-linux-amd64 -O /usr/local/bin/codeCrafter
chmod +x /usr/local/bin/codeCrafter</code></pre>

### Clone from GitHub

<pre><code>git clone https://github.com/golanguzb71/codeCrafter.git
cd codeCrafter</code></pre>

<pre><code>make build</code></pre>

<pre><code>sudo install -m 0755 bin/codeCrafter /usr/local/bin/codeCrafter</code></pre>

## Usage

<pre><code>codeCrafter
Tool for helping make project structures. Join my Telegram channel => https://t.me/kuyov_taraf

Usage:
  codeCrafter [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  gen         Generate linking classes
  genCrafter  Generate codeCrafter.yml file for project structure
  genFoFI     Generate folders and files declared in codeCrafter.yml
  help        Help about any command
  version     Print the version number</code></pre>

## Setup

<ol>
    <li><strong>Generate a YAML Configuration:</strong> <code>codeCrafter genCrafter</code></li>
    <li><strong>Generate Project Files and Folders:</strong> <code>codeCrafter genFoFI</code></li>
    <li><strong>Generate Linking Classes:</strong> <code>codeCrafter gen</code></li>
    <li><strong>Enable Autocompletion:</strong> <code>codeCrafter completion [bash|zsh|fish]</code></li>
</ol>

## Contributing

We welcome contributions to CodeCrafter! To get started, please refer to our <a href="CONTRIBUTING.md" target="_blank">contributing guidelines</a>.

## License

CodeCrafter is licensed under the MIT License. See the <a href="LICENSE" target="_blank">LICENSE</a> file for details.
