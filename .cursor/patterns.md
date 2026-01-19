# Coding Patterns and Examples

This document outlines common coding patterns used in the Calm-Map codebase. Use these as reference when adding new features or modifying existing code.

## Pattern: Marker Creation with Metadata

**Purpose**: Create Leaflet markers with attached properties for filtering

**Pattern**:
```javascript
const marker = L.marker([place.lat, place.lng]);
marker.name = place.name || "";
marker.address = place.address || "";
marker.category = place["vibe-type"] || place.category || "";
marker.cost = place.cost || "";
marker.accessibility = place.accessibility || "";
```

**Why**: Leaflet markers don't natively store custom data, so we attach properties directly to the marker object for easy access during filtering.

**Usage**: Used when creating markers from `points.json` data.

---

## Pattern: Popup HTML Template

**Purpose**: Generate HTML content for marker popups

**Pattern**:
```javascript
const popupHtml = `
  <div class="popup-content">
    <div class="popup-name">${place.name}</div>
    <img src="${place.image}" alt="${place.name}" class="popup-image">
    <div class="popup-category"><strong>Vibe type:</strong> ${place["vibe-type"] || 'N/A'}</div>
    <div class="popup-price"><strong>Price:</strong> ${place.cost || 'N/A'}</div>
    <div class="popup-accessibility"><strong>Accessibility:</strong> ${place.accessibility || 'N/A'}</div>
    <div class="popup-address"><strong>Address:</strong> ${place.address || 'N/A'}</div>
  </div>
`;
marker.bindPopup(popupHtml, { maxWidth: 320 });
```

**Why**: Template literals allow clean HTML generation with data interpolation. Fallback values (`|| 'N/A'`) handle missing data gracefully.

**Usage**: When creating markers, bind this popup HTML to each marker.

---

## Pattern: Multi-Criteria Filtering

**Purpose**: Filter markers based on multiple criteria simultaneously

**Pattern**:
```javascript
function filterMarkers() {
  const q = (searchEl?.value || "").trim().toLowerCase();
  const cat = categoryEl?.value || "";
  const costIdx = parseInt(costEl?.value ?? "0", 10);
  const costVal = COST_VALUES[costIdx];
  const accessIdx = parseInt(accessEl?.value ?? "0", 10);
  const accessVal = ACCESS_VALUES[accessIdx];

  markers.forEach(m => map.removeLayer(m));

  const filtered = markers.filter(m => {
    const matchesText = !q || 
      (m.name || "").toLowerCase().includes(q) || 
      (m.address || "").toLowerCase().includes(q);
    const matchesCat = !cat || cat === "all" || m.category === cat;
    const matchesCost = (m.cost || "").toLowerCase() === costVal.toLowerCase();
    const matchesAccess = (m.accessibility || "").toLowerCase() === accessVal.toLowerCase();
    return matchesText && matchesCat && matchesCost && matchesAccess;
  });

  filtered.forEach(m => m.addTo(map));
}
```

**Why**: 
- Uses logical AND (`&&`) to combine all filter criteria
- Each filter checks if it's "empty" (no filter applied) or matches
- Removes all markers first, then re-adds filtered ones (cleaner than tracking state)

**Usage**: Called whenever any filter input changes.

---

## Pattern: Slider Label Updates

**Purpose**: Update text label to match slider value in real-time

**Pattern**:
```javascript
const COST_VALUES = ["Easy on the wallet", "Something reasonable", "Big treat to self"];

costEl?.addEventListener("input", () => {
  costLabel.textContent = COST_VALUES[parseInt(costEl.value, 10)];
});
```

**Why**: Sliders use numeric values (0-2) but display human-readable labels. This pattern maps the numeric value to the corresponding label.

**Usage**: Applied to both cost and accessibility sliders.

---

## Pattern: Reset All Filters

**Purpose**: Clear all filters and show all markers

**Pattern**:
```javascript
showAllBtn?.addEventListener("click", () => {
  if (searchEl) searchEl.value = "";
  if (categoryEl) categoryEl.value = "";
  if (costEl) { 
    costEl.value = "0"; 
    costLabel.textContent = COST_VALUES[0]; 
  }
  if (accessEl) { 
    accessEl.value = "0"; 
    accessLabel.textContent = ACCESS_VALUES[0]; 
  }

  markers.forEach(m => map.removeLayer(m));
  markers.forEach(m => m.addTo(map));
  renderPlacesList(placesData);

  const allLatLngs = markers.map(m => m.getLatLng());
  if (allLatLngs.length) {
    const bounds = L.latLngBounds(allLatLngs);
    map.fitBounds(bounds, { padding: [50, 50] });
  }
});
```

**Why**: 
- Resets each filter to its default state
- Re-adds all markers
- Zooms map to show all locations using `fitBounds()`

**Usage**: "Show all" button click handler.

---

## Pattern: Dynamic List Rendering

**Purpose**: Render filtered list of places in sidebar

**Pattern**:
```javascript
function renderPlacesList(places) {
  const list = document.getElementById("places-list");
  list.innerHTML = "";

  places.forEach(place => {
    const item = document.createElement("div");
    item.className = "place-item";
    const cat = place["vibe-type"] || place.category || "";
    item.innerHTML = `<b>${place.name}</b><br><span>${cat}</span>`;
    item.onclick = () => map.setView([place.lat, place.lng], 18);
    list.appendChild(item);
  });
}
```

**Why**: 
- Clears existing list first (`innerHTML = ""`)
- Creates DOM elements programmatically
- Attaches click handler to zoom map to location

**Usage**: Called after filtering to update sidebar list.

---

## Pattern: Safe Property Access

**Purpose**: Access nested properties safely with fallbacks

**Pattern**:
```javascript
// Using optional chaining and nullish coalescing
const value = element?.value ?? "default";

// Using logical OR for fallbacks
const category = place["vibe-type"] || place.category || "";

// Using optional chaining with method calls
costEl?.addEventListener("input", handler);
```

**Why**: Prevents errors when DOM elements or data properties might be undefined. Multiple fallback options handle different data structures.

**Usage**: Throughout the codebase when accessing DOM elements or data properties.

---

## Pattern: Map Bounds Calculation

**Purpose**: Zoom map to show all markers

**Pattern**:
```javascript
const allLatLngs = markers.map(m => m.getLatLng());
if (allLatLngs.length) {
  const bounds = L.latLngBounds(allLatLngs);
  map.fitBounds(bounds, { padding: [50, 50] });
}
```

**Why**: 
- Extracts lat/lng from all markers
- Creates Leaflet bounds object
- Fits map view to bounds with padding

**Usage**: "Show all" button and potentially for initial map setup.

---

## Pattern: Data Loading with Fetch

**Purpose**: Load JSON data asynchronously

**Pattern**:
```javascript
fetch("points.json")
  .then(r => r.json())
  .then(data => {
    placesData = data;
    // Process data...
  });
```

**Why**: Standard async data loading pattern. Could be converted to async/await if preferred.

**Usage**: Initial data load on page initialization.

---

## Common Anti-Patterns to Avoid

1. **Don't** create markers inside filter functions (they already exist)
2. **Don't** use `var` (use `const`/`let`)
3. **Don't** forget to handle missing data (use fallbacks)
4. **Don't** hardcode filter values (use constants like `COST_VALUES`)
5. **Don't** modify `placesData` directly (it's the source of truth)

## When Adding New Patterns

If you introduce a new pattern that might be reused:
1. Document it here with purpose, code example, and why
2. Update `.cursorrules` if it becomes a project convention
3. Consider creating a reusable function if it's used multiple times
