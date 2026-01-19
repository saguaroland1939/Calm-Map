# Guide to Markdown-Based Task Management

This guide teaches you how to use markdown files for effective task management in your development workflow. Markdown-based task management is simple, version-controlled, and works great with AI coding agents.

## Why Markdown for Task Management?

### Advantages

1. **Version Control**: Tasks are tracked in git alongside your code
2. **Simple**: No special tools needed, just text files
3. **Flexible**: Structure tasks however works for you
4. **AI-Friendly**: Agents can read and update markdown easily
5. **Portable**: Works anywhere markdown is supported
6. **Searchable**: Easy to search across all tasks

### When to Use

- Small to medium projects
- Solo or small team development
- When you want tasks in your repository
- When working with AI coding agents
- When you prefer simple tools

### When Not to Use

- Large teams (consider GitHub Issues or project management tools)
- Complex project management needs
- Need for advanced features (time tracking, dependencies, etc.)

## Task Management Structure

### Three Levels of Organization

1. **TODO.md** - Quick, active task list
2. **PROJECTS.md** - Long-term project planning
3. **TASKS/** - Detailed task documentation

### TODO.md - Active Task List

**Purpose**: Quick reference for what needs to be done now.

**Structure**:
- Priority levels (P0, P1, P2, P3)
- Status tracking (todo, in-progress, done, blocked)
- Simple checkboxes
- Brief descriptions

**Best For**:
- Daily task tracking
- Quick status updates
- Prioritizing work
- Seeing what's next

**Example**:
```markdown
## P1 - High Priority

- [ ] **in-progress** - Add mobile responsive design
  - Make sidebar collapsible on mobile
  - Adjust map width calculation

- [ ] **todo** - Add marker clustering
  - Research Leaflet plugins
  - Implement clustering
```

### PROJECTS.md - Long-term Planning

**Purpose**: Track larger projects broken into milestones.

**Structure**:
- Project status (planning, active, on-hold, completed)
- Milestones with checkboxes
- Related tasks references
- Project descriptions

**Best For**:
- Multi-week/month projects
- Feature planning
- Breaking down large work
- Tracking progress over time

**Example**:
```markdown
### Mobile Responsiveness
**Status**: active  
**Priority**: P2

**Milestones**:
- [ ] Research mobile UX patterns
- [ ] Design mobile sidebar
- [ ] Implement responsive CSS
- [ ] Test on devices

**Related Tasks**: TODO.md - "Add mobile responsive design"
```

### TASKS/ Directory - Detailed Task Docs

**Purpose**: Detailed documentation for complex tasks.

**Structure**:
- Full task description
- Acceptance criteria
- Implementation plan
- Technical details
- Testing requirements
- Progress log

**Best For**:
- Complex tasks needing detailed planning
- Tasks with multiple steps
- Tasks requiring documentation
- Tasks that span multiple sessions

**Example**: See `TASKS/template.md` and `TASKS/example.md`

## Task Status System

### Status Values

- **todo**: Not started yet
- **in-progress**: Currently working on
- **blocked**: Waiting on something/someone
- **review**: Ready for review
- **done**: Completed

### Priority Levels

- **P0**: Critical - blocks other work or urgent
- **P1**: High - important feature or fix
- **P2**: Medium - nice to have, can wait
- **P3**: Low - future consideration

### Using Status

```markdown
- [ ] **todo** - Task not started
- [ ] **in-progress** - Currently working
- [ ] **blocked** - Waiting on dependency
- [x] **done** - Completed
```

## Task Lifecycle

### 1. Creating a Task

**Simple Task** → Add to `TODO.md`:
```markdown
- [ ] **todo** - Add debouncing to search input
```

**Complex Task** → Create file in `TASKS/`:
```markdown
# Task: Add Mobile Responsiveness
**Status**: todo
**Priority**: P1
...
```

**Project** → Add to `PROJECTS.md`:
```markdown
### Mobile Responsiveness
**Status**: planning
**Milestones**: ...
```

### 2. Working on a Task

Update status:
```markdown
- [ ] **in-progress** - Add debouncing to search input
```

For detailed tasks, add progress log:
```markdown
## Progress Log

### 2025-01-15 - Started
- Added debounce function
- Next: Wire up to search input
```

### 3. Completing a Task

Mark as done:
```markdown
- [x] **done** - Add debouncing to search input
```

Move to completed section:
```markdown
## Completed Tasks

- [x] **done** - Add debouncing to search input (2025-01-15)
```

### 4. Blocked Tasks

Mark as blocked with reason:
```markdown
- [ ] **blocked** - Add feature X (waiting on API access)
```

## Best Practices

### 1. Keep TODO.md Current

- Update status regularly
- Move completed tasks to completed section
- Remove obsolete tasks
- Keep priorities accurate

### 2. Use Detailed Tasks for Complexity

- If a task has >3 steps, consider `TASKS/` file
- If a task needs documentation, use `TASKS/` file
- If a task spans multiple sessions, use `TASKS/` file

### 3. Link Related Items

```markdown
**Related Tasks**: TODO.md - "Add mobile design"
**Related Projects**: PROJECTS.md - "Mobile Responsiveness"
**Related Docs**: .cursor/workflows.md
```

### 4. Update as You Work

- Change status when starting work
- Add progress notes for complex tasks
- Update when blocked or unblocked
- Mark done when complete

### 5. Regular Review

- Weekly: Review priorities and status
- Monthly: Review projects and milestones
- Quarterly: Archive completed projects

## Working with AI Agents

### How Agents Use Task Files

Agents can:
- Read task files to understand what needs to be done
- Update task status as they work
- Create new tasks when needed
- Reference tasks in their work

### Best Practices for Agent Tasks

1. **Be Specific**: Clear acceptance criteria help agents understand goals
2. **Provide Context**: Link to related code, docs, or issues
3. **Break Down**: Complex tasks into smaller, actionable steps
4. **Update Status**: Agents should update status as they work
5. **Document Progress**: For complex tasks, log what's been done

### Example Agent-Friendly Task

```markdown
# Task: Add Accessibility Filter

**Status**: todo
**Priority**: P1

## Description
Add accessibility filter slider to sidebar.

## Acceptance Criteria
- [ ] Slider appears in sidebar
- [ ] Three levels: Easy, Moderate, Challenging
- [ ] Filter works with other filters
- [ ] "Show all" resets filter

## Implementation Plan
1. Add HTML structure (see existing cost filter)
2. Add CSS styling (match cost filter)
3. Add JavaScript logic (follow filterMarkers pattern)

## Related
- Pattern: `.cursor/patterns.md` - Multi-criteria filtering
- Workflow: `.cursor/workflows.md` - Adding a new filter
```

## Task Templates

### Simple Task Template

```markdown
- [ ] **todo** - [Task description]
  - [Brief detail or sub-task]
  - [Another detail if needed]
```

### Detailed Task Template

See `TASKS/template.md` for full template.

### Project Template

```markdown
### [Project Name]
**Status**: planning|active|on-hold|completed
**Priority**: P0|P1|P2|P3
**Description**: [What this project accomplishes]

**Milestones**:
- [ ] Milestone 1
- [ ] Milestone 2

**Related Tasks**: [Links to TODO.md or TASKS/]
```

## Tips for Effective Task Management

1. **Start Simple**: Use TODO.md for most tasks
2. **Upgrade When Needed**: Move to TASKS/ for complex tasks
3. **Keep It Current**: Update status regularly
4. **Be Specific**: Clear tasks are easier to complete
5. **Link Everything**: Connect tasks, projects, and docs
6. **Review Regularly**: Keep priorities and status accurate
7. **Archive Completed**: Move done tasks to completed section

## Common Patterns

### Pattern: Breaking Down Large Tasks

```markdown
# Task: Add Mobile Support

## Implementation Plan
### Step 1: Research
- [ ] Research mobile UX patterns
- [ ] Review similar applications

### Step 2: Design
- [ ] Design mobile sidebar
- [ ] Create mockups

### Step 3: Implement
- [ ] Add responsive CSS
- [ ] Test on devices
```

### Pattern: Task Dependencies

```markdown
## Dependencies
- Requires: API access (blocked)
- Blocks: Feature X (can't start until this is done)
```

### Pattern: Time Tracking

```markdown
**Estimated Time**: 4 hours
**Actual Time**: 3.5 hours
```

## Learning Exercise

Try setting up task management for a feature:

1. **Create a simple task** in TODO.md
2. **Work on it** and update status
3. **If it gets complex**, create detailed task in TASKS/
4. **If it's part of larger work**, add to PROJECTS.md
5. **Complete it** and move to completed section

## Resources

- `TODO.md` - See active task list example
- `PROJECTS.md` - See project planning example
- `TASKS/template.md` - See detailed task template
- `TASKS/example.md` - See completed task example

## Remember

Effective task management:
- **Keeps you organized**: Know what needs to be done
- **Tracks progress**: See what's done and what's next
- **Helps planning**: Break down large work
- **Works with agents**: Agents can read and update tasks
- **Stays current**: Regular updates keep it useful

The simpler you keep it, the more likely you'll use it!
