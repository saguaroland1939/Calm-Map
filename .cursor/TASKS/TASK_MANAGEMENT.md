# Markdown-Based Task Management

How to use markdown files to manage tasks in your development workflow.

## Why Markdown?

- **Version controlled**: Tasks live in git alongside code
- **Simple**: No special tools, just text files
- **AI-friendly**: Agents can read and update tasks easily
- **Portable and searchable**: Works anywhere

Best for solo/small team projects. For large teams, consider GitHub Issues or dedicated tools.

## Three Levels of Organization

### TODO.md -- Active Tasks
Quick daily task list with priorities (P0-P3) and status (todo, in-progress, done, blocked). See `TODO.md` for the live example.

### PROJECTS.md -- Long-term Planning
Larger projects broken into milestones. See `PROJECTS.md` for the live example.

### TASKS/ -- Detailed Task Docs
Individual files for complex tasks needing acceptance criteria, implementation plans, and progress logs. See `TASKS/template.md` for the format.

## Task Lifecycle

1. **Create**: Simple tasks go in `TODO.md`. Complex tasks get a file in `TASKS/`. Large efforts go in `PROJECTS.md`.
2. **Work**: Update status to `in-progress`. Add progress notes for complex tasks.
3. **Complete**: Mark `[x] **done**`. Move to completed section.
4. **Blocked**: Mark `**blocked**` with reason.

## Best Practices

- Keep `TODO.md` current -- update status regularly, archive completed tasks
- Use `TASKS/` files for anything with >3 steps or spanning multiple sessions
- Link related items across `TODO.md`, `PROJECTS.md`, and `TASKS/`
- Review weekly to keep priorities accurate
