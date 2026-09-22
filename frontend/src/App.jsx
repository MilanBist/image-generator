import { Routes, Route } from 'react-router-dom';
import './App.css'
import Navbar from '../components/homeComponents/Navbar';

import Home from '../components/Home';
import Register from '../components/Register';
import Login from '../components/Login';
import History from '../components/History'
import Output from '../components/Output';
import UploadedFiles from '../components/UploadedFiles';


import { useEffect, useState } from 'react';


function App() {  
  const[currentStatus, setCurrentStatus] = useState("Logged Out");
  const[output, setOuptut] = useState([]);
  const[history, setHistory] = useState([]);
  const[uploaded, setUploaded] = useState([]);
  

  // if there is accesstoken and refresh token the user is normally logged in
  let refreshToken = localStorage.getItem("refreshToken");
  useEffect(()=>{
    if (refreshToken !== null){
      setCurrentStatus("Logged In");
    }
  }, []);

  console.log("Current status is: ", currentStatus);

  return (
    <Routes>
      <Route path="/" element={
        <>
          <Navbar currentStatus={currentStatus}/>
          <Home setOutputSection={setOuptut} setHistorySection={setHistory} setUploadedSection={setUploaded}/>
        </>
      }
          />


      <Route path="/history" element={
        <>
          <Navbar currentStatus={currentStatus} />
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
