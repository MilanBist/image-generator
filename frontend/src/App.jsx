import { Routes, Route } from 'react-router-dom';
import './App.css'


import Home from '../components/Home';
import Register from '../components/Register';
import Login from '../components/Login';
import History from '../components/History'
import Output from '../components/Output';
import UploadedFiles from '../components/UploadedFiles';


function App() {

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />
      <Route path="register" element={<Register />} />
      <Route path="/history" element={<History />} />
      <Route path="/uploadedFiles" element={<UploadedFiles />} />
      <Route path="/output" element={<Output />} />
    </Routes>
  )
}

export default App;
