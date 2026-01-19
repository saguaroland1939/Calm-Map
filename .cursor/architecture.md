# Calm-Map Architecture Documentation

## Overview

Calm-Map is a client-side only GIS web application. All functionality runs in the browser with no backend server required. The application uses Leaflet.js for map rendering and interaction.

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────┐
│                    Browser Window                        │
│                                                           │
│  ┌──────────────┐              ┌──────────────────┐    │
│  │   Sidebar    │              │                  │    │
│  │              │              │       Map        │    │
│  │ - Filters    │              │  (Leaflet.js)    │    │
│  │ - Search     │              │                  │    │
│  │ - Results    │              │  - Markers       │    │
│  │   List       │              │  - Popups        │    │
│  └──────────────┘              └──────────────────┘    │
│                                                           │
│  ┌──────────────────────────────────────────────────┐   │
│  │         JavaScript (embedded in index.html)      │   │
│  │  - Data loading (fetch points.json)              │   │
│  │  - Marker creation                                │   │
│  │  - Filter logic                                   │   │
│  │  - Event handlers                                 │   │
│  └──────────────────────────────────────────────────┘   │
│                                                           │
│  ┌──────────────────────────────────────────────────┐   │
│  │              External Resources                    │   │
│  │  - points.json (location data)                    │   │
│  │  - Leaflet.js CDN                                 │   │
│  │  - Esri Vector Tiles (basemap)                    │   │
│  │  - Image assets (GitHub raw URLs)                 │   │
│  └──────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
```

## Component Breakdown

### 1. HTML Structure (`index.html`)

The application uses a single HTML file with embedded JavaScript:

- **Sidebar (`#sidebar`)**: Fixed right-side panel containing:
  - Logo and branding
  - Search input
  - Category dropdown
  - Cost slider filter
  - Accessibility slider filter
  - "Show all" button
  - Results list (`#places-list`)

- **Map Container (`#map`)**: Full-screen Leaflet map instance

### 2. CSS Styling (`style.css`)

- **Layout**: Absolute positioning for sidebar and map
- **Map**: Takes `calc(100vw - 320px)` width, full height
- **Sidebar**: Fixed 320px width, right-aligned
- **Popups**: Custom styling for Leaflet popups with calm aesthetic
- **Responsive**: Currently desktop-focused, but structure supports mobile

### 3. JavaScript Logic (embedded in `index.html`)

#### Data Flow

1. **Initialization**:
   ```
   Page Load → Fetch points.json → Parse JSON → Create Markers → Add to Map
   ```

2. **Filtering Flow**:
   ```
   User Input → filterMarkers() → Remove All Markers → Filter Array → 
   Add Filtered Markers → Update Sidebar List
   ```

3. **Marker Creation**:
   ```
   Place Object → Create L.marker() → Attach Properties → 
   Create Popup HTML → Bind Popup → Add to Map
   ```

#### Key Data Structures

**`placesData`**: Array of place objects from `points.json`
```javascript
[
  {
    name: "Place Name",
    lat: 33.1234,
    lng: -111.5678,
    address: "Full address",
    "vibe-type": "Nature",
    cost: "Easy on the wallet",
    accessibility: "Easy",
    image: "https://..."
  }
]
```

**`markers`**: Array of Leaflet marker objects with attached properties
```javascript
[
  L.marker([lat, lng]) {
    name: "Place Name",
    category: "Nature",
    cost: "Easy on the wallet",
    accessibility: "Easy",
    address: "Full address"
  }
]
```

#### State Management

- **Global Variables**:
  - `markers`: Array of all marker instances
  - `placesData`: Raw data from JSON
  - `map`: Leaflet map instance

- **Filter State**: Managed through DOM elements:
  - `searchEl.value`: Search query
  - `categoryEl.value`: Selected category
  - `costEl.value`: Cost filter (0-2)
  - `accessEl.value`: Accessibility filter (0-2)

### 4. External Dependencies

#### Leaflet.js (v1.9.4)
- Core mapping library
- Handles map rendering, markers, popups, bounds
- CDN loaded: `https://unpkg.com/leaflet@1.9.4/dist/leaflet.js`

#### Esri Leaflet Plugins
- **esri-leaflet**: Base plugin for Esri services
- **esri-leaflet-vector**: Vector tile support
- Used for basemap: `World_Basemap_v2/VectorTileServer`

#### Data Source
- **points.json**: Static JSON file in project root
- Loaded via `fetch()` on page load
- Contains array of place objects

## Key Functions

### `filterMarkers()`
- Reads filter values from DOM
- Removes all markers from map
- Filters markers array based on criteria
- Re-adds filtered markers to map
- Updates sidebar list via `renderPlacesList()`

### `renderPlacesList(places)`
- Takes filtered places array
- Clears existing list
- Creates DOM elements for each place
- Adds click handlers to zoom map to location

### Marker Creation (inline)
- Maps over `placesData` array
- Creates Leaflet marker for each place
- Attaches metadata as properties
- Generates popup HTML template
- Binds popup and adds to map

## Event Handlers

- **Search input**: `input` event → `filterMarkers()`
- **Category select**: `change` event → `filterMarkers()`
- **Cost slider**: `input` event → update label + `filterMarkers()`
- **Accessibility slider**: `input` event → update label + `filterMarkers()`
- **Show all button**: `click` event → reset filters + show all + zoom to bounds
- **Place list items**: `onclick` → zoom map to place location

## Map Configuration

- **Initial View**: `[33.4484, -112.074]` (Phoenix, AZ), zoom level 11
- **Basemap**: Esri Vector Tile Service
- **Default Location**: `[41.024893, 28.958593]` (Istanbul, Turkey) - fallback
- **Location Services**: Currently commented out

## Data Relationships

```
points.json (source of truth)
    ↓
placesData (loaded into memory)
    ↓
markers (Leaflet objects with attached properties)
    ↓
Map Display + Sidebar List (filtered views)
```

## Performance Considerations

- All markers created upfront (no lazy loading)
- Filtering removes/adds markers from map (efficient for small datasets)
- No debouncing on filter inputs (acceptable for current scale)
- Images loaded on popup open (lazy loading)

## Future Architecture Considerations

- Consider marker clustering for large datasets
- Add debouncing to search input
- Implement URL parameters for shareable filtered views
- Consider splitting JavaScript into separate file for maintainability
- Add service worker for offline capability
