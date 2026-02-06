# Guide to Writing Effective Agent Instructions

How to write markdown files that help AI coding agents work well with your codebase.

## Why It Matters

Without clear instructions, agents make wrong assumptions, suggest inconsistent code, and miss your conventions. Good instructions give agents the context they need to be effective.

## Types of Instruction Files

### `.cursorrules` (project root)
High-level rules: project overview, tech stack, code style, key patterns, important notes. This is the first file agents read -- keep it concise.

### `.cursor/` directory
Detailed docs organized by topic: `architecture.md`, `patterns.md`, `workflows.md`, `skills.md`. Use these when a topic needs more depth than `.cursorrules` can hold.

## Five Principles for Good Instructions

### 1. Be Specific

Bad: "Use good variable names."
Good: "Use `placesData` for the array, `filterMarkers` for the function, `marker` for instances."

### 2. Provide Context

Bad: "Markers are created from JSON data."
Good: "Data loads via `fetch('points.json')`, each object becomes a Leaflet marker with metadata attached as properties."

### 3. Show Examples

Include real code snippets from your codebase. Agents learn patterns best from concrete examples -- see `patterns.md` for how this project documents them.

### 4. Explain Why

Bad: "Attach properties to markers."
Good: "Attach properties to markers because Leaflet doesn't natively store custom data -- this enables filtering."

### 5. Document Patterns, Not Just Rules

Bad: "Don't create markers in filter functions."
Good: "**Pattern**: Create markers once, filter by adding/removing from map. **Why**: Creating markers is expensive."

## Keeping Docs Current

Always update documentation when code changes. See the Documentation Maintenance section in `.cursorrules` for specific guidance on which files to update.

## Tips

1. Start with the big picture, then details
2. Use consistent formatting (headers, code blocks, lists)
3. Cross-reference related docs instead of duplicating content
4. Test instructions by asking an agent to implement a similar feature
5. Iterate based on agent performance
