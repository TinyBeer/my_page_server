import React, { useState } from 'react';
import LoginModal from './login';
import Game from './game';
import Nav from './nav';
import Memo from './memo';
import Cinema from './cinema';
import Video from './video';
import ErrorPage from './error';

import { createBrowserRouter, RouterProvider } from 'react-router-dom';

export default function App() {
  const router = createBrowserRouter([
    {
      path: '/login',
      element: <LoginModal />,
      errorElement: <ErrorPage />,
    },
    {
      path: '/',
      element: <Nav />,
      errorElement: <ErrorPage />,
    },
    {
      path: '/nav',
      element: <Nav />,
      errorElement: <ErrorPage />,
    },
    {
      path: '/game',
      element: <Game />,
      errorElement: <ErrorPage />,
    },
    {
      path: '/memo',
      element: <Memo />,
      errorElement: <ErrorPage />,
    },
    {
      path: '/cinema',
      element: <Cinema />,
      errorElement: <ErrorPage />,
      children: [
        {
          path: '',
          element: <Video />,
        },
        {
          path: '/cinema/video',
          element: <Video />,
        },
      ],
    },
  ]);
  return <RouterProvider router={router} />;
}
