import { ChakraProvider } from '@chakra-ui/react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Layout from './components/Layout'
import HomePage from './pages/HomePage'
import EditPage from './pages/EditPage'
import StatsPage from './pages/StatsPage'

function App() {
  return (
    <ChakraProvider resetCSS>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route index element={<HomePage />} />
            <Route path="edit/:shortCode" element={<EditPage />} />
            <Route path="stats/:shortCode" element={<StatsPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </ChakraProvider>
  )
}

export default App
