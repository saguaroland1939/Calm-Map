# Development Workflows

Step-by-step procedures for common Calm-Map development tasks.

## Adding a New Place

1. **Prepare image**: Add to `docs/`, optimize file size
2. **Get coordinates**: Use Google Maps to find exact lat/lng
3. **Add to `points.json`**: Follow the schema in `AGENT_RULES.md` "Data Structure"
   - Image URLs must use GitHub raw format (`?raw=true`)
   - All fields are required
4. **Verify in browser**: Check marker, popup, and filters work with new place

---

## Adding a New Filter

1. **Add UI element** in `index.html` sidebar:
   ```html
   <div class="new-filter">
     <label for="new-filter-input">Filter Label</label>
     <input type="..." id="new-filter-input" />
   </div>
   ```

2. **Add CSS** in `style.css` (match existing filter styling)

3. **Wire up in JavaScript**:
   - Get DOM element: `const newFilterEl = document.getElementById("new-filter-input");`
   - Add condition to `filterMarkers()` using AND logic
   - Add event listener: `newFilterEl?.addEventListener("input", filterMarkers);`
   - Add reset logic to "Show all" button handler

4. **Test**: Filter works alone, with other filters, and resets correctly

---

## Modifying Map Styling

1. **Identify target**: `#map`, `.leaflet-popup-content-wrapper`, `.popup-*`, `#sidebar`
2. **Edit `style.css`**: Use browser DevTools to test live
3. **Test**: Different screen sizes, with/without filters, popups open/closed
4. **Maintain aesthetic**: Keep calm, peaceful color scheme and readability

> Leaflet adds its own classes -- inspect in browser to see actual class names. Popup styling requires careful CSS specificity.

---

## Debugging Filter Issues

1. **Log filter values**:
   ```javascript
   console.log("Search:", searchEl?.value);
   console.log("Category:", categoryEl?.value);
   console.log("Cost:", costEl?.value);
   console.log("Access:", accessEl?.value);
   ```

2. **Log marker properties**:
   ```javascript
   markers.forEach(m => console.log(m.name, m.category, m.cost, m.accessibility));
   ```

3. **Check data**: Verify `points.json` values match expected strings (case, whitespace, typos)

**Common issues**: Case sensitivity (use `.toLowerCase()`), missing properties (use fallbacks), slider values are strings (need `parseInt()`)

---

## Testing Checklist

After making changes, verify:

- [ ] Map loads, all markers display, popups show correct info
- [ ] All filters work individually and in combination
- [ ] "Show all" resets everything
- [ ] Sidebar displays correctly, no console errors
- [ ] Edge cases: empty results, all filters applied, no filters applied
- [ ] Documentation updated if needed
