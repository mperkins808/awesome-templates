import { Link } from "react-router-dom";
import styles from "./styles.module.css";

export default function About() {
  return (
    <div className={styles.container}>
      <h1>About</h1>
      <Link to={"/"}>Go back</Link>
    </div>
  );
}
