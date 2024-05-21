import React from 'react'
import ReactDOM from 'react-dom/client'

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom"

import HomePage from './pages/HomePage/HomePage'
import CreateUser from './pages/CreateUser/CreateUser'
import Restaurants from './pages/Restaurants/Restaurants'
import UserPage from './pages/UserPage/UserPage'
import ErrorPage from './pages/Error/ErrorPage'

import './index.css'

const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
    errorElement: <ErrorPage />
  },
  {
    path: "/create-account",
    element: <CreateUser />,
    errorElement: <ErrorPage />
  },
  {
    path: "/restaurants",
    element: <Restaurants />,
    errorElement: <ErrorPage />
  },
  {
    path: "/userpage",
    element: <UserPage />,
    errorElement: <ErrorPage />
  }
])

const root = ReactDOM.createRoot(document.getElementById('root'))
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)

