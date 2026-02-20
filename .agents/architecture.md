# Calm-Map Architecture

## Overview

Client-side only GIS web app. All functionality runs in the browser -- no backend required. Uses Leaflet.js for map rendering.

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────┐
│                    Browser Window                        │
│                                                         │
│  ┌──────────────┐              ┌──────────────────┐    │
│  │   Sidebar    │              │       Map        │    │
│  │ - Filters    │              │  (Leaflet.js)    │    │
│  │ - Search     │              │  - Markers       │    │
│  │ - Results    │              │  - Popups        │    │
│  └──────────────┘              └──────────────────┘    │
│                                                         │
│  ┌──────────────────────────────────────────────────┐  │
│  │         JavaScript (embedded in index.html)      │  │
│  │  - Data loading (fetch points.json)              │  │
│  │  - Marker creation + Filter logic                │  │
│  │  - Event handlers                                │  │
│  └──────────────────────────────────────────────────┘  │
│                                                         │
│  ┌──────────────────────────────────────────────────┐  │
│  │  External: points.json, Leaflet CDN, Esri Tiles  │  │
│  └──────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────┘
```

## Component Breakdown

### HTML Structure (`index.html`)

Single HTML file with embedded JavaScript:

- **Sidebar (`#sidebar`)**: Fixed right-side panel -- logo, search input, category dropdown, cost slider, accessibility slider, "Show all" button, results list (`#places-list`)
- **Map Container (`#map`)**: Full-screen Leaflet map instance

### CSS Styling (`style.css`)

- **Layout**: Absolute positioning; map `calc(100vw - 320px)` wide, sidebar 320px right-aligned
- **Popups**: Custom styling for Leaflet popups with calm aesthetic
- **Responsive**: Currently desktop-focused

### JavaScript Logic

#### Data Flow

```
Page Load → fetch points.json → Create Markers → Add to Map
User Input → filterMarkers() → Remove All → Filter → Re-add Filtered → Update Sidebar
```

#### State

- `placesData`: Raw array from JSON (source of truth)
- `markers`: Array of Leaflet marker objects with attached metadata
- `map`: Leaflet map instance
- Filter state lives in DOM elements (`searchEl`, `categoryEl`, `costEl`, `accessEl`)

> For data structure schema, see `AGENT_RULES.md` "Data Structure" section.
> For code patterns, see `patterns.md`.

### Key Functions

- **`filterMarkers()`**: Reads DOM filter values, removes all markers, re-adds matching ones, updates sidebar
- **`renderPlacesList(places)`**: Clears and rebuilds sidebar list with click-to-zoom handlers

### Event Handlers

- **Search/Category/Cost/Accessibility inputs** → `filterMarkers()`
- **Show all button** → reset filters + show all + zoom to bounds
- **Place list items** → zoom map to location

## External Dependencies

- **Leaflet.js v1.9.4**: Map rendering, markers, popups, bounds (CDN)
- **esri-leaflet + esri-leaflet-vector**: Vector tile basemap (`World_Basemap_v2`)
- **points.json**: Static data file loaded via `fetch()`

## Map Configuration

- **Initial View**: Phoenix, AZ `[33.4484, -112.074]`, zoom 11
- **Fallback Location**: Istanbul `[41.024893, 28.958593]`
- **Location Services**: Currently commented out

## Data Relationships

```
points.json (source of truth)
    ↓
placesData (in memory)
    ↓
markers (Leaflet objects with attached properties)
    ↓
Map Display + Sidebar List (filtered views)
```

## Performance Notes

- All markers created upfront (no lazy loading)
- Filtering removes/adds markers (efficient for small datasets)
- No debouncing on filter inputs (acceptable at current scale)
- Images loaded on popup open
