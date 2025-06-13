# 🌀 gitpulse

**gitpulse** is an open-source CLI tool for instantly analyzing Git repository statistics — including contributor activity, commit patterns, and directory heat — right from your terminal.  
It supports both **command-line** and **interactive terminal UI (TUI)** modes for a seamless developer experience.

---

## ✨ Features

- 📊 **Repository Summary**  
  Total commits, unique contributors, first & last commits, average commits/day  
- 🧑‍💻 **Top Contributors**  
  Most active committers, ranked by number of commits  
- 🔥 **Commit Activity Heatmap**  
  Visualize commit frequency by weekday  
- 📂 **Most Active Directories**  
  See which folders change the most  
- 🧪 **TUI Mode**  
  Interactive, keyboard-navigable terminal interface with tabs

---


### 📦 Download

➡️ [Latest Release](https://github.com/Mertozturkk/gitpulse/releases/latest)

## ⚙️ Installation

> Requires **Go 1.21+** installed

### 🧱 Build from Source

#### Linux & macOS

```sh
git clone https://github.com/Mertozturkk/gitpulse.git
cd gitpulse
go build -o gitpulse
```

### Windows

```sh
git clone https://github.com/Mertozturkk/gitpulse.git
cd gitpulse
go build -o gitpulse.exe
```

🚀 Usage

Run from the root of a Git repository:

🔹 Summary

```sh
gitpulse summary
```

🔹 Top Contributors

```sh
gitpulse authors --top 5
```

🔹 Commit Heatmap

```sh
gitpulse heatmap --since 30d
```

🔹 Active Directories

```sh
gitpulse dirs
```

🔹 Interactive Terminal UI (TUI)

```sh
gitpulse tui
```

•	Navigate tabs with ← → arrows
•	Press q to quit


🔹 Help

```sh
gitpulse --help
```

🛠️ Contributing

Pull requests and issues are welcome!
Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute.
1.	Fork the repository
2.	Create a feature branch
3.	Open a pull request
4.	Describe your changes clearly

## License

This project is licensed under the [MIT License](LICENSE).

---

> Developed by [Mertozturkk](https://github.com/Mertozturkk)