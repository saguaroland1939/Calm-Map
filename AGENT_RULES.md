# Calm-Map Project Rules for AI Agents

## Conversation Start Protocol

**REQUIRED**: At the start of every new conversation, agents MUST:

1. **Start with an emoji** (e.g., 🗺️, 📍, 🌿)
2. **Acknowledge documentation**: List the files you've reviewed (`AGENT_RULES.md`, `architecture.md`, `patterns.md`, `workflows.md`, `skills.md`)
3. **Confirm understanding**: Briefly state the project context

**Example**: 🗺️ I've reviewed the Calm-Map docs and understand this is a vanilla JS GIS app using Leaflet.js. How can I help?

## Documentation Structure

Before starting ANY task, agents MUST read:
- `AGENT_RULES.md` - This file (project overview, rules, patterns)
- `.agents/START_HERE.md` - Agent protocol and conversation guidelines
- `.agents/architecture.md` - Technical architecture and component breakdown
- `.agents/patterns.md` - Code patterns with examples
- `.agents/workflows.md` - Step-by-step development procedures

These documents are kept current and contain critical context. Ignoring them leads to inconsistent code and poor results.

## Project Overview

Calm-Map is a GIS web app helping users discover calm, peaceful places in their community. Single-page application using vanilla HTML/CSS/JS with Leaflet.js.

## Tech Stack

- **Frontend**: Vanilla HTML5, CSS3, JavaScript (ES6+)
- **Mapping**: Leaflet.js v1.9.4 + Esri Vector Tile basemap
- **Data**: `points.json` (static JSON)
- **No build tools**: Static site served directly

## Project Structure

```
Calm-Map/
├── index.html       # HTML + embedded JavaScript
├── style.css        # All styling
├── points.json      # Location data
├── docs/            # Images and documentation
└── README.md
```

## Code Style

- **JS**: ES6+ (const/let, arrows, template literals), meaningful names, single-purpose functions
- **CSS**: Semantic class names, BEM-like naming, organized by component
- **HTML**: Semantic elements, accessibility (alt text, labels)

## Data Structure

Places in `points.json`:
```json
{
  "name": "Place Name",
  "lat": 33.1234,
  "lng": -111.5678,
  "address": "Full address string",
  "vibe-type": "Nature|Eats & Drinks|Exercise|Wellness|Arts|Quiet",
  "cost": "Easy on the wallet|Something reasonable|Big treat to self",
  "accessibility": "Easy|Moderate|Challenging",
  "image": "URL to image"
}
```

## Key Patterns

- Markers store metadata as properties for filtering (see `patterns.md`)
- Multiple filters use AND logic; both map and sidebar update together
- "Show all" resets filters and zooms to bounds

## Important Notes

- Map centered on Phoenix, AZ; fallback location is Istanbul, Turkey
- Location services currently commented out
- Images hosted on GitHub (raw URLs)
- No external dependencies beyond CDN links

## Common Tasks

See `.agents/workflows.md` for detailed step-by-step procedures.

## Documentation Maintenance

**CRITICAL**: When making code changes, update the relevant docs:

- Code structure changes → `architecture.md` + project structure above
- New patterns → `patterns.md`
- Workflow changes → `workflows.md`
- Feature additions → `README.md` + `AGENT_RULES.md`
- Data structure changes → `AGENT_RULES.md` + `architecture.md`

## When Making Changes

- Test in browser after changes
- Verify markers, filters, and popups work correctly
- Maintain the calm, peaceful aesthetic
- Update relevant documentation
