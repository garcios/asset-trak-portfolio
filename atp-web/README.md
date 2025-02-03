# Web application

## Create new project
```shell
npm create vite@latest atp-web -- --template react
cd atp-web
npm install
```

## GraphQL client
```js
import './App.css'
import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import BalanceSummary from "./components/BalanceSummary.tsx";

const client = new ApolloClient({
    uri: "/query",
    cache: new InMemoryCache(),
});

function App() {

    return (
        <ApolloProvider client={client}>
            <BalanceSummary/>
        </ApolloProvider>
    );
}

export default App
```

## Setup Proxy in React Vite for CORS
When developing proof-of-concept applications on your local machine using Vite React, you may encounter CORS errors 
while attempting to integrate the front end, designed with React, and the backend API or GraphQL.

See the vite.config.js for configuration of the proxy.

```js
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Target the GraphQL gateway
      '/query': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        //rewrite: (path) => path.replace(/^\/api/, ''),

        configure: (proxy, options) => {
          proxy.on('error', (err, _req, _res) => {
            console.log('error', err);
          });
          proxy.on('proxyReq', (proxyReq, req, _res) => {
            console.log('Request sent to target:', req.method, req.url);
          });
          proxy.on('proxyRes', (proxyRes, req, _res) => {
            console.log('Response received from target:', proxyRes.statusCode, req.url);
          });
        },
      },
    },
  },
})
```

## Run the app
```shell
npm run dev
```

## Material UI
```shell
npm install @mui/material @emotion/react @emotion/styled
```
