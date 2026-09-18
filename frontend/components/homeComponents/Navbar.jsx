import logo from "../../src/assets/logo.png"
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";
import { NavLink } from "react-router-dom";
function Navbar({currentStatus = "Logged Out", color = "red"}) {
  const navigate = useNavigate();

  if (currentStatus === "Logged In"){
    color = "green";
  }

  const navigateToLocation = (location)=>{
    navigate(location);
  }

  return (
    <header className="navbar">
      <div className="brand">
        <div className="brand-icon">
            <img src={logo} alt="Pixel forge" />
        </div>

        <div>
          <h2>PixelForge</h2>
          <span>Generate with us</span>
        </div>
      </div>

      <nav className="nav-links">
        <NavLink to="/">Home</NavLink>
        <NavLink to="/output">Output</NavLink>
        <NavLink to="/history">History</NavLink>
        <NavLink to="/contact">Contact</NavLink>
      </nav>

      <nav className="auth-links">
        <NavLink to="/login">Login</NavLink>
        <NavLink to="/status" style={{color:color}}>{currentStatus}</NavLink>
      </nav>

    </header>
  );
}

export default Navbar;