# gitignore

```text
        _  _    _
  __ _ (_)| |_ (_)  __ _  _ __    ___   _ __   ___
 / _` || || __|| | / _` || '_ \  / _ \ | '__| / _ \
| (_| || || |_ | || (_| || | | || (_) || |   |  __/
 \__, ||_| \__||_| \__, ||_| |_| \___/ |_|    \___|
 |___/             |___/

```

## ✨ Features

- Create `.gitignore` files for specific languages or frameworks.
- List all available templates.

## 🚀 Installation

Binaries for all major platforms are available on the
[releases page](https://github.com/trinhminhtriet/gitignore/releases).

### Manually

You'll need Go for this. You can grab it over at http://golang.org/.

```sh
go get github.com/trinhminhtriet/gitignore
go install github.com/trinhminhtriet/gitignore
```

## 💡 Usage

To create a `.gitignore` file for specific languages or frameworks, use the `gitignore create` command followed by the names of the templates you want to include. For example, to create a `.gitignore` file for Go, Java, and Ruby, you would run:

```sh
gitignore create go java ruby
```

This will generate a `.gitignore` file with the appropriate rules for those languages.

To see a list of all available templates, use the `gitignore list` command:

```sh
gitignore list
```

## 🤝 How to contribute

We welcome contributions!

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m "feat: my new feature"`;
- Push to your branch: `git push origin my-feature`.

Once your pull request has been merged, you can delete your branch.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
