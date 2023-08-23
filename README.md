# Go Security Assessment Tool (`go-sec-assess`)

A command-line tool written in Go to perform basic security assessments on websites. This tool fetches and analyzes HTTP security headers and checks the TLS configuration of a given website.

## Features

- **TLS Version & Cipher Checker**: Determines the TLS version and cipher suite used by the website.
- **Header Checker**: Checks for the presence of key security headers.
- **URL Validation**: Validates the provided URL to ensure it's well-formed and uses a valid scheme.

## Installation & Usage

### Using Pre-built Binaries

1. **Download the Binary**:
    - Navigate to the [releases](https://github.com/indranandjha1993/go-sec-assess/releases) page of the repository.
    - Look for the latest release and download the appropriate binary for your operating system (e.g., `go-sec-assess-linux-amd64` for Linux).

2. **Make the Binary Executable (Linux/macOS)**:
    ```bash
    chmod +x go-sec-assess-linux-amd64
    ```

3. **Run the Tool**:
    ```bash
    ./go-sec-assess-linux-amd64 -url=https://example.com
    ```
   For Windows, you'd execute the `.exe` file in the command prompt or PowerShell.

Replace `go-sec-assess-linux-amd64` with the name of the downloaded binary and `https://example.com` with the URL of the website you want to assess.

## Output

Information about the TLS configuration of the site (version and cipher suite) and the presence (or absence) of important security headers will be displayed.

## Contributing

### Making Code Contributions

1. Fork the repository.
2. Clone your fork: `git clone https://github.com/indranandjha1993/go-sec-assess.git`.
3. Create a new branch for your features or bug fixes: `git checkout -b your-feature-branch`.
4. Make your changes and commit them: `git commit -am "Add some feature"`.
5. Push your branch: `git push origin your-feature-branch`.
6. Open a Pull Request from your forked repository to `indranandjha1993/go-sec-assess`.

### Pushing a New Release (Maintainers)

1. Merge the desired changes to the `main` branch.
2. Update version numbers or relevant documentation if needed.
3. Commit your changes: `git commit -am "Prepare for vX.X.X release"`.
4. Tag the release: `git tag vX.X.X`.
5. Push the tag: `git push origin vX.X.X`.

