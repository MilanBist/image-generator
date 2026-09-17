import logo from "../../public/logo.png"

function Navbar() {
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
        <a className="active" href="#studio">Home</a>
        <a href="#output">Output</a>
        <a href="#history">History</a>
        <a href="#settings">Contact us</a>
      </nav>

      <nav className="auth-links">
        <a href="#Login">Login</a>
        <a href="#Status">In</a>
      </nav>


    </header>
  );
}

export default Navbar;