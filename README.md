<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/6295/6295417.png" width="100" />
</p>
<p align="center">
    <h1 align="center">SQL-EDITOR</h1>
</p>
<p align="center">
    <em>SQL editor for local sqlite dbs</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/PritishMishraa/sql-editor?style=flat&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/PritishMishraa/sql-editor?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/PritishMishraa/sql-editor?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/PritishMishraa/sql-editor?style=flat&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
		<em>Developed with the software and tools below.</em>
</p>
<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
	<img src="https://img.shields.io/badge/Vite-646CFF.svg?style=flat&logo=Vite&logoColor=white" alt="Vite">
	<img src="https://img.shields.io/badge/React-61DAFB.svg?style=flat&logo=React&logoColor=black" alt="React">
	<img src="https://img.shields.io/badge/TypeScript-3178C6.svg?style=flat&logo=TypeScript&logoColor=white" alt="TypeScript">
</p>
<hr>

##  Quick Links

> - [ Features](#features)
> - [ Getting Started](#getting-started)
>   - [ Installation](#installation)
>   - [ Running sql-editor](#running-sql-editor)
>   - [ Creating executable](#reating-executable)

---

##  Features

- view all sqlite dbs
- connect to any db
- query and view the results in ui

---

##  Getting Started

***Requirements***

Ensure you have the [GO](https://go.dev/doc/install) installed on your system:

* **TypeScript**: `version x.y.z`

###  Installation

1. Clone the sql-editor repository:

```sh
git clone https://github.com/PritishMishraa/sql-editor
```

2. Change to the project directory:

```sh
cd sql-editor && go install
```

```sh
cd web && pnpm i
```

###  Running sql-editor

Use the following command to run sql-editor:

```sh
cd sql-editor && go run .
```

```sh
cd web && pnpm dev
```

###  Creating executable

```sh
go build -o sql-editor .
```

You can throw this binary into any of your projects and query its sqlite dbs.
