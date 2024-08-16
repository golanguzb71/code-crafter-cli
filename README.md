<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CodeCrafter</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            margin: 20px;
            color: #333;
        }
        h1, h2 {
            color: #007bff;
        }
        code {
            background: #f8f9fa;
            padding: 2px 4px;
            border-radius: 4px;
        }
        pre {
            background: #f8f9fa;
            padding: 10px;
            border-radius: 4px;
            overflow-x: auto;
        }
        a {
            color: #007bff;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <h1>CodeCrafter</h1>
    <p><strong>CodeCrafter</strong> is a powerful CLI tool designed to streamline the creation and management of project structures. Ideal for developers working on monolithic applications, CodeCrafter automates the setup and organization of your project's directory and file structure.</p>
    <p>Join our community on Telegram for updates, tips, and more: <a href="https://t.me/kuyov_taraf" target="_blank">https://t.me/kuyov_taraf</a></p>

    <h2>Features</h2>
    <ul>
        <li><strong>Generate Linking Classes</strong>: Automatically create linking classes based on your project's requirements.</li>
        <li><strong>Project Structure Setup</strong>: Establish a comprehensive project structure, including directories, files, and configurations.</li>
        <li><strong>YAML Configuration</strong>: Use a <code>codeCrafter.yml</code> file to efficiently manage your project's structure.</li>
        <li><strong>Shell Autocompletion</strong>: Enhance your development workflow with autocompletion scripts for various shells.</li>
    </ul>

    <h2>Installation</h2>

    <h3>Download the Binary</h3>
    <ol>
        <li><strong>Download the Binary</strong>:
            <pre><code>wget https://github.com/golanguzb71/codeCrafter/releases/download/v1.0.0/codeCrafter-linux-amd64 -O /usr/local/bin/codeCrafter
chmod +x /usr/local/bin/codeCrafter</code></pre>
        </li>
        <li><strong>Verify Installation</strong>:
            <pre><code>codeCrafter version</code></pre>
        </li>
    </ol>

    <h3>Clone from GitHub</h3>
    <ol>
        <li><strong>Clone the Repository</strong>:
            <pre><code>git clone https://github.com/golanguzb71/codeCrafter.git
cd codeCrafter</code></pre>
        </li>
        <li><strong>Build the Project</strong>:
            <pre><code>make build</code></pre>
        </li>
        <li><strong>Install the Binary</strong>:
            <pre><code>sudo install -m 0755 bin/codeCrafter /usr/local/bin/codeCrafter</code></pre>
        </li>
    </ol>

    <h2>Usage</h2>
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

    <h2>Setup</h2>
    <ol>
        <li><strong>Generate a YAML Configuration</strong>:
            <pre><code>codeCrafter genCrafter</code></pre>
        </li>
        <li><strong>Generate Project Files and Folders</strong>:
            <pre><code>codeCrafter genFoFI</code></pre>
        </li>
        <li><strong>Generate Linking Classes</strong>:
            <pre><code>codeCrafter gen</code></pre>
        </li>
        <li><strong>Enable Autocompletion</strong>:
            <pre><code>codeCrafter completion [bash|zsh|fish]</code></pre>
        </li>
    </ol>

    <h2>Contributing</h2>
    <p>We welcome contributions to CodeCrafter! To get started, please refer to our <a href="CONTRIBUTING.md" target="_blank">contributing guidelines</a>.</p>

    <h2>License</h2>
    <p>CodeCrafter is licensed under the MIT License. See the <a href="LICENSE" target="_blank">LICENSE</a> file for details.</p>
</body>
</html>
