# 🤝 Contributing to gitpulse

Thanks for your interest in contributing to **gitpulse**!  
We welcome all kinds of contributions — from bug reports and documentation fixes to feature suggestions and code improvements.

---

## 🧭 How to Contribute

1. **Fork** the repository
2. **Clone** your fork locally:

```bash
git clone https://github.com/your-username/gitpulse.git
cd gitpulse
```

3. **Create** a new branch for your changes:

```bash
git checkout -b feature/your-feature-name
```

4.	Make your changes
5.	Test your changes (go test ./... if applicable)
6.	Commit your changes with a clear message:

```bash
git commit -m "feat: add awesome feature"
```

7.	Push to your fork:

```bash
git push origin feature/your-feature-name
```

8.	Create a Pull Request from your fork to the main branch of gitpulse

✅ Contribution Guidelines

    •	Follow the Go code style conventions

    •	Keep your changes focused and atomic

    •	Avoid unrelated formatting or refactoring in the same PR

    •	Use descriptive commit messages (prefix with feat:, fix:, refactor: etc.)

    •	Test your code (manually or via go test) before pushing



📋 Issues and Suggestions

Found a bug? Have a feature request?
Feel free to open an issue with:

    •  ✅ Clear title

    •  📄 Steps to reproduce (if a bug)

    •  💡 Motivation or use case (if a feature request)

    •  📷 Screenshots or logs (if helpful)


⸻

📚 Project Structure

    • main.go – CLI entry point

    • cmd/ – Command definitions (summary, authors, dirs, etc.)

    • ui/ – TUI (Terminal UI) logic

    • utils/ – Shared utilities

    • go.mod / go.sum – Dependencies and versions
    

⸻

🛡️ Code of Conduct

Please be respectful in your communication and maintain a friendly and inclusive environment for all contributors.

⸻

🙌 Thank You

Whether you’re helping fix a typo or implementing a complex feature — we appreciate your time, effort, and contribution to making gitpulse better!