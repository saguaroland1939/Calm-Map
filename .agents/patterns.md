# Coding Patterns

Common coding patterns used in Calm-Map. Reference these when adding features or modifying code.

## Marker Creation with Metadata

**Purpose**: Attach filterable properties to Leaflet markers (which don't natively store custom data).

```javascript
const marker = L.marker([place.lat, place.lng]);
marker.name = place.name || "";
marker.address = place.address || "";
marker.category = place["vibe-type"] || place.category || "";
marker.cost = place.cost || "";
marker.accessibility = place.accessibility || "";
```

---

## Popup HTML Template

**Purpose**: Generate popup content using template literals with fallback values.

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

---

## Multi-Criteria Filtering

**Purpose**: Filter markers using AND logic across all criteria. Removes all markers first, then re-adds matches.

```javascript
function filterMarkers() {
  const q = (searchEl?.value || "").trim().toLowerCase();
  const cat = categoryEl?.value || "";
  const costVal = COST_VALUES[parseInt(costEl?.value ?? "0", 10)];
  const accessVal = ACCESS_VALUES[parseInt(accessEl?.value ?? "0", 10)];

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

---

## Slider Label Updates

**Purpose**: Map numeric slider values (0-2) to human-readable labels in real-time.

```javascript
const COST_VALUES = ["Easy on the wallet", "Something reasonable", "Big treat to self"];

costEl?.addEventListener("input", () => {
  costLabel.textContent = COST_VALUES[parseInt(costEl.value, 10)];
});
```

---

## Reset All Filters

**Purpose**: Clear filters, show all markers, zoom to full extent.

```javascript
showAllBtn?.addEventListener("click", () => {
  if (searchEl) searchEl.value = "";
  if (categoryEl) categoryEl.value = "";
  if (costEl) { costEl.value = "0"; costLabel.textContent = COST_VALUES[0]; }
  if (accessEl) { accessEl.value = "0"; accessLabel.textContent = ACCESS_VALUES[0]; }

  markers.forEach(m => map.removeLayer(m));
  markers.forEach(m => m.addTo(map));
  renderPlacesList(placesData);

  const allLatLngs = markers.map(m => m.getLatLng());
  if (allLatLngs.length) {
    map.fitBounds(L.latLngBounds(allLatLngs), { padding: [50, 50] });
  }
});
```

---

## Dynamic List Rendering

**Purpose**: Rebuild sidebar list from filtered data with click-to-zoom.

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

---

## Safe Property Access

**Purpose**: Prevent errors from undefined DOM elements or data properties.

```javascript
const value = element?.value ?? "default";          // nullish coalescing
const category = place["vibe-type"] || place.category || "";  // OR fallback
costEl?.addEventListener("input", handler);          // optional chaining
```

---

## Anti-Patterns to Avoid

1. Don't create markers inside filter functions (they already exist)
2. Don't use `var` (use `const`/`let`)
3. Don't forget fallbacks for missing data
4. Don't hardcode filter values (use constants like `COST_VALUES`)
5. Don't modify `placesData` directly (it's the source of truth)

## Adding New Patterns

Document new reusable patterns here with: purpose, code example, and why.
