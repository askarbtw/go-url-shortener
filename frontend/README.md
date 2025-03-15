# URL Shortener Frontend

A React frontend for the URL Shortener API.

## Features

- Create short URLs
- View short URL details
- Update existing short URLs
- View access statistics
- Redirect to original URLs

## Technology Stack

- React 18
- TypeScript
- Vite
- React Router
- Chakra UI
- Axios

## Getting Started

### Prerequisites

- Node.js 14.x or later
- npm or yarn

### Installation

1. Install dependencies:

```bash
npm install
```

### Development

To start the development server:

```bash
npm run dev
```

The application will be available at `http://localhost:5173`.

### Building for Production

To create a production build:

```bash
npm run build
```

The build output will be in the `dist` directory.

### Integration with Backend

This frontend is designed to work with the URL Shortener Go backend. Make sure the backend is running at `http://localhost:8080` (or update the API_BASE_URL in `src/services/api.ts` to match your backend URL).

In development mode, API requests are proxied to avoid CORS issues.
