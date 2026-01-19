# Guide to Writing Effective Agent Instructions

This guide teaches you how to write markdown files that help AI coding agents work effectively with your codebase. Good agent instructions lead to better code suggestions, fewer errors, and more productive AI-assisted development.

## Why Agent Instructions Matter

AI coding agents need context to be effective. Without clear instructions, they:
- Make assumptions that might be wrong
- Suggest code that doesn't fit your project's style
- Miss important patterns or conventions
- Create inconsistencies

Well-written agent instructions help agents:
- Understand your project's architecture
- Follow your coding conventions
- Use established patterns
- Make informed decisions

## Types of Agent Instruction Files

### 1. `.cursorrules` - Project Rules

**Purpose**: High-level project rules and conventions that apply to all work.

**Location**: Root of your project

**What to Include**:
- Project overview and purpose
- Tech stack and versions
- Code style guidelines
- Naming conventions
- Common patterns
- Important notes and gotchas

**Example Structure**:
```markdown
# Project Rules

## Project Overview
[What the project does]

## Tech Stack
[Technologies used]

## Code Style Guidelines
[How code should be written]

## Key Patterns
[Common patterns used]

## When Making Changes
[Important reminders]
```

**Best Practices**:
- Keep it concise but complete
- Use clear section headers
- Include code examples for patterns
- Update when conventions change
- Reference detailed docs in `.cursor/` directory

### 2. `.cursor/` Directory - Detailed Documentation

**Purpose**: Organized, detailed documentation for specific aspects of your project.

**Structure**:
- `README.md` - Guide to the directory
- `architecture.md` - Technical architecture
- `patterns.md` - Coding patterns and examples
- `workflows.md` - Development workflows

**When to Use**:
- `.cursorrules` is getting too long
- You need detailed explanations
- You want to organize by topic
- You need diagrams or extensive examples

## Writing Effective Agent Instructions

### 1. Be Specific

**Bad**:
```markdown
Use good variable names.
```

**Good**:
```markdown
Use descriptive variable names:
- `placesData` for the array of places
- `filterMarkers` for the filtering function
- `marker` for individual marker instances
```

### 2. Provide Context

**Bad**:
```markdown
Markers are created from JSON data.
```

**Good**:
```markdown
Markers are created from `points.json` data:
1. Data is loaded via `fetch("points.json")`
2. Each place object becomes a Leaflet marker
3. Metadata is attached as properties (`marker.name`, `marker.category`)
4. Popups are bound with HTML templates
```

### 3. Show Examples

**Bad**:
```markdown
Use the filtering pattern.
```

**Good**:
```markdown
Filtering uses this pattern:
```javascript
const filtered = markers.filter(m => {
  const matchesText = !q || m.name.includes(q);
  const matchesCat = !cat || cat === "all" || m.category === cat;
  return matchesText && matchesCat;
});
```
```

### 4. Explain Why

**Bad**:
```markdown
Attach properties to markers.
```

**Good**:
```markdown
Attach properties to markers because Leaflet markers don't natively store custom data. This allows easy access during filtering:
```javascript
marker.name = place.name;
marker.category = place["vibe-type"];
```
```

### 5. Document Patterns, Not Just Rules

**Bad**:
```markdown
Don't create markers in filter functions.
```

**Good**:
```markdown
**Pattern**: Markers are created once from data, then filtered by adding/removing from map.

**Why**: Creating markers is expensive. Filtering by showing/hiding is more efficient.

**Example**:
```javascript
// Create all markers upfront
markers = data.map(place => createMarker(place));

// Filter by removing/adding to map
function filterMarkers() {
  markers.forEach(m => map.removeLayer(m));
  const filtered = markers.filter(/* criteria */);
  filtered.forEach(m => m.addTo(map));
}
```
```

## Documentation Maintenance

### The Documentation Maintenance Rule

**Critical**: Always update documentation when code changes.

This is both a rule (for agents to follow) and a skill (for you to develop).

### When to Update Documentation

1. **Code Structure Changes**:
   - New files or directories
   - Changed file organization
   - Updated architecture

2. **Pattern Changes**:
   - New coding patterns introduced
   - Existing patterns modified
   - New conventions established

3. **Workflow Changes**:
   - New development processes
   - Changed procedures
   - Updated tools or dependencies

4. **Feature Additions**:
   - New functionality added
   - New components created
   - New data structures introduced

### How to Maintain Documentation

1. **Make code changes first**
2. **Identify affected documentation**:
   - What changed?
   - Which docs reference this?
   - What examples need updating?
3. **Update documentation**:
   - Keep examples accurate
   - Update diagrams if needed
   - Add new patterns/workflows
4. **Verify accuracy**:
   - Read through updated sections
   - Ensure examples match current code
   - Check for broken references

### Documentation Checklist

After making code changes:
- [ ] Code changes complete
- [ ] Relevant documentation files identified
- [ ] Examples updated to match current code
- [ ] Architecture diagrams updated (if applicable)
- [ ] New patterns/workflows documented
- [ ] README updated (if project-level changes)
- [ ] Documentation reviewed for accuracy

## Common Patterns for Agent Instructions

### Pattern: Data Structure Documentation

```markdown
### Data Structure
Places in `points.json` have this structure:
```json
{
  "name": "Place Name",
  "lat": 33.1234,
  "lng": -111.5678,
  "vibe-type": "Nature",
  "cost": "Easy on the wallet",
  "accessibility": "Easy",
  "image": "https://..."
}
```
```

### Pattern: Function Documentation

```markdown
### `filterMarkers()`
**Purpose**: Filter markers based on user input

**How it works**:
1. Reads filter values from DOM
2. Removes all markers from map
3. Filters markers array based on criteria
4. Re-adds filtered markers to map
5. Updates sidebar list

**Usage**: Called on filter input events
```

### Pattern: Workflow Documentation

```markdown
## Workflow: Adding a New Place

### Steps
1. Add entry to `points.json`
2. Ensure image URL is accessible
3. Test marker appears correctly

### Notes
- Image URLs must use GitHub raw format
- All required fields must be present
```

## Tips for Effective Agent Instructions

1. **Start with the big picture**: Give context before details
2. **Use consistent formatting**: Headers, code blocks, lists
3. **Link related docs**: Cross-reference for deeper understanding
4. **Keep it current**: Outdated docs are worse than no docs
5. **Test with agents**: See if agents understand your instructions
6. **Iterate**: Improve based on agent performance

## Learning Exercise

Try writing agent instructions for a feature in your project:

1. **Choose a feature** (e.g., "adding a new filter")
2. **Write instructions** covering:
   - What the feature does
   - How it's implemented
   - What patterns it uses
   - How to modify it
3. **Test with an agent**: Ask the agent to implement a similar feature
4. **Refine**: Improve instructions based on results

## Resources

- `.cursorrules` - See project rules example
- `.cursor/` - See detailed documentation examples
- `.cursor/workflows.md` - See workflow documentation examples
- `.cursor/patterns.md` - See pattern documentation examples

## Remember

Good agent instructions are:
- **Accurate**: Reflect current codebase state
- **Complete**: Cover what agents need to know
- **Clear**: Easy to understand
- **Maintained**: Updated when code changes
- **Practical**: Include examples and patterns

The better your instructions, the better your AI-assisted development experience will be!
