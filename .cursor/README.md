# .cursor Directory Guide

This directory contains organized documentation and resources specifically designed to help AI coding agents understand and work effectively with the Calm-Map codebase.

## Purpose

The `.cursor/` directory serves as a structured knowledge base for AI agents. While `.cursorrules` provides high-level project rules and conventions, this directory contains detailed documentation that agents can reference when working on specific aspects of the project.

## File Structure

- **`README.md`** (this file): Overview of the `.cursor/` directory and how to use it
- **`architecture.md`**: Detailed technical architecture, component relationships, and data flow
- **`patterns.md`**: Common coding patterns, examples, and best practices used in this project
- **`workflows.md`**: Development workflows, common tasks, and step-by-step procedures
- **`skills.md`**: Reusable agent skills and techniques (e.g., TDD, refactoring)

## How Agents Should Use This Directory

1. **When starting a task**: Read relevant files to understand context
2. **When modifying code**: Reference patterns and architecture docs
3. **When adding features**: Check workflows for established procedures
4. **When applying techniques**: Reference skills for approaches like TDD
5. **After changes**: Update relevant documentation files (see Documentation Maintenance in `.cursorrules`)

## For Developers

These files are living documentation. As the codebase evolves, keep these files updated to ensure AI agents have accurate information. The more accurate and detailed these files are, the better AI agents can assist with development tasks.

## Learning Resource

This directory structure also serves as a learning example for:
- How to organize agent-friendly documentation
- What information agents need to be effective
- How to maintain documentation alongside code
