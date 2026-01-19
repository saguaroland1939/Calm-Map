# Agent Skills

This document defines reusable skills that AI agents should apply when working on this codebase. Skills are techniques, approaches, and best practices that agents can use to solve problems effectively.

## Skill: Test-Driven Development (TDD)

### Overview

Test-Driven Development (TDD) is a development approach where you write tests before writing the code that makes them pass. This skill guides agents on how to apply TDD principles in the context of this vanilla JavaScript, client-side application.

### TDD Cycle: Red-Green-Refactor

The core TDD cycle consists of three phases:

1. **Red**: Write a failing test that describes the desired behavior
2. **Green**: Write the minimum code needed to make the test pass
3. **Refactor**: Improve the code while keeping tests passing

### Applying TDD to Calm-Map

#### Context Considerations

This project uses:
- Vanilla JavaScript (no build tools)
- DOM manipulation
- Leaflet.js for mapping
- Fetch API for data loading
- No existing test framework

#### Testing Strategy Options

**Option 1: Manual Testing Approach**
- Write test descriptions as comments
- Manually verify behavior in browser
- Document expected vs actual results
- Good for: Quick iterations, learning TDD concepts

**Option 2: Simple Test Framework**
- Use a lightweight framework like QUnit or Mocha (CDN)
- Write actual test code
- Run tests in browser
- Good for: More rigorous testing, regression prevention

**Option 3: Test-Driven Thinking**
- Write test descriptions before code
- Think through edge cases upfront
- Verify manually but systematically
- Good for: Building TDD habits without setup overhead

### TDD Workflow for This Project

#### Step 1: Understand the Requirement

Before writing tests, understand:
- What feature/functionality is needed?
- What are the inputs and expected outputs?
- What edge cases exist?
- How does it integrate with existing code?

#### Step 2: Write Test Description (Red Phase)

Document what you're testing:

```javascript
// Test: filterMarkers() should filter by search query
// Given: markers array with ["Coffee Shop", "Park", "Museum"]
// When: search query is "Coffee"
// Then: only "Coffee Shop" marker should be visible

// Test: filterMarkers() should handle empty search
// Given: markers array with multiple places
// When: search query is empty string
// Then: all markers should be visible

// Test: filterMarkers() should be case-insensitive
// Given: marker with name "Coffee Shop"
// When: search query is "coffee"
// Then: marker should be visible
```

#### Step 3: Implement Minimum Code (Green Phase)

Write the simplest code that makes the test pass:

```javascript
function filterMarkers() {
  const q = (searchEl?.value || "").trim().toLowerCase();
  
  markers.forEach(m => map.removeLayer(m));
  
  const filtered = markers.filter(m => {
    const name = (m.name || "").toLowerCase();
    return !q || name.includes(q);
  });
  
  filtered.forEach(m => m.addTo(map));
}
```

#### Step 4: Refactor (Refactor Phase)

Improve code while keeping behavior the same:

```javascript
function filterMarkers() {
  const searchQuery = getSearchQuery();
  const categoryFilter = getCategoryFilter();
  // ... other filters
  
  const filtered = applyFilters(markers, {
    search: searchQuery,
    category: categoryFilter,
    // ... other filters
  });
  
  updateMapMarkers(filtered);
  updateSidebarList(filtered);
}

// Extract helper functions for clarity
function getSearchQuery() {
  return (searchEl?.value || "").trim().toLowerCase();
}
```

### TDD Patterns for This Codebase

#### Pattern: Testing DOM Manipulation

**Challenge**: Testing DOM code requires a browser environment.

**Approach**:
1. Write test descriptions for DOM behavior
2. Test in browser console or DevTools
3. Verify DOM state matches expectations
4. Document test cases for future reference

**Example**:
```javascript
// Test: renderPlacesList() should create list items
// Given: places array with 3 places
// When: renderPlacesList(places) is called
// Then: #places-list should contain 3 .place-item elements
// Then: Each item should have onclick handler
// Then: Clicking item should zoom map to place location

function renderPlacesList(places) {
  const list = document.getElementById("places-list");
  list.innerHTML = "";
  
  places.forEach(place => {
    const item = document.createElement("div");
    item.className = "place-item";
    item.innerHTML = `<b>${place.name}</b>`;
    item.onclick = () => map.setView([place.lat, place.lng], 18);
    list.appendChild(item);
  });
}

// Manual verification:
// console.log(document.querySelectorAll('.place-item').length) // Should be 3
```

#### Pattern: Testing Filter Logic

**Challenge**: Filters involve multiple criteria and state.

**Approach**:
1. Test each filter independently
2. Test filter combinations
3. Test edge cases (empty, null, undefined)
4. Test with sample data

**Example**:
```javascript
// Test: filterByCategory() with sample data
const testMarkers = [
  { name: "Park", category: "Nature" },
  { name: "Cafe", category: "Eats & Drinks" },
  { name: "Gym", category: "Exercise" }
];

// Test case 1: Filter by "Nature"
// Expected: Only "Park" marker passes
const filtered1 = testMarkers.filter(m => m.category === "Nature");
console.assert(filtered1.length === 1, "Should filter to one Nature place");
console.assert(filtered1[0].name === "Park", "Should be Park");

// Test case 2: Filter by "all"
// Expected: All markers pass
const filtered2 = testMarkers.filter(m => true);
console.assert(filtered2.length === 3, "Should show all places");
```

#### Pattern: Testing Map Interactions

**Challenge**: Leaflet map requires initialization and DOM.

**Approach**:
1. Test marker creation logic separately from map
2. Test popup content generation
3. Test filter application (marker array manipulation)
4. Verify map behavior manually in browser

**Example**:
```javascript
// Test: createMarkerPopupHtml() generates correct HTML
function createMarkerPopupHtml(place) {
  return `
    <div class="popup-content">
      <div class="popup-name">${place.name}</div>
      <img src="${place.image}" alt="${place.name}" class="popup-image">
    </div>
  `;
}

// Test case
const testPlace = {
  name: "Test Park",
  image: "https://example.com/image.jpg"
};

const html = createMarkerPopupHtml(testPlace);
console.assert(html.includes("Test Park"), "Should include place name");
console.assert(html.includes("popup-image"), "Should include image class");
console.assert(html.includes("https://example.com/image.jpg"), "Should include image URL");
```

### TDD Best Practices for This Project

#### 1. Start Small

- Begin with the simplest test case
- Add complexity gradually
- Don't try to test everything at once

#### 2. Test Behavior, Not Implementation

**Bad**: "Test that filterMarkers calls map.removeLayer"
**Good**: "Test that only matching markers are visible after filtering"

#### 3. Use Descriptive Test Names

- Describe what is being tested
- Include context (given/when/then)
- Make it clear what should happen

#### 4. Test Edge Cases

Common edge cases for this project:
- Empty arrays/lists
- Missing or null data
- Empty search queries
- All filters applied (very restrictive)
- No filters applied (show all)

#### 5. Keep Tests Simple

- One assertion per test concept
- Test one thing at a time
- Avoid complex test setup

### TDD Workflow Integration

#### When Adding a New Feature

1. **Write test descriptions** (Red)
   - What should the feature do?
   - What are the inputs/outputs?
   - What edge cases exist?

2. **Implement feature** (Green)
   - Write minimum code to pass tests
   - Verify in browser

3. **Refactor** (Refactor)
   - Improve code structure
   - Extract functions if needed
   - Ensure tests still pass

#### When Fixing a Bug

1. **Reproduce the bug**
   - Write a test that demonstrates the bug
   - Test should fail (showing the bug)

2. **Fix the bug** (Green)
   - Make the test pass
   - Verify fix works

3. **Refactor if needed** (Refactor)
   - Improve code while keeping fix

#### When Refactoring

1. **Ensure tests exist** (or write them)
   - Tests document current behavior
   - They'll catch regressions

2. **Refactor code**
   - Improve structure/readability
   - Keep behavior the same

3. **Verify tests still pass**
   - If tests fail, you broke something
   - Fix until tests pass again

### Example: TDD for Adding a New Filter

#### Step 1: Write Test Descriptions (Red)

```javascript
// Test: New filter "distance" should filter markers
// Given: markers at various distances from user location
// When: distance filter is set to "within 5 miles"
// Then: only markers within 5 miles should be visible

// Test: Distance filter should work with other filters
// Given: markers filtered by category "Nature"
// When: distance filter is also applied
// Then: only Nature markers within distance should be visible

// Test: Distance filter should handle missing user location
// Given: user location is not available
// When: distance filter is applied
// Then: all markers should be visible (fallback behavior)
```

#### Step 2: Implement Feature (Green)

```javascript
// Add distance filter UI
const distanceEl = document.getElementById("distance-slider");

// Add to filterMarkers()
function filterMarkers() {
  // ... existing filters ...
  const distance = getDistanceFilter();
  
  const filtered = markers.filter(m => {
    // ... existing filter logic ...
    const matchesDistance = !distance || isWithinDistance(m, distance);
    return matchesText && matchesCat && matchesCost && matchesAccess && matchesDistance;
  });
  
  // ... update map and list ...
}
```

#### Step 3: Refactor (Refactor)

```javascript
// Extract distance calculation
function calculateDistance(marker, userLocation) {
  // Haversine formula or Leaflet distance calculation
}

// Extract filter application
function applyFilters(markers, filters) {
  return markers.filter(m => {
    return matchesAllFilters(m, filters);
  });
}
```

### Benefits of TDD for This Project

1. **Better Design**: Writing tests first forces you to think about API/interface
2. **Documentation**: Tests serve as executable documentation
3. **Confidence**: Know that changes don't break existing functionality
4. **Edge Cases**: TDD encourages thinking about edge cases upfront
5. **Refactoring Safety**: Tests catch regressions during refactoring

### Limitations and Adaptations

#### Limitations in This Context

- No automated test runner (manual verification required)
- DOM/map testing requires browser environment
- No build process for test setup

#### Adaptations

- Use test descriptions as documentation
- Verify manually but systematically
- Use browser console for quick tests
- Consider adding a simple test framework if project grows

### When to Use TDD

**Good for**:
- New features with clear requirements
- Bug fixes (write test that reproduces bug)
- Refactoring (ensure tests exist first)
- Complex logic that needs verification

**Less critical for**:
- Simple UI tweaks
- Styling changes
- One-off experiments
- Very simple functions

### Skill Application Checklist

When applying TDD:
- [ ] Understand the requirement clearly
- [ ] Write test descriptions (what should happen)
- [ ] Implement minimum code to pass tests
- [ ] Verify behavior in browser
- [ ] Refactor if needed
- [ ] Document test cases for future reference
- [ ] Consider edge cases
- [ ] Update documentation if behavior changes

### Related Resources

- `.cursor/workflows.md` - Development workflows
- `.cursor/patterns.md` - Coding patterns
- `.cursor/architecture.md` - Understanding the codebase structure

### Remember

TDD is a skill that improves with practice. Start with simple cases, write test descriptions even if you can't automate them, and use TDD thinking to guide your development. The goal is better code through test-first thinking, not perfect test coverage.
