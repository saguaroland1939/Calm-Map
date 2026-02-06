# Agent Skills

Reusable techniques and approaches agents should apply when working on this codebase.

## Skill: Test-Driven Development (TDD)

### The Red-Green-Refactor Cycle

1. **Red**: Write a failing test (or test description) for the desired behavior
2. **Green**: Write the minimum code to make it pass
3. **Refactor**: Improve code while keeping tests passing

### TDD Workflow for This Project

Since this is a vanilla JS project with no test framework, use **test-driven thinking**:

**Step 1 -- Write test descriptions (Red)**:
```javascript
// Test: filterMarkers() should filter by search query
// Given: markers ["Coffee Shop", "Park", "Museum"]
// When: search query is "Coffee"
// Then: only "Coffee Shop" marker should be visible

// Test: filterMarkers() should be case-insensitive
// Given: marker "Coffee Shop"
// When: search query is "coffee"
// Then: marker should be visible
```

**Step 2 -- Implement minimum code (Green)**: Write the simplest code that satisfies the test descriptions.

**Step 3 -- Refactor**: Extract helpers, improve naming, simplify logic -- verify behavior still matches descriptions.

### Testing Approaches

For browser-console verification:
```javascript
// Quick assertions
console.assert(filtered.length === 1, "Should filter to one result");
console.assert(filtered[0].name === "Park", "Should be Park");

// DOM verification
console.log(document.querySelectorAll('.place-item').length);
```

For this project, separate testable logic from DOM/map code when possible (e.g., filter predicates can be pure functions).

### When to Apply TDD

**New features**: Write test descriptions first to clarify requirements and edge cases.

**Bug fixes**: Write a test description that reproduces the bug, fix it, verify.

**Refactoring**: Ensure test descriptions exist for current behavior before changing code.

### Best Practices

- Test behavior, not implementation ("markers are visible" not "removeLayer was called")
- Use Given/When/Then format for clarity
- Start with the simplest case, add complexity gradually
- Test edge cases: empty arrays, missing data, no filters, all filters applied

### When TDD Is Less Critical

- Simple UI/styling tweaks
- One-off experiments
- Very simple functions

### Checklist

When applying TDD:
- [ ] Write test descriptions before code
- [ ] Implement minimum code to pass
- [ ] Verify in browser
- [ ] Refactor if needed
- [ ] Consider edge cases
- [ ] Update documentation if behavior changes
