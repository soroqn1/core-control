# Core Control MVP — Implementation Plan (macOS & Linux)

This document outlines the key functionality of the Minimum Viable Product (MVP) for the cross-platform monitoring and management system.

### 1. System Health Monitoring

The bedrock of the app, providing real-time insights into the hardware state.

- **What it does:** Collects data on CPU load, used RAM, and network activity (incoming/outgoing traffic).
- **Tech stack:** Go (`gopsutil`) for data collection (natively supports both macOS and Linux), Angular for visualization.

### 2. Activity Monitor (Processes)

A deep dive into resource consumption.

- **What it does:** Displays the top processes by CPU and RAM usage. The key feature is **History**: you can track hungry apps through 24-hour and 7-day statistics.
- **How it works:** Process data is saved to a local **PostgreSQL** database. A background worker takes a snapshot every 30 minutes.

### 3. Docker API (Container Manager)

Manage containers without opening Docker Desktop or a terminal.

- **What it does:** Displays a list of all containers, their statuses, and ports. Allows for one-click **Start / Stop / Restart** and live streaming logs.
- **How it works:** Connects to the local **Docker Unix Socket** (common for both macOS and Linux).

### 4. Quick Ops (Terminal and Macros)

- **What it does:**
  - **Terminal:** A fully functional console.
  - **Macros:** Trigger predefined actions (e.g., "Clear Cache", "Dev Mode").
- **How it works:** Implements a **Pseudo Terminal (PTY)**. Macros are stored in JSON/YAML and executed via `os/exec` using system-specific shells (Zsh/Bash).

### 5. App Control (Platform Specific)

Monitoring what is actually running on your machine.

- **What it does:** Distinguishes between system services and user-facing applications. Displays uptime and active window status.
- **How it works:**
  - **macOS:** Uses native APIs (like `NSWorkspace`) via CGO.
  - **Linux:** Uses X11/Wayland APIs or command-line tools (like `wmctrl` or `xprop`) to detect active windows and GUI apps.

---

# Core Control MVP — План реализации (macOS и Linux)

Этот документ описывает ключевой функционал MVP (Minimum Viable Product) для кроссплатформенной системы мониторинга и управления.

### 1. Мониторинг системных ресурсов (System Health)

Основа приложения, предоставляющая информацию о состоянии «железа» в реальном времени.

- **Что делает:** Собирает данные о загрузке процессора, занятой оперативной памяти и сетевой активности (входящий/исходящий трафик).
- **Технологический стек:** Go (`gopsutil`) для сбора данных (поддерживает и macOS, и Linux), Angular для визуализации.

### 2. Мониторинг активности (Процессы)

Глубокий анализ потребления ресурсов.

- **Что делает:** Отображает топ процессов по загрузке CPU и RAM. Ключевая фишка — **История**: возможность отслеживать «прожорливые» приложения не только в моменте, но и через статистику за 24 часа и 7 дней.
- **Как работает:** Данные о процессах сохраняются в локальную базу данных **PostgreSQL**. Фоновый воркер делает «слепок» системы каждые 30 минут.

### 3. Docker API (Управление контейнерами)

Управление контейнерами без открытия Docker Desktop или терминала.

- **Что делает:** Отображает список всех контейнеров, их статусы и порты. Позволяет одной кнопкой выполнять **Start / Stop / Restart**, а также просматривать логи в реальном времени.
- **Как работает:** Бекенд подключается к локальному **Docker Unix Socket** (общий для macOS и Linux).

### 4. Quick Ops (Терминал и Макросы)

- **Что делает:**
  - **Терминал:** Полноценная консоль.
  - **Макросы:** Кнопки-заготовки для частых действий (например, «Очистить кэш», «Режим разработки»).
- **Как работает:** Реализация **Pseudo Terminal (PTY)**. Макросы хранятся в формате JSON/YAML и выполняются через `os/exec` с использованием системных оболочек (Zsh/Bash).

### 5. Управление приложениями (Специфично для платформ)

Контроль за тем, что реально запущено на вашем компьютере.

- **Что делает:** Разделяет системные сервисы и пользовательские приложения. Отображает время работы (uptime) и статус активного окна.
- **Как работает:**
  - **macOS:** Использует нативные API (например, `NSWorkspace`) через CGO.
  - **Linux:** Использует API X11/Wayland или консольные инструменты (например, `wmctrl` или `xprop`) для определения активных окон и GUI-приложений.
