import { useState } from 'react'
import { Routes, Route } from 'react-router-dom';
import './App.css'


import Home from '../components/Home';
import Register from '../components/Register';
import Login from '../components/Login';
import History from '../components/History'

function App() {

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />
      <Route path="register" element={<Register />} />
      <Route path="/history" element={<History />} />
    </Routes>
  )
}

export default App;
