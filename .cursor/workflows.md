# Development Workflows

This document outlines common development workflows and procedures for working with the Calm-Map codebase.

## Workflow: Adding a New Place Location

### Steps

1. **Prepare the image**:
   - Add image file to `docs/` directory
   - Ensure image is optimized (reasonable file size)
   - Note the filename

2. **Get location coordinates**:
   - Use Google Maps or similar to find exact lat/lng
   - Verify coordinates are accurate

3. **Add to `points.json`**:
   ```json
   {
     "name": "Place Name",
     "lat": 33.1234,
     "lng": -111.5678,
     "address": "Full street address, City, State ZIP",
     "vibe-type": "Nature|Eats & Drinks|Exercise|Wellness|Arts|Quiet",
     "cost": "Easy on the wallet|Something reasonable|Big treat to self",
     "accessibility": "Easy|Moderate|Challenging",
     "image": "https://github.com/[repo]/blob/main/docs/filename.jpg?raw=true"
   }
   ```

4. **Verify**:
   - Open `index.html` in browser
   - Check marker appears on map
   - Click marker to verify popup displays correctly
   - Test that filters work with new place

### Notes
- Image URLs must use GitHub raw URL format: `?raw=true`
- Ensure all required fields are present
- Follow existing naming conventions for vibe-type, cost, accessibility

---

## Workflow: Adding a New Filter

### Steps

1. **Add UI element** in `index.html` sidebar:
   ```html
   <div class="new-filter">
     <label for="new-filter-input">Filter Label</label>
     <input type="..." id="new-filter-input" />
   </div>
   ```

2. **Add CSS styling** in `style.css`:
   ```css
   .new-filter {
     margin: 12px 0;
   }
   #new-filter-input {
     width: 100%;
     /* styling */
   }
   ```

3. **Get DOM element** in JavaScript:
   ```javascript
   const newFilterEl = document.getElementById("new-filter-input");
   ```

4. **Update `filterMarkers()` function**:
   - Read value from new filter element
   - Add filter condition to `filtered` array logic
   - Ensure it works with other filters (AND logic)

5. **Add event listener**:
   ```javascript
   newFilterEl?.addEventListener("input", filterMarkers);
   ```

6. **Update "Show all" button**:
   - Reset new filter to default value
   - Ensure it's included in reset logic

7. **Test**:
   - Verify filter works independently
   - Verify filter works with other filters
   - Verify "Show all" resets new filter

### Notes
- Keep filter logic consistent with existing patterns
- Update both marker filtering and sidebar list filtering
- Consider adding filter to `renderPlacesList()` if needed

---

## Workflow: Modifying Map Styling

### Steps

1. **Identify target element**:
   - Map container: `#map`
   - Popups: `.leaflet-popup-content-wrapper`, `.popup-*` classes
   - Sidebar: `#sidebar` and child elements

2. **Make CSS changes** in `style.css`:
   - Use browser DevTools to test changes live
   - Verify changes don't break layout

3. **Test across scenarios**:
   - Different screen sizes (if responsive)
   - With/without filters applied
   - With popups open/closed

4. **Maintain aesthetic**:
   - Keep calm, peaceful color scheme
   - Ensure readability
   - Maintain consistency with existing design

### Notes
- Leaflet adds its own classes - inspect in browser to see actual classes
- Popup styling requires careful CSS specificity
- Test with multiple markers to ensure styling scales

---

## Workflow: Debugging Filter Issues

### Steps

1. **Check filter values**:
   ```javascript
   console.log("Search:", searchEl?.value);
   console.log("Category:", categoryEl?.value);
   console.log("Cost:", costEl?.value);
   console.log("Access:", accessEl?.value);
   ```

2. **Check marker properties**:
   ```javascript
   markers.forEach(m => {
     console.log(m.name, m.category, m.cost, m.accessibility);
   });
   ```

3. **Check filter logic**:
   - Add console.logs in `filterMarkers()` to see which markers pass/fail
   - Verify string comparisons (case sensitivity, whitespace)

4. **Check data consistency**:
   - Verify `points.json` has expected values
   - Check for typos in vibe-type, cost, accessibility values

### Common Issues
- Case sensitivity: Use `.toLowerCase()` for comparisons
- Missing properties: Use fallbacks (`|| ""`)
- Slider values: Remember they're strings, need `parseInt()`

---

## Workflow: Updating Documentation

### When to Update

- **`.cursorrules`**: When adding project-wide conventions or changing tech stack
- **`.cursor/architecture.md`**: When changing code structure, adding components, or modifying data flow
- **`.cursor/patterns.md`**: When introducing new coding patterns or modifying existing ones
- **`.cursor/workflows.md`**: When adding new workflows or changing procedures
- **`README.md`**: When changing project description, setup, or usage

### Steps

1. **Make code changes first**
2. **Identify affected documentation**:
   - What changed? (structure, patterns, workflows, etc.)
   - Which docs reference this?
3. **Update documentation**:
   - Keep examples accurate
   - Update diagrams if needed
   - Add new patterns/workflows
4. **Verify documentation accuracy**:
   - Read through updated sections
   - Ensure examples match current code
   - Check for broken references

### Documentation Maintenance Checklist

- [ ] Code changes complete
- [ ] Relevant documentation files identified
- [ ] Examples updated to match current code
- [ ] Architecture diagrams updated (if applicable)
- [ ] New patterns/workflows documented
- [ ] README updated (if project-level changes)
- [ ] Documentation reviewed for accuracy

---

## Workflow: Testing Changes

### Pre-Change Checklist

- [ ] Understand what you're changing
- [ ] Check if documentation needs updating
- [ ] Identify potential side effects

### Testing Checklist

- [ ] **Basic functionality**:
  - [ ] Map loads correctly
  - [ ] All markers display
  - [ ] Popups show correct information

- [ ] **Filtering**:
  - [ ] Search filter works
  - [ ] Category filter works
  - [ ] Cost filter works
  - [ ] Accessibility filter works
  - [ ] Multiple filters work together
  - [ ] "Show all" resets everything

- [ ] **UI/UX**:
  - [ ] Sidebar displays correctly
  - [ ] Map displays correctly
  - [ ] Popups styled correctly
  - [ ] No console errors

- [ ] **Edge cases**:
  - [ ] Empty search results
  - [ ] All filters applied (very restrictive)
  - [ ] No filters applied (show all)

### Post-Change Checklist

- [ ] Code works as expected
- [ ] No regressions introduced
- [ ] Documentation updated (if needed)
- [ ] Changes tested in browser

---

## Workflow: Code Review Process

### For AI Agents

When reviewing code changes:

1. **Check consistency**:
   - Follows existing patterns?
   - Uses established conventions?
   - Matches code style?

2. **Check functionality**:
   - Logic is correct?
   - Edge cases handled?
   - No obvious bugs?

3. **Check documentation**:
   - Documentation updated?
   - Examples still accurate?
   - New patterns documented?

4. **Suggest improvements**:
   - Can code be simplified?
   - Are there better patterns?
   - Performance considerations?

### For Human Developers

- Review AI-suggested changes
- Verify tests pass
- Check documentation accuracy
- Ensure changes align with project goals

---

## Best Practices

1. **Always test in browser** - Don't assume code works without testing
2. **Update documentation** - Keep docs in sync with code
3. **Follow existing patterns** - Consistency helps maintainability
4. **Handle edge cases** - Empty data, missing properties, etc.
5. **Use meaningful names** - Variables and functions should be descriptive
6. **Comment complex logic** - Especially map-related operations
7. **Keep functions focused** - Single responsibility principle
8. **Test filters together** - Ensure they work in combination
