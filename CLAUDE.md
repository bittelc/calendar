# Cole Calendar - Project Documentation

## Important Notes for AI Assistants
**CRITICAL**: When working on this project, you MUST update this CLAUDE.md file and the README.md file whenever:
- New design decisions are made
- Architecture changes or additions are implemented
- Technical choices are finalized
- Feature requirements are added or modified
- Development approaches are established

This ensures continuity across different AI sessions and maintains accurate project documentation.

## Project Vision
A better calendar application that inverts the traditional Google Calendar layout structure. Built in Go as a web application with calendar synchronization capabilities.

## Core Design Decisions

### Interface Design
- **Inverted Layout**: Time of day on X-axis, days of week on Y-axis (opposite of Google Calendar's standard layout)
- **Platform**: Web application (initially)
- **Target**: Personal calendar management with external sync capabilities

### Technical Architecture
- **Backend**: Go web application
- **Frontend**: JavaScript frontend with Go REST API backend
- **Database**: SQLite (chosen for simplicity, single-file deployment, and zero configuration)
- **Framework**: Gin (chosen for robust HTTP handling, JSON support, and extensive middleware ecosystem)
- **Authentication**: Simple session-based auth (start simple, migrate to OAuth when adding external calendar integrations)

### Key Features
1. **Calendar Synchronization**
   - Pull events from external calendars (Google Calendar, Proton Calendar)
   - Push events to external calendars when possible
   - Bidirectional sync where supported

2. **Event Groups & Sharing**
   - Group related events together (e.g., "Child's School Events", "Work Projects")
   - Share entire groups with specific users with different permission levels (view/edit/admin)
   - Multi-user collaboration on event groups with invitation system
   - Events can belong to multiple groups simultaneously

3. **Unique Interface**
   - Inverted time/day axis compared to traditional calendar views
   - Focus on improving user experience over existing calendar applications

## Architecture Notes
- Go backend serving REST API + static files
- JavaScript frontend for interactive calendar interface
- External calendar integration via CalDAV/REST APIs
- Authentication system for external calendar access
- Real-time sync capabilities

## Development Approach
- Start with core calendar functionality
- Add external sync integrations incrementally
- Focus on the unique interface design as primary differentiator
- Test-Driven Development (TDD): All backend APIs must be developed using TDD methodology (Red-Green-Refactor cycle)

## Critical Design Questions to Resolve
These are key decisions that need to be made before implementation:

### Frontend Implementation
1. **JavaScript Framework Choice**: React/Vue/Svelte vs vanilla JavaScript
2. **Calendar Rendering Approach**: How to efficiently render the inverted grid (canvas, SVG, or DOM elements)
3. **CSS Strategy**: Tailwind, CSS modules, styled-components, or vanilla CSS
4. **Build Tooling**: Vite, Webpack, or simple bundling approach

### Data & API Design
1. **Event Data Model**: Structure for events (title, time, metadata, external calendar IDs, recurring patterns)
2. **API Endpoint Structure**: RESTful design (`/api/events`, `/api/calendars`, etc.)
3. **Real-time Updates**: WebSockets, Server-Sent Events, or polling for calendar sync updates
4. **Time Zone Handling**: Client-side vs server-side timezone conversion

### Performance & Architecture
1. **Calendar Data Loading**: How much data to load at once, lazy loading strategy
2. **Static File Serving**: Go serves frontend files vs separate deployment
3. **Caching Strategy**: Browser and server-side caching for calendar data
4. **Development Workflow**: Hot reload and build process integration
