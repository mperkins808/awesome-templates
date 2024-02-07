import "./App.css";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";

// pages
import Home from "./pages/Home/Home";
import About from "./pages/About/About";
export default function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
        </Routes>
      </Router>
    </>
  );
}
