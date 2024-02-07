import { Link } from "react-router-dom";
import styles from "./styles.module.css";

export default function Home() {
  return (
    <div className={styles.container}>
      <h1>Home</h1>
      <strong>The Stack</strong>
      <ul>
        <li>
          <strong>Frontend</strong> React
        </li>
        <li>
          <strong>Bundler</strong> Vite
        </li>
        <li>
          <strong>Backend</strong> Go
        </li>
      </ul>
      <Link to={"/about"}>Go to About page</Link>
    </div>
  );
}
