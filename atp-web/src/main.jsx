import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
// Import your mock worker
import worker from "./mocks/browser";

console.log('Vite mode:', import.meta.env.MODE);

// Start the worker in mock
if (import.meta.env.MODE === "development") {
    console.log("development environment");
    await worker.start();
}



createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
