import { Routes, Route } from 'react-router-dom';
import './App.css'
import Navbar from '../components/homeComponents/Navbar';

import Home from '../components/Home';
import Register from '../components/Register';
import Login from '../components/Login';
import History from '../components/History'
import Output from '../components/Output';
import UploadedFiles from '../components/UploadedFiles';


import { useState } from 'react';


function App() {

  const[currentStatus, setCurrentStatus] = useState("Logged Out")

  return (
    <Routes>
      <Route path="/" element={
        <>
          <Navbar currentStatus={currentStatus}/>
          <Home />
        </>
      }
         />


      <Route path="/history" element={
        <>
          <Navbar currentStatus={currentStatus}/>
          <History />
        </>
      } />
      <Route path="/uploadedFiles" element={
        <>
        <Navbar currentStatus={currentStatus}/>
          <UploadedFiles />
        </>
      } />

      <Route path="/output" element={
        <>
          <Navbar currentStatus={currentStatus}/>
          <Output />
        </>
      } />
      <Route path="/login" element={<Login setCurrentStatus={setCurrentStatus}/>} />
      <Route path="register" element={<Register setCurrentStatus={setCurrentStatus}/>} />
      
    </Routes>
  )
}

export default App;
